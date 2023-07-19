package model

import (
	"context"
	"fmt"
	"github.com/xu756/imlogic/internal/xerr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserGroupModel = (*customUserGroupModel)(nil)
var cacheUserGroupsPrefix = "cache:userGroups:id:"

type (
	// UserGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupModel.
	UserGroupModel interface {
		userGroupModel
		FindUserGroups(ctx context.Context, userId int64) ([]int64, error)
	}

	customUserGroupModel struct {
		*defaultUserGroupModel
	}
)

// NewUserGroupModel returns a model for the database table.
func NewUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserGroupModel {
	return &customUserGroupModel{
		defaultUserGroupModel: newUserGroupModel(conn, c, opts...),
	}
}

// FindUserGroups 查询用户组
func (m *customUserGroupModel) FindUserGroups(ctx context.Context, userId int64) ([]int64, error) {
	userGroupsKey := fmt.Sprintf("%s%v", cacheUserGroupsPrefix, userId)
	var resp []int64
	err := m.QueryRowCtx(ctx, &resp, userGroupsKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select `id` from %s where `user_id` = ? ", m.table)
		return conn.QueryRowsCtx(ctx, v, query, userId)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return []int64{}, nil
	default:
		return nil, xerr.NewDbErr("用户组查询失败", err)
	}
}
