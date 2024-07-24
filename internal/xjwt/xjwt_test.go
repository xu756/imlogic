package xjwt

import (
	"imlogic/common/config"
	"testing"
)

func TestNewJwt(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	jwt := NewJwt()
	token, err := jwt.NewJwtToken(1, "222222")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
