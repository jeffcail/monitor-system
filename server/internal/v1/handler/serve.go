package handler

import (
	"fmt"
	"net/http"
	"time"

	"bz.service.cloud.monitoring/common/request"
	"bz.service.cloud.monitoring/common/ubzer"
	"go.uber.org/zap"

	"github.com/spf13/cast"

	"bz.service.cloud.monitoring/server/internal/v1/daos"
	"bz.service.cloud.monitoring/server/internal/v1/service"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/params"
	"github.com/labstack/echo"
)

// ServeList
func ServeList(c echo.Context) error {
	params := &params.ServeListParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	count, data, err := service.ServeList(params, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "查看列表失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功",
		utils.Resp.ResponsePagination(count, data)))
}

// CreateServe
func CreateServe(c echo.Context) error {
	serveParams := &params.CreateServeParams{}
	_ = c.Bind(serveParams)
	msg := utils.ValidateParam(serveParams)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.CreateServe(serveParams, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "创建成功", ""))
}

// DeleteServe
func DeleteServe(c echo.Context) error {
	params := &params.DeleteServeParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.DeleteServe(params, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "删除成功", ""))
}

// UpgradeServe
func UpgradeServe(c echo.Context) error {
	param := &params.UpgradeServeParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	//lastIndex := strings.LastIndex(param.ServeAddress, ":")
	//s := param.ServeAddress[:lastIndex]
	//index := strings.LastIndex(s, ":")
	//s2 := s[index+3:]

	header := make(map[string]string)
	p := make(map[string]interface{})
	p["package_name"] = param.PackageName
	p["package_path"] = param.PackagePath
	res, err := request.GetParams("http://"+param.ServeIp+":9093/c/serve/upgrade", header, p)
	ubzer.MLog.Info(fmt.Sprintf("升级=========== res: %v", string(res)))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("升级服务请求失败 ServeAddress: %v", param.ServeIp), zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "升级服务请求失败", ""))
	}
	type FRes struct {
		str string
	}

	//if f.str != "success" {
	//	ubzer.MLog.Error(fmt.Sprintf("升级服务失败 ServeAddress: %v", param.ServeIp), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, string(res), ""))
	//}
	admin := GetAdminInfoFromParseToken(c)
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, c.Request().URL.Path, c.Request().Method,
		fmt.Sprintf("%v 在 %v 时间升级了服务地址为 %v 的服务", admin.Username, time.Now(), param.ServeIp))
	if err != nil {
		ubzer.MLog.Error("记录操作日志失败", zap.Error(err))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, string(res), ""))
}
