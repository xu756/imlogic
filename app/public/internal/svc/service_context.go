package svc

import (
	"github.com/xu756/imlogic/app/model"
	"github.com/xu756/imlogic/app/public/internal/config"
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	Jwt       *xjwt.Jwt
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Jwt:       xjwt.NewJwt(c.Jwt),
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DbSource), c.Cache),
	}
}
