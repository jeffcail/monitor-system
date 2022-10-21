package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cast"

	_const "github.com/c/monitor-system/common/const"
	"github.com/c/monitor-system/common/utils"
	params2 "github.com/c/monitor-system/server/internal/v1/params"

	"github.com/c/monitor-system/common/db"

	"github.com/gorilla/websocket"

	"github.com/c/monitor-system/server/internal/v1/daos"

	"go.uber.org/zap"

	"github.com/c/monitor-system/common/ubzer"
	"github.com/labstack/echo"
)

type ServerClient struct {
	Conn *websocket.Conn
	Send chan *Message
}

type Message struct {
	Exists int `json:"exists"`
}

var (
	Ser      *ServerClient
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (s *ServerClient) write(exist int) {
	m := &Message{Exists: exist}
	res, _ := json.Marshal(m)
	err := s.Conn.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error("往客户端推送系统信息失败", zap.Error(err))
	}
}

// InitClient
func InitClient(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}
	Ser = &ServerClient{
		Conn: conn,
		Send: make(chan *Message, 100),
	}

	// 从客户端读取数据
	type ClientMessage struct {
		Ip          string `json:"ip"`
		HostName    string `json:"host_name"`
		MachineCode string `json:"machine_code"`
	}
	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error("获取客户端初始化推送过来的数据失败", zap.Error(err))
	}
	cm := &ClientMessage{}
	_ = json.Unmarshal(content, cm)
	fmt.Printf("======================cm: %v", cm)

	// 业务逻辑，检测缓存和数据库中是否存在客户端的ip和机器码
	// 从缓存中拿机器码
	getMachineCode, err := daos.GetMachineCode(cm.Ip)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("从缓存中获取唯一机器码失败"), zap.Error(err))
	}
	if getMachineCode == "" {
		// 从数据库中拿
		m, err := daos.GetMachineInfoFromDbByIp(cm.Ip)
		if err != nil {
			ubzer.MLog.Error(fmt.Sprintf("从数据库中获取唯一机器码失败"), zap.Error(err))
		}
		if m.Id == 0 {
			// 数据库中不存在 第一次部署
			err = daos.SetMachineCode(cm.Ip, cm.MachineCode)
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("生成唯一机器码写入缓存失败 code: %v", cm.MachineCode), zap.Error(err))
			}
			err = daos.SaveHostName(cm.HostName, cm.MachineCode, cm.Ip)
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("生成唯一机器码之后保存服务器信息失败 ip: %v code: %v name: %v",
					cm.Ip, cm.MachineCode, cm.HostName), zap.Error(err))
			}
			go Ser.write(2)
			return nil
		}

	}

	go Ser.write(1)

	return nil
}

// ClientCpuPercent
func ClientCpuPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	result, err := db.Rc.LRange(params.Ip+"-cpu", 0, 120).Result()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", result))
}

// ClientMemPercent
func ClientMemPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	result, err := db.Rc.LRange(params.Ip+"-cpu", 0, 120).Result()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", result))
}

// ClientDiskPercent
func ClientDiskPercent(c echo.Context) error {
	params := &params2.ClientPercentParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	result, err := db.Rc.LRange(params.Ip+"-cpu", 0, 120).Result()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", result))
}

// ------------------------------------------------  废弃    ----------------------------------------------

//// ClientCpuPercent
//func ClientCpuPercent(c echo.Context) error {
//	params := &params2.ClientPercentParams{}
//	_ = c.Bind(params)
//	msg := utils.ValidateParam(params)
//	if msg != "" {
//		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
//	}
//
//	// 由原来的http 升级为websocket
//	dl := websocket.Dialer{}
//	d := "ws://" + params.Ip + config.Config().ClientHttpBind + "/c/sys/cpu"
//	conn, _, err := dl.Dial(d, nil)
//
//	if err != nil {
//		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的cpu使用率, websocket连接客户端失败", params.Ip), zap.Error(err))
//		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
//	}
//
//	type Result struct {
//		Percent string
//	}
//
//	_, content, err := conn.ReadMessage()
//	if err != nil {
//		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的cpu使用率, 从websocket中读取客户端推送信息失败", params.Ip),
//			zap.Error(err))
//	}
//	res := &Result{}
//	_ = json.Unmarshal(content, res)
//	if res.Percent != "" {
//		if utils.LessThenAndEqual("60", res.Percent) {
//			err := daos.CreateMachineCheckRecord(params.Ip, "cpu", cast.ToInt(res.Percent))
//			if err != nil {
//				ubzer.MLog.Error(fmt.Sprintf("记录服务器报警信息失败"))
//			}
//		}
//	}
//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", cast.ToInt(res.Percent)))
//}
//
//// ClientMemPercent
//func ClientMemPercent(c echo.Context) error {
//	params := &params2.ClientPercentParams{}
//	_ = c.Bind(params)
//	msg := utils.ValidateParam(params)
//	if msg != "" {
//		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
//	}
//
//	// 由原来的http 升级为websocket
//	dl := websocket.Dialer{}
//	d := "ws://" + params.Ip + config.Config().ClientHttpBind + "/c/sys/mem"
//	conn, _, err := dl.Dial(d, nil)
//
//	if err != nil {
//		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的内存使用率, websocket连接客户端失败", params.Ip), zap.Error(err))
//		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
//	}
//	type Result struct {
//		Percent string
//	}
//	defer conn.Close()
//
//	_, content, err := conn.ReadMessage()
//	if err != nil {
//		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的内存使用率, 从websocket中读取客户端推送信息失败", params.Ip),
//			zap.Error(err))
//	}
//	res := &Result{}
//	_ = json.Unmarshal(content, res)
//	if res.Percent != "" {
//		if utils.LessThenAndEqual("60", res.Percent) {
//			err := daos.CreateMachineCheckRecord(params.Ip, "mem", cast.ToInt(res.Percent))
//			if err != nil {
//				ubzer.MLog.Error(fmt.Sprintf("记录服务器报警信息失败"))
//			}
//		}
//	}
//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", cast.ToInt(res.Percent)))
//}
//
//// ClientDiskPercent
//func ClientDiskPercent(c echo.Context) error {
//	params := &params2.ClientPercentParams{}
//	_ = c.Bind(params)
//	msg := utils.ValidateParam(params)
//	if msg != "" {
//		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
//	}
//
//	// 由原来的http 升级为websocket
//	dl := websocket.Dialer{}
//	d := "ws://" + params.Ip + config.Config().ClientHttpBind + "/c/sys/disk"
//	conn, _, err := dl.Dial(d, nil)
//
//	if err != nil {
//		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的磁盘使用率, websocket连接客户端失败", params.Ip), zap.Error(err))
//		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
//	}
//	type Result struct {
//		Percent string
//	}
//
//	_, content, err := conn.ReadMessage()
//	if err != nil {
//		ubzer.MLog.Error(fmt.Sprintf("获取服务器: %v 的磁盘使用率, 从websocket中读取客户端推送信息失败", params.Ip),
//			zap.Error(err))
//	}
//	res := &Result{}
//	_ = json.Unmarshal(content, res)
//	if res.Percent != "" {
//		if utils.LessThenAndEqual("60", res.Percent) {
//			err := daos.CreateMachineCheckRecord(params.Ip, "cpu", cast.ToInt(res.Percent))
//			if err != nil {
//				ubzer.MLog.Error(fmt.Sprintf("记录服务器报警信息失败"))
//			}
//		}
//	}
//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", cast.ToInt(res.Percent)))
//}

// ------------------------------------------------  废弃    ----------------------------------------------
