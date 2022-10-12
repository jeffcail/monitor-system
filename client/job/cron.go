package job

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/websocket"

	"github.com/spf13/cast"

	"github.com/c/server-monitoring/client/config"

	"go.uber.org/zap"

	"github.com/c/server-monitoring/common/ubzer"
	"github.com/c/server-monitoring/common/utils"
)

// CheckClientVersion
func CheckClientVersion() {
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 ============="))
	ip := utils.GetIP()
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 拿到ip: %v =============", ip))

	dl := websocket.Dialer{}
	d := "ws://" + config.Config().GoFileServe + "/client/up"
	//d := "ws://192.168.0.159:9999/init/client"
	conn, _, err := dl.Dial(d, nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 开始检测升级 连接服务端websocket失败 ============="), zap.Error(err))
	}
	type Message struct {
		Ip string `json:"ip"`
	}
	m := &Message{Ip: ip}
	res, _ := json.Marshal(m)
	err = conn.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 开始检测升级 连接服务端websocket发送客户端ip失败 ============="), zap.Error(err))
	}

	_, resContent, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("============ 开始检测升级 连接服务端websocket接收服务端推送过来的数据失败 ============="), zap.Error(err))
	}

	type CMessage struct {
		PackageName    string `json:"package_name"`
		UpgradeVersion string `json:"upgrade_version"`
	}
	cm := &CMessage{}
	_ = json.Unmarshal(resContent, cm)
	fmt.Printf("============cm: %v", cm)

	if cm.UpgradeVersion == "" && cm.PackageName == "" {
		ubzer.MLog.Info(fmt.Sprintf("=======  开始检测升级 从数据库中没有拿到ip为: %v 的数据 =============", ip))
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 此客户端: %v 没有要升级的版本 =============", ip), zap.Error(err))
		return
	}

	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 拿到最新的升级记录: %v =============", m))

	content, err := os.ReadFile("./version.txt")
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 定时检测更新客户端版本 读取版本号失败 ============="), zap.Error(err))
		return
	}
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 当前版本号文件保存的版本为: %v =============", string(content)))
	ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 要升级的版本为: %v =============", cm.UpgradeVersion))
	if utils.LessThenAndEqual(cm.UpgradeVersion, string(content)) {
		ubzer.MLog.Info(fmt.Sprintf("============ 开始检测升级 当前没有可升级的版本 ============="))
		return
	}

	file2, err := os.OpenFile("./version.txt", os.O_RDWR|os.O_TRUNC, 0766)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 客户端版本升级打开版本号文件失败"), zap.Error(err))
	}
	defer file2.Close()

	str := cm.UpgradeVersion
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

	pn := ip + ":" + cm.PackageName + ":" + cm.UpgradeVersion
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
	err = utils.ExecCommand("rm -f /root/client/client:" + cast.ToString(cast.ToInt(cm.UpgradeVersion)-3))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("=======  开始检测升级 删除历史包 : %v失败", "client:"+cast.ToString(cast.ToInt(cm.UpgradeVersion)-3)), zap.Error(err))
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
}
