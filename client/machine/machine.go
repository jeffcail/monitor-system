package machine

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"bz.service.cloud.monitoring/client/daos"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/common/utils"
)

// GenerateUniqueMachineCode
func GenerateUniqueMachineCode() {
	ip := getIp()
	name, err := GetHostName()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取机器hostname失败 机器ip: %v", ip), zap.Error(err))
	}

	// 从缓存中拿机器码
	getMachineCode, err := daos.GetMachineCode(ip)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取唯一机器码失败"), zap.Error(err))
	}

	if getMachineCode == "" {
		uid := uuid.New().String()
		code := machineCode(ip, uid)
		err = daos.SetMachineCode(ip, code)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("生成唯一机器码写入缓存失败 code: %v", code), zap.Error(err))
		}

		err = daos.SaveHostName(name, code, ip)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("生成唯一机器码之后保存服务器信息失败 ip: %v code: %v name: %v",
				ip, code, name), zap.Error(err))
		}
	}

}

// touchMachineCodeFile 后续机器码写入缓存的同时也写入文件
func touchMachineCodeFile(code string) {
	file, err := os.OpenFile("./code.txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("生成唯一机器码打开文件失败 code: %v", code), zap.Error(err))
	}
	_, err = file.WriteString(code)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("生成唯一机器码保存写入文件失败 code: %v", code), zap.Error(err))
	}
	file.Close()
}

func machineCode(ip, uid string) string {
	return utils.Md5s(ip + "-" + uid)
}

func getIp() string {
	return utils.GetIP()
}
