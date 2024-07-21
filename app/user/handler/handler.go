package handler

import (
	"imlogic/common/db"
	"imlogic/internal/xjwt"
)

type PublicSrvImpl struct {
	Model db.Model
	Jwt   *xjwt.Jwt
}

func NewPublicSrvImpl() *PublicSrvImpl {
	return &PublicSrvImpl{

		Model: db.NewModel(),
		Jwt:   xjwt.NewJwt(),
	}

}
