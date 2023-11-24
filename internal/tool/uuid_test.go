package tool

import "testing"

func TestRandomCode(t *testing.T) {
	code := RandomCode(6)
	t.Log(code)
}
