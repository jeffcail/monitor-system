package machine

import "os"

// GetHostName
// 获取系统主机名
func GetHostName() (string, error) {
	hostname, err := os.Hostname()
	return hostname, err
}
