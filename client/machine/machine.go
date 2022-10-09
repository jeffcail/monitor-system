package machine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"bz.service.cloud.monitoring/client/config"

	"github.com/gorilla/websocket"

	"github.com/google/uuid"

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
	dl := websocket.Dialer{}
	d := "ws://" + config.Config().GoFileServe + "/init/client"
	//d := "ws://192.168.0.159:9092/init/client"
	conn, _, err := dl.Dial(d, nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("初始化客户端 连接服务端websocket失败"), zap.Error(err))
	}

	type SMessage struct {
		Ip          string `json:"ip"`
		HostName    string `json:"host_name"`
		MachineCode string `json:"machine_code"`
	}
	m := &SMessage{
		Ip:          ip,
		HostName:    name,
		MachineCode: machineCode(ip, uuid.New().String()),
	}

	res, _ := json.Marshal(m)
	err = conn.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("初始化客户端 连接服务端发送数据失败"), zap.Error(err))
	}
	_, content, err := conn.ReadMessage()
	fmt.Printf("=======content: %v", string(content))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("初始化客户端 连接服务端接收数据失败"), zap.Error(err))
	}

	type ResponseContent struct {
		Exists int `json:"exists"`
	}
	rc := &ResponseContent{}
	_ = json.Unmarshal(content, rc)
	fmt.Printf("============rc.exists: %d , %T", rc.Exists, rc.Exists)
	if rc.Exists == 2 {
		ubzer.MLog.Info("开始创建版本号文件")
		writeFirstVersion()
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
	ubzer.MLog.Info(fmt.Sprintf("版本号文件创建成功"))
}

func machineCode(ip, uid string) string {
	return utils.Md5s(ip + "-" + uid)
}

func getIp() string {
	return utils.GetIP()
}

// GetHostName
// 获取系统主机名
func GetHostName() (string, error) {
	hostname, err := os.Hostname()
	return hostname, err
}
