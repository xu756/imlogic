package logic

import (
	"imlogic/common/cos"
)

var service *Service

type Service struct {
	Cos *cos.Cos
}

func Init() {
	service = &Service{
		Cos: cos.NewCos(),
	}
}
