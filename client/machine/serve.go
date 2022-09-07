package machine

import (
	"fmt"
	"net/http"
	"os"

	"bz.service.cloud.monitoring/client/config"

	"bz.service.cloud.monitoring/common/utils"

	"bz.service.cloud.monitoring/common/ubzer"
	"go.uber.org/zap"

	"github.com/labstack/echo"
)

// ServeUpgrade
func ServeUpgrade(c echo.Context) error {
	pn := c.FormValue("package_name")
	pt := c.FormValue("package_path")

	_, err := os.Stat("./" + pn)
	if !os.IsNotExist(err) {
		utils.ExecCommand("rm -f ./" + pn)
	}

	err = utils.ExecCommand("wget http://" + config.Config().GoFileServe + "/dl/" + pn)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("服务升级拉取升级需要的包文件失败"), zap.Error(err))
		return c.JSON(http.StatusOK, "服务升级拉取升级需要的包文件失败")
	}

	cmd := "mv ./" + pn + " " + pt
	err = utils.ExecCommand(cmd)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("移动包到项目目录下失败"), zap.Error(err))
		return c.JSON(http.StatusOK, "移动包到项目目录下失败")
	}

	cmd2 := "chmod +x " + pt + "/" + pn
	ubzer.MLog.Info(fmt.Sprintf("========== cmd2: %v", cmd2))
	err = utils.ExecCommand(cmd2)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("赋予升级包可执行权限失败"), zap.Error(err))
		return c.JSON(http.StatusOK, "赋予升级包可执行权限失败")
	}

	cmd3 := "cd " + pt + " && ./start.sh"
	go upgrade(cmd3)

	return c.JSON(http.StatusOK, "success")
}

// upgrade
func upgrade(cmd string) {
	ubzer.MLog.Info("====== 开始升级服务 ======")
	ubzer.MLog.Info(fmt.Sprintf("========== cmd3: %v", cmd))
	err := utils.ExecCommand(cmd)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("开始升级服务失败"), zap.Error(err))
	}
}
