package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/user"
	"imlogic/internal/xerr"
)

type dbUserModel interface {
	FindUserByUsername(ctx context.Context, username, device string) (userInfo *ent.User, err error)
	FindUserByMobile(ctx context.Context, mobile, device string) (userInfo *ent.User, err error)
	FindUserByUUID(ctx context.Context, uuid, device string) (userInfo *ent.User, err error)
}

// FindUserByUsername 根据用户名查找用户
func (m *customModel) FindUserByUsername(ctx context.Context, username, device string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.Username(username), user.Deleted(false), user.Device(device)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 用户名:%s", username)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 用户名:%s", username)
	default:
		return userInfo, nil
	}
}

// FindUserByMobile 根据手机号查找用户
func (m *customModel) FindUserByMobile(ctx context.Context, mobile, device string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.Mobile(mobile), user.Device(device), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 手机号:%s", mobile)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 手机号:%s", mobile)
	default:
		return userInfo, nil
	}
}

// CreateUser 创建用户
func (m *customModel) CreateUser(ctx context.Context, username, password, mobile, device string, creator int64) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Create().
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
		return nil, xerr.WarnMsg("用户已存在 用户名:%s 手机号:%s", username, mobile)
	case err != nil:
		return nil, xerr.DbErr(err, "创建用户失败 用户名:%s 手机号:%s", username, mobile)
	default:
		return userInfo, nil
	}
}

func (m *customModel) FindUserByUUID(ctx context.Context, uuid, device string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.UUID(uuid), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 用户uuid:%s", uuid)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 用户uuid:%s", uuid)
	default:
		return userInfo, nil
	}
}
