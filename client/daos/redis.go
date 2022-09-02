package daos

import (
	"errors"
	"fmt"

	"bz.service.cloud.monitoring/common/db"
)

// SetMachineCode
func SetMachineCode(ip, code string) error {
	_, err := db.Rc.Set(ip, code, -1).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetMachineCode
func GetMachineCode(ip string) (string, error) {
	b, err := db.Rc.Exists(ip).Result()
	if err != nil {
		return "", err
	}
	if !b {
		return "", errors.New(fmt.Sprintf("ip: %v key不存在", ip))
	}
	result, err := db.Rc.Get(ip).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
