package job

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cast"

	"bz.service.cloud.monitoring/client/config"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/client/models"
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/common/utils"
)

// CheckClientVersion
func CheckClientVersion() {
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 ============="))
	ip := utils.GetIP()
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 拿到ip: %v =============", ip))

	m := &models.MonUpgradeMachineRecord{}
	has, err := db.Mysql.Where("machine_ip = ?", ip).Desc("id").Limit(1).Get(m)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 定时从数据库中检测客户端最新版本失败 ip: %v", ip), zap.Error(err))
		return
	}
	if !has {
		ubzer.MLog.Info(fmt.Sprintf("=======  开始检测升级 从数据库中没有拿到ip为: %v 的数据", ip))
		return
	}
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 拿到最新的升级记录: %v =============", m))

	content, err := os.ReadFile("./version.txt")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 定时检测更新客户端版本 读取版本号失败"), zap.Error(err))
		return
	}
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 当前版本号文件保存的版本为: %v =============", string(content)))
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 要升级的版本为: %v =============", m.UpgradeVersion))
	if utils.LessThenAndEqual(m.UpgradeVersion, string(content)) {
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 当前没有可升级的版本 ============="))
		return
	}

	file2, err := os.OpenFile("./version.txt", os.O_RDWR|os.O_TRUNC, 0766)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 客户端版本升级打开版本号文件失败"), zap.Error(err))
	}
	defer file2.Close()

	str := m.UpgradeVersion
	writer := bufio.NewWriter(file2)
	_, err = writer.WriteString(str)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 版本号写入缓存失败"), zap.Error(err))
	}
	err = writer.Flush()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 版本号写入文件失败"), zap.Error(err))
	}

	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 版本写入版本控制文件成功 ============="))

	pn := ip + ":" + m.PackageName + ":" + m.UpgradeVersion
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 要升级的包名: %v =============", pn))

	exec := "wget " + config.Config().GoFileServe + "/api/dl\\?file\\=" + pn
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 下载要升级的包: %v =============", exec))

	err = utils.ExecCommand(exec)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 定时检测更新客户端版本 下载对应的版本失败"), zap.Error(err))
		return
	}

	// 将原来的包重命名成对应的版本名字
	err = utils.ExecCommand("mv /root/client/client /root/client/client:" + string(content))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 重命名原包名称失败"), zap.Error(err))
		return
	}
	// 检测保留最近三份包
	err = utils.ExecCommand("rm -f /root/client/client:" + cast.ToString(cast.ToInt(m.UpgradeVersion)-3))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 删除历史包 : %v失败", "client:"+cast.ToString(cast.ToInt(m.UpgradeVersion)-3)), zap.Error(err))
		return
	}

	err = utils.ExecCommand("mv ./dl?file=" + pn + " /root/client/client")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 移动下载的客户端版本到项目目录下失败"), zap.Error(err))
		return
	}

	err = utils.ExecCommand("chmod +x /root/client/client")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 给包赋予可执行的权限失败"), zap.Error(err))
		return
	}
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 给包赋予可执行的权限成功: %v =============", pn))

	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 准备启动客户端程序 ============="))
	os.Exit(1)

	//err = utils.ExecCommand("cd /etc/systemd/system && systemctl start client-monitor.service")
	//if err != nil {
	//	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 启动客户端服务失败 ============="), zap.Error(err))
	//	return
	//}
	//ubzer.MLog.Info(fmt.Sprintf("=======  开始检测升级 定时检测更新客户端版本成功, 原版本号为: %v 更新成功的版本号为: %v", string(content),
	//	m.UpgradeVersion))

	//err = utils.ExecCommand("./" + pn + " uninstall")
	//if err != nil {
	//	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 卸载客户端服务失败 ============="), zap.Error(err))
	//	return
	//}
	//
	//err = utils.ExecCommand("./" + pn + " install")
	//if err != nil {
	//	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 安装客户端服务失败 ============="), zap.Error(err))
	//	return
	//}
	//
	//err = utils.ExecCommand("cd /etc/systemd/system && systemctl start client-monitor.service")
	//if err != nil {
	//	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 启动客户端服务失败 ============="), zap.Error(err))
	//	return
	//}

	//err = utils.ExecCommand("kill -9 `cat pidfile.txt`")
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("定时检测更新客户端版本 杀死原版本失败: %v 要更新的版本号为: %v", string(content),
	//		m.UpgradeVersion))
	//	return
	//}
	//ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 杀死要升级的服务的进程成功 ============="))
	//
	//exec2 := "nohup ./ " + pn + " > nohup.out & echo $! > pidfile.txt"
	//ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 开始启动进程的命令为: %v =============", exec2))
	//err = utils.ExecCommand(exec2)
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("定时检测更新客户端版本失败 原版本号为: %v 要更新的版本号为: %v", string(content),
	//		m.UpgradeVersion))
	//} else {
	//	ubzer.MLog.Info(fmt.Sprintf("定时检测更新客户端版本成功, 原版本号为: %v 更新成功的版本号为: %v", string(content),
	//		m.UpgradeVersion))
	//}
}
