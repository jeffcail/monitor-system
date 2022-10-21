package handler

import (
	"fmt"
	"net/http"

	"github.com/c/monitor-system/common/ubzer"
	"go.uber.org/zap"

	"github.com/spf13/cast"

	_const "github.com/c/monitor-system/common/const"
	"github.com/c/monitor-system/common/utils"
	"github.com/c/monitor-system/server/connections"
	"github.com/labstack/echo"
)

// ShellWeb
func ShellWeb(c echo.Context) error {
	var err error

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error("websocket upgrade 失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "websocket upgrade 失败", ""))
	}
	_, readContent, err := conn.ReadMessage()
	if err != nil {
		ubzer.MLog.Error("websocket 读取ip、用户名、密码 失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ readContent: %v\n", string(readContent))

	sshClient, err := connections.DecodeMsgToSSHClient(string(readContent))
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ sshClient: %v\n", sshClient)

	terminal := connections.Terminal{
		Columns: 150,
		Rows:    35,
	}

	var port = 22
	err = sshClient.GenerateClient(sshClient.IpAddress, sshClient.Username, sshClient.Password, port)
	if err != nil {
		conn.WriteMessage(1, []byte(err.Error()))
		conn.Close()
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	sshClient.RequestTerminal(terminal)
	sshClient.Connect(conn)
	return nil
}
