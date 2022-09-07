package handler

import (
	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/utils"
	"bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/internal/v1/service"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// log query
func LogQuery(e echo.Context) error {
	//接参
	params := &params.LogQueryParam{}
	//验证
	_ = e.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	//请求service
	count, list, err := service.LogQueryList(params, GetAdminInfoFromParseToken(e))
	if err != nil {
		fmt.Printf("失败原因：%v\n", err)
		return e.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "查询失败", ""))
	}
	//返回
	data := utils.Resp.ResponsePagination(count, list) //拼接
	return e.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功！", data))
}
