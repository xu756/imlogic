package ip

import (
	"testing"
)

func TestGetCurrentIp(t *testing.T) {
	ip := GetCurrentIp()
	if ip == "" {
		t.Error("ip is empty")
	}
	t.Logf("ip: %s", ip)

}
