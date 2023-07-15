package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/xu756/imlogic/app/public/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	Captcha     *base64Captcha.Captcha
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	driver := base64Captcha.DriverMath{
		Height:          100,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
	}
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	store := base64Captcha.NewMemoryStore(1000000, 120)
	return &ServiceContext{
		Config:      c,
		Captcha:     base64Captcha.NewCaptcha(driver.ConvertFonts(), store),
		RedisClient: redisClient,
	}
}
