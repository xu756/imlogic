package model

import (
	"context"

	"github.com/xu756/imlogic/internal/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		LoginByPassword(ctx context.Context, username string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) LoginByPassword(ctx context.Context, username string) (*User, error) {
	var resp User
	err := m.QueryRowNoCacheCtx(ctx, &resp, "select * from user where username = ? or mobile = ?", username, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, xerr.NewMsgError("用户不存在")
	default:
		return nil, xerr.NewDbErr("查询用户失败 ", err)
	}

}
