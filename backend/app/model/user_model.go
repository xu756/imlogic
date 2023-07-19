package model

import (
	"context"
	"fmt"
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
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *defaultUserModel) LoginByPassword(ctx context.Context, username string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `name` = ? or `mobile` = ? ", userRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, username, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, xerr.NewMsgError("用户不存在,请注册")
	default:
		return nil, xerr.NewDbErr("查询失败", err)
	}
}
