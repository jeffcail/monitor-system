package machine

import (
	"bufio"
	"fmt"
	"os"

	"bz.service.cloud.monitoring/client/daos"
	"github.com/google/uuid"

	"github.com/labstack/echo"

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
	//
	// 从缓存中拿机器码
	getMachineCode, err := daos.GetMachineCode(ip)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("从缓存中获取唯一机器码失败"), zap.Error(err))
	}

	if getMachineCode == "" {
		// 从数据库中拿
		m, err := daos.GetMachineInfoFromDbByIp(ip)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("从数据库中获取唯一机器码失败"), zap.Error(err))
		}
		if m.Id == 0 {
			// 数据库中不存在 第一次部署

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

			writeFirstVersion()
		}
	}

}

// writeFirstVersion
func writeFirstVersion() {
	file, err := os.OpenFile("./version.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("打开文件失败"), zap.Error(err))
	}
	defer file.Close()

	str := "1"
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(str)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("版本号写入缓存失败"), zap.Error(err))
	}
	err = writer.Flush()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("版本号写入文件失败"), zap.Error(err))
	}
}

func machineCode(ip, uid string) string {
	return utils.Md5s(ip + "-" + uid)
}

func getIp() string {
	return utils.GetIP()
}

// ReceiveCom
func ReceiveCom(c echo.Context) error {
	content := c.FormValue("content")
	err := utils.ExecCommand(content)
	if err != nil {
		ubzer.MLog.Error("指令执行失败", zap.Error(err))
		return err
	}
	return nil
}
