package tool

import (
	"github.com/google/uuid"
	"github.com/xu756/imlogic/internal/xerr"
)

func NewUid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", xerr.SystemErr()
	}
	return uid.String(), nil
}
