package db

import (
	"context"
	"github.com/xu756/imlogic/ent"
	"github.com/xu756/imlogic/ent/user"
	"github.com/xu756/imlogic/internal/xerr"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, username, device string) (userInfo *ent.User, err error)
	FindUserByMobile(ctx context.Context, mobile, device string) (userInfo *ent.User, err error)
	FindUserByUUID(tx *ent.Tx, uuid, device string) (userInfo *ent.User, err error)
	CreateUser(ctx context.Context, username, password, mobile, device string, creator int64) (userInfo *ent.User, err error)
	UpdateUserAvatar(ctx context.Context, uuid string, avatar, device string, editor int64) (userInfo *ent.User, err error)
}

// FindUserByUsername 根据用户名查找用户
func (m *customModel) FindUserByUsername(ctx context.Context, username, device string) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, xerr.DbConnectErr()
	}
	userInfo, err = tx.User.Query().
		Where(user.Username(username), user.Deleted(false), user.Device(device)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbFindErr()
	default:
		return userInfo, tx.Commit()
	}
}

// FindUserByUUID 根据UUID查找用户
func FindUserByUUID(tx *ent.Tx, uuid string) (userInfo *ent.User, err error) {
	userInfo, err = tx.User.Query().
		Where(user.UUID(uuid), user.Deleted(false)).First(context.Background())
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbFindErr()
	default:
		return userInfo, nil
	}
}

// FindUserByMobile 根据手机号查找用户
func (m *customModel) FindUserByMobile(ctx context.Context, mobile, device string) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, xerr.DbConnectErr()
	}
	userInfo, err = tx.User.Query().
		Where(user.Mobile(mobile), user.Device(device), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbFindErr()
	default:
		return userInfo, tx.Commit()
	}
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, username, password, mobile, device string, creator int64) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, xerr.DbConnectErr()
	}
	userInfo, err = tx.User.Create().
		SetUsername(username).
		SetPassword(password).
		SetMobile(mobile).
		SetCreator(creator).
		SetEditor(creator).
		SetDevice(device).
		SetVersion(1).
		Save(ctx)
	switch {
	case ent.IsConstraintError(err):
		return nil, xerr.ErrMsg(xerr.UserExist)
	case err != nil:
		return nil, xerr.DbCreateErr()
	default:
		return userInfo, tx.Commit()
	}
}

// UpdateUserAvatar 更新用户头像
func (m *customModel) UpdateUserAvatar(ctx context.Context, uuid string, avatar, device string, editor int64) (userInfo *ent.User, err error) {
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return nil, xerr.DbConnectErr()
	}
	User, err := FindUserByUUID(tx, uuid)
	if err != nil {
		return nil, err
	}
	if User.Avatar != avatar {
		User, err := tx.User.UpdateOne(User).
			SetAvatar(avatar).
			SetEditor(editor).
			SetVersion(User.Version + 1).
			Where(user.UUID(uuid)).
			Save(ctx)
		if err != nil {
			return nil, xerr.DbUpdateErr()
		}
		return User, tx.Commit()
	} else {
		return User, tx.Commit()
	}
}

func (m *customModel) FindUserByUUID(tx *ent.Tx, uuid, device string) (userInfo *ent.User, err error) {
	userInfo, err = tx.User.Query().
		Where(user.UUID(uuid), user.Deleted(false)).First(context.Background())
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.ErrMsg(xerr.UserNotExist)
	case err != nil:
		return nil, xerr.DbFindErr()
	default:
		return userInfo, nil
	}
}
