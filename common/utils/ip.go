package utils

import "net"

var ip = ""

// GetIP
func GetIP() string {
	if ip != "" {
		return ip
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("获取本机地址失败: " + err.Error())
	}
	for _, addrss := range addrs {
		if ipNet, ok := addrss.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.IsPrivate() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				return ip
			}
		}
	}
	panic("获取本机地址失败")
}
