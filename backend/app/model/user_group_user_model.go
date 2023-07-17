package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserGroupUserModel = (*customUserGroupUserModel)(nil)

type (
	// UserGroupUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupUserModel.
	UserGroupUserModel interface {
		userGroupUserModel
	}

	customUserGroupUserModel struct {
		*defaultUserGroupUserModel
	}
)

// NewUserGroupUserModel returns a model for the database table.
func NewUserGroupUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserGroupUserModel {
	return &customUserGroupUserModel{
		defaultUserGroupUserModel: newUserGroupUserModel(conn, c, opts...),
	}
}
