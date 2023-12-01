package tool

import (
	"math/rand"
	"strconv"
)

// RandomCode 随机生成n位数
func RandomCode(n int) string {
	// 生成n位数
	var code string
	for i := 0; i < n; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code

}
