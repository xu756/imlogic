package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AccountModel = (*customAccountModel)(nil)

type (
	// AccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountModel.
	AccountModel interface {
		accountModel
	}

	customAccountModel struct {
		*defaultAccountModel
	}
)

// NewAccountModel returns a model for the database table.
func NewAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AccountModel {
	return &customAccountModel{
		defaultAccountModel: newAccountModel(conn, c, opts...),
	}
}
