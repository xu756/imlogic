package tool

import (
	"github.com/xu756/imlogic/pkg/xerr"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

func NewUuid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", xerr.SystemErr()
	}
	return uid.String(), nil
}

// RandomCode 随机生成n位数
func RandomCode(n int) string {
	// 生成n位数
	var code string
	for i := 0; i < n; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code

}
