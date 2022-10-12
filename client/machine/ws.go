package machine

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/c/server-monitoring/common/ubzer"

	"github.com/c/server-monitoring/common/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type Client struct {
	Ip   string
	Conn *websocket.Conn
	Send chan *SystemMessage
}

type SystemMessage struct {
	Percent string `json:"percent"` // 使用率 cpu memory disk
}

var (
	WsClient *Client
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (c *Client) write(percent string) {
	m := &SystemMessage{
		Percent: percent,
	}
	res, _ := json.Marshal(m)
	err := c.Conn.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		ubzer.MLog.Error("往服务器推送系统信息失败", zap.Error(err))
	}
}

var exec string

// WsReceiveCom
func WsReceiveCom(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("websocket 服务升级失败"), zap.Error(err))
		return err
	}
	_, content, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("从 websocket中读取服务端发送的指令失败"), zap.Error(err))
	}
	exec = string(content)
	err = utils.ExecCommand(exec)
	if err != nil {
		ubzer.MLog.Error("指令执行失败", zap.Error(err))
	}
	return nil
}
