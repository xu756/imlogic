package ip

import (
	"fmt"
	"log"
	"net"
)

// 获取本机ip
func GetCurrentIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	log.Print(addrs)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("【 获取本机ip失败 】")
}
