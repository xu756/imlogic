package tool

import (
	"log"

	"github.com/google/uuid"
	"github.com/xu756/imlogic/internal/xerr"
)

func NewUid() string {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Print(xerr.NewSystemError("生成uid错误"))
	}
	return uid.String()
}
