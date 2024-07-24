package db

import (
	"context"
	"imlogic/ent"
	"imlogic/ent/user"
	"imlogic/internal/xerr"
)

type dbUserModel interface {
	GetOneUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error)
	GetOneUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error)
	FindUserByUUID(ctx context.Context, uuid string) (userInfo *ent.User, err error)
}

// 根据用户名查找用户
func (m *customModel) GetOneUserByUsername(ctx context.Context, username string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.Username(username), user.Deleted(false)).First(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, xerr.WarnMsg("用户不存在 用户名:%s", username)
	case err != nil:
		return nil, xerr.DbErr(err, "查询用户失败 用户名:%s", username)
	default:
		return userInfo, nil
	}
}

// 根据手机号查找用户
func (m *customModel) GetOneUserByMobile(ctx context.Context, mobile string) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Query().
		Where(user.Mobile(mobile), user.Deleted(false)).First(ctx)
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
func (m *customModel) CreateUser(ctx context.Context, username, password, mobile string, editor int64) (userInfo *ent.User, err error) {
	userInfo, err = m.client.User.Create().
		SetUsername(username).
		SetPassword(password).
		SetMobile(mobile).
		SetEditor(editor).
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

func (m *customModel) FindUserByUUID(ctx context.Context, uuid string) (userInfo *ent.User, err error) {
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
