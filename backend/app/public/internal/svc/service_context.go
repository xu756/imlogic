package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/xu756/imlogic/app/model"
	"github.com/xu756/imlogic/app/public/internal/config"
	"github.com/xu756/imlogic/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	Captcha       *base64Captcha.Captcha
	RedisClient   *redis.Redis
	UserModel     model.UserModel
	UserRoleModel model.UserRoleModel
	Jwt           *xjwt.Jwt
}

func NewServiceContext(c config.Config) *ServiceContext {
	driver := base64Captcha.DriverMath{
		Height:          100,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
	}
	redisClient, err := redis.NewRedis(c.RedisConf)
	conn := sqlx.NewMysql(c.DbSource)
	if err != nil {
		panic(err)
	}
	store := base64Captcha.NewMemoryStore(1000000, 120)
	return &ServiceContext{
		Config:        c,
		Captcha:       base64Captcha.NewCaptcha(driver.ConvertFonts(), store),
		RedisClient:   redisClient,
		UserModel:     model.NewUserModel(conn, c.Cache),
		UserRoleModel: model.NewUserRoleModel(conn, c.Cache),
		Jwt:           xjwt.NewJwt(c.Jwt),
	}
}
