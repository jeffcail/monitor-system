package machine

import (
	"errors"
	"fmt"

	"bz.service.cloud.monitoring/common/utils"

	"bz.service.cloud.monitoring/common/ubzer"
	"go.uber.org/zap"

	"github.com/labstack/echo"
)

// ServeUpgrade
func ServeUpgrade(c echo.Context) error {
	pn := c.FormValue("package_name")

	err := utils.ExecCommand("wget http://127.0.0.1:9095/dl/" + pn)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("服务升级拉取升级需要的包文件失败"), zap.Error(err))
		return errors.New("服务升级拉取升级需要的包文件失败")
	}

	err = utils.ExecCommand("mv ./" + pn + " /root/" + pn)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("移动包到项目目录下失败"), zap.Error(err))
		return errors.New("移动包到项目目录下失败")
	}

	ubzer.MLog.Info("====== 开始升级服务 ======")
	err = utils.ExecCommand("cd /root/" + pn + " && ./start")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("开始升级服务失败"), zap.Error(err))
		return errors.New("开始升级服务失败")
	}

	return errors.New("success")
}
