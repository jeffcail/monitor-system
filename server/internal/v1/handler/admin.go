package handler

import (
	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"github.com/labstack/echo"
	"net/http"
)

// controller层
// 管理员创建
func AdminRegister(e echo.Context) error {
	//实例化一个对象接受信息
	params := &params.AdminParam{}
	_ = e.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	res, _ := service.AdminRegister(params, GetAdminInfoFromParseToken(e), e.Request().URL.Path, e.Request().Method)
	if !res {

		return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Fail, "用户创建失败", ""))
	}
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "管理员用户创建成功！", ""))
}

// 管理员信息列表查询
func SelectAdmin(e echo.Context) error {
	// 接参
	params := &params.SelAdminParam{}
	_ = e.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	// 请求service
	count, list, err := service.SelAdmin(params, GetAdminInfoFromParseToken(e), e.Request().URL.Path, e.Request().Method)
	if err != nil {
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "查询失败", ""))
	}
	// 返回
	data := utils.Resp.ResponsePagination(count, list) //拼接
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功！", data))
}

// 管理员信息变更
func UpdateAdminById(e echo.Context) error {
	// 接参
	params := &params.UpdAdminParamById{}
	//参数验证
	_ = e.Bind(params)
	msg := utils.ValidateParam(params)
	msg += service.PasswordValidate(params)
	if msg != "" {
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	//请求service
	_, err := service.UpdAdminById(params, GetAdminInfoFromParseToken(e), e.Request().URL.Path, e.Request().Method)
	//返回结果处理
	if err != nil {
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "变更失败", ""))
	}
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "变更成功！", ""))
}
