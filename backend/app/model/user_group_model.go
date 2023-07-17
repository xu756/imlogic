package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserGroupModel = (*customUserGroupModel)(nil)

type (
	// UserGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupModel.
	UserGroupModel interface {
		userGroupModel
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
