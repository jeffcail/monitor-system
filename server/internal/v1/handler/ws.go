package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"

	"golang.org/x/crypto/ssh"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type MSsh struct {
	Websocket *websocket.Conn
	Stdin     io.WriteCloser
	Stdout    *wsBufferWriter
	Session   *ssh.Session
}

type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

// RunWebSSH
func RunWebSSH(c echo.Context) error {
	//params := &params2.RunSshServerParams{}
	//_ = c.Bind(params)
	//msg := utils.ValidateParam(params)
	//if msg != "" {
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	//}
	mssh := &MSsh{}
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"service-cloud-monitoring"},
	}
	webConn, err := upGrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		ubzer.MLog.Error("websocket upgrade 失败", zap.Error(err))
	}
	mssh.Websocket = webConn

	sshClient, err := ssh.Dial("tcp", "192.168.0.125:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("112233")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		//ubzer.MLog.Error(fmt.Sprintf("连接 %v 失败, 填写的用户名 %v 和密码 %v", params.Ip, params.Name, params.Password),
		//	zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "连接失败，请检查ip账号和密码是否正确", ""))
	}

	session, err := sshClient.NewSession()
	//defer session.Close()
	if err != nil {
		ubzer.MLog.Error("打开一个新会话失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "开启会话失败", ""))
	}
	mssh.Session = session

	mssh.Stdin, err = session.StdinPipe()
	if err != nil {
		ubzer.MLog.Error("保存用户输入的信息流失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "保存用户输入的信息流失败", ""))
	}

	sshOut := new(wsBufferWriter)
	session.Stdout = sshOut // 会话输出关联到系统标准输出设备
	session.Stderr = sshOut // 会话错误输出关联到系统标准错误输出设备
	mssh.Stdout = sshOut    // 会话输入关联到系统标准输入设备

	modes := ssh.TerminalModes{
		ssh.ECHO:          1, // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err = session.RequestPty("linux", 30, 120, modes); err != nil {
		ubzer.MLog.Error("绑定pty失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "绑定pty失败", ""))
	}

	session.Shell()

	// 执行远程命令
	go Send2SSH(mssh)
	go Send2Web(mssh)

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", ""))
}

// 读取websocket数据，发送到ssh输入流中
func Send2SSH(mssh *MSsh) {
	for {
		_, wsData, err := mssh.Websocket.ReadMessage()
		if err != nil {
			ubzer.MLog.Error("读取websocket数据失败", zap.Error(err))
			return
		}
		_, err = mssh.Stdin.Write(wsData)
		if err != nil {
			ubzer.MLog.Error("ssh发送数据失败", zap.Error(err))
		}
	}
}

// 读取ssh输出，发送到websocket中
func Send2Web(mssh *MSsh) {
	for {
		if mssh.Stdout.buffer.Len() > 0 {
			err := mssh.Websocket.WriteMessage(websocket.TextMessage, mssh.Stdout.buffer.Bytes())
			fmt.Println(string(mssh.Stdout.buffer.Bytes()))
			if err != nil {
				ubzer.MLog.Error("websocket发送数据失败", zap.Error(err))
			}
			mssh.Stdout.buffer.Reset() // 读完清空
		}
	}
}

//func CloseSsh(c echo.Context) error {
//	mssh := &MSsh{}
//	upGrader := websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//		Subprotocols: []string{"service-cloud-monitoring"},
//	}
//	webConn, err := upGrader.Upgrade(c.Response().Writer, c.Request(), nil)
//	if err != nil {
//		ubzer.MLog.Error("websocket upgrade 失败", zap.Error(err))
//	}
//	mssh.Websocket = webConn
//	_, readContent, _ := mssh.Websocket.ReadMessage()
//	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ %v", string(readContent))
//	if string(readContent) == "close" {
//		mssh.Session.Wait()
//		fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 关闭ssh成功")
//	}
//	return nil
//}
