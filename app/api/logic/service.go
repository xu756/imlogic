package logic

import "imlogic/internal/xjwt"

var service *Service

type Service struct {
	Jwt *xjwt.Jwt
}

func Init() {
	service = &Service{
		Jwt: xjwt.NewJwt(),
	}
}
