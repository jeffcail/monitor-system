package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"

	"github.com/spf13/cast"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/common/ubzer"

	"bz.service.cloud.monitoring/server/config"
	"bz.service.cloud.monitoring/server/internal/v1/params"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"github.com/labstack/echo"
)

// MachineList
func MachineList(c echo.Context) error {
	params := &params.MachineListParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	count, list, err := service.MachineList(params, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取机器列表失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功",
		utils.Resp.ResponsePagination(count, list)))
}

// AllMachine
func AllMachine(c echo.Context) error {
	machines, err := service.AllMachine()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "获取所有客户端机器失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", machines))
}

// SendCom
func SendCom(c echo.Context) error {
	param := &params.SendComParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	//h := make(map[string]string)
	//p := make(map[string]interface{})
	//p["content"] = param.Content
	//bytes, err := request.GetParams("http://"+param.Ip+":9093/c/machine/receive/com", h, p)
	//if err != nil {
	//	ubzer.MLog.Error("向客户端发送指令请求异常", zap.Error(err))
	//}

	// 原http 升级为websocket
	dl := websocket.Dialer{}
	d := "ws://" + param.Ip + config.Config().ClientHttpBind + "/c/machine/receive/com"
	conn, _, err := dl.Dial(d, nil)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("向客户端服务: %v 发送指令, websocket连接客户端失败", param.Ip), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte(param.Content))
	if err != nil {
		ubzer.MLog.Error("往客户端服务推送系统指令失败", zap.Error(err))
	}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "指令发送成功", ""))
}

// UpdateMachineRemark
func UpdateMachineRemark(c echo.Context) error {
	param := &params.UpdateMachineRemarkParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.UpdateMachineRemark(param, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "修改成功", ""))
}

// UpgradeClientMachine
func UpgradeClientMachine(c echo.Context) error {
	param := &params.UpgradeClientMachineParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	if strings.Contains(param.UpgradeVersion, ".") {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "版本号只能为整数版本", ""))
	}

	err := service.CheckClientUpgradeVersion(param.MachineIp, param.UpgradeVersion)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	file, err := c.FormFile("go_file")
	if err != nil {
		ubzer.MLog.Error("文件接收失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "文件接收失败", ""))
	}

	src, err := fileClientCheck(file)
	defer src.Close()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	packageName := param.MachineIp + ":" + file.Filename + ":" + param.UpgradeVersion
	err = checkClientFileIsExists(packageName)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	dst, err := os.Create(config.Config().UpClientPkgPath + packageName)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	md5Sum := GetFileMd5Sum(config.Config().UpClientPkgPath + packageName)

	// 保留最近三份历史包记录
	exec := "rm -f " + config.Config().UpClientPkgPath + param.MachineIp + ":" + file.Filename + ":" + cast.ToString(cast.ToInt64(param.UpgradeVersion)-3)
	err = utils.ExecCommand(exec)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("升级客户端程序 删除历史包失败, 当前版本为: %v", param.UpgradeVersion), zap.Error(err))
	}

	err = service.UpgradeClientServe(param, file.Filename, md5Sum, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	//dl := websocket.Dialer{}
	//d := "ws://" + param.MachineIp + config.Config().ClientHttpBind + "/c/machine/upgrade/client"
	//conn, _, err := dl.Dial(d, nil)
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("升级服务器: %v 的客户端, websocket连接客户端失败", param.MachineIp), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	//}
	//
	//err = conn.WriteMessage(websocket.TextMessage, []byte("ok"))
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("升级服务器: %v 的客户端,发送升级指令失败", param.MachineIp), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "失败", ""))
	//}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "上传成功，60秒内自动升级此服务", ""))
}

// GetFileMd5Sum
func GetFileMd5Sum(file string) string {
	ubzer.MLog.Info(fmt.Sprintf("========== file: %v", file))
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	ctx := make([]byte, 8)
	md5sum := md5.New()
	num, err := f.Read(ctx)
	if err == io.EOF {
		return ""
	}
	md5sum.Write(ctx[:num])
	return fmt.Sprintf("%x", md5sum.Sum(nil))
}

// fileClientCheck
func fileClientCheck(file *multipart.FileHeader) (multipart.File, error) {
	// 字节 100 * 1024 * 1024
	if file.Size > 1024*1024*100 {
		return nil, errors.New("文件过大,不能超过100M")
	}

	if strings.Contains(file.Filename, ".") {
		return nil, errors.New("非法文件,文件不允许带后缀")
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	return src, nil
}

// checkClientFileIsExists
func checkClientFileIsExists(packageName string) error {
	_, err := os.Stat(config.Config().UpClientPkgPath)
	if err != nil {
		err = os.Mkdir(config.Config().UpClientPkgPath, 0777)
		return err
	}
	_, err = os.Stat(config.Config().UpClientPkgPath + packageName)
	if !os.IsNotExist(err) {
		return errors.New("已经存在此版本发布的包，请检查输入的版本号，或重新输入版本号")
	}
	return nil
}

// UpgradeClientMachineRecord
func UpgradeClientMachineRecord(c echo.Context) error {
	param := &params.UpgradeClientMachineRecord{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(msg), ""))
	}

	count, lists, err := service.UpgradeClientMachineRecord(param)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功",
		utils.Resp.ResponsePagination(count, lists)))
}

// Download
func Download(c echo.Context) error {
	file := c.QueryParam("file")
	return c.Attachment(config.Config().UpClientPkgPath+file, file)
}

//// DeleteMachine
//func DeleteMachine(c echo.Context) error {
//
//}
