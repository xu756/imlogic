package handler

import (
	"github.com/xu756/imlogic/common/db"
	"github.com/xu756/imlogic/internal/xjwt"
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
