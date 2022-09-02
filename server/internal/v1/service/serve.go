package service

import (
	"errors"
	"fmt"
	"time"

	_const "bz.service.cloud.monitoring/common/const"

	"bz.service.cloud.monitoring/server/internal/v1/models"

	"go.uber.org/zap"

	"bz.service.cloud.monitoring/server/internal/v1/daos"

	"bz.service.cloud.monitoring/common/request"
	"bz.service.cloud.monitoring/common/ubzer"
	params2 "bz.service.cloud.monitoring/server/internal/v1/params"
)

type ServeListView struct {
	Id            int64  `json:"id"`
	ServeName     string `json:"serve_name"`
	ServeAddress  string `json:"serve_address"`
	LastCheckTime string `json:"last_check_time"`
	ServeState    int    `json:"serve_state"`
}

// ServeList
func ServeList(params *params2.ServeListParams, admin *models.MonAdmin, url string, method string) {
	list, serves, err := daos.ServeList(params)
}

var serveState int

// CreateServe
func CreateServe(params *params2.CreateServeParams, admin *models.MonAdmin, path string, method string) error {
	checkTime := time.Now()
	ubzer.MLog.Info(fmt.Sprintf("创建服务开始检测 服务名称: %v 服务地址: %v", params.ServeName, params.ServeAddress))
	res, err := request.Get(params.ServeAddress)
	ubzer.MLog.Info(fmt.Sprintf("创建服务开始检测 服务名称: %v 服务地址: %v 检测结果: %v", params.ServeName,
		params.ServeAddress, string(res)))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("检测服务: %v 地址: %v ping不通", params.ServeName, params.ServeAddress))
		return errors.New("服务错误，联系技术排查")
	}
	if string(res) != "pong" {
		serveState = 2
	} else {
		serveState = 1
	}
	err = daos.CreateServe(params, serveState, checkTime)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("创建服务检测失败, 名称: %v 地址: %v", params.ServeName, params.ServeAddress),
			zap.Error(err))
		return errors.New("创建服务检测失败")
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, path, method, fmt.Sprintf("%v 在 %v 创建了名称为 %v的服务",
		admin.Username, time.Now().Format(_const.Layout), params.ServeName))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录 %v 创建服务检测日志失败,服务名称: %v", admin.Username, params.ServeName))
	}

	return nil
}

// DeleteServe
func DeleteServe(params *params2.DeleteServeParams, admin *models.MonAdmin, url string, method string) error {
	name, err := daos.DeleteServe(params.ID)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 删除服务失败 服务id: %d", admin.Username, params.ID))
		return errors.New("删除服务失败")
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 删除了名称为 %v 的服务",
		admin.Username, time.Now().Format(_const.Layout), name))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录 %v 删除服务检测日志失败,服务名称: %v", admin.Username, name))
	}
	return nil
}
