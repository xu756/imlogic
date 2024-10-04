package xjwt

import (
	"testing"
)

func TestNewJwt(t *testing.T) {
	jwt := NewJwt()
	token, err := jwt.NewJwtToken(1, "222222")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
