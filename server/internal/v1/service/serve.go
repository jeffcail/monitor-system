package service

import (
	"errors"
	"fmt"
	"time"

	_const "github.com/c/server-monitoring/common/const"

	"github.com/c/server-monitoring/server/internal/v1/models"

	"go.uber.org/zap"

	"github.com/c/server-monitoring/server/internal/v1/daos"

	"github.com/c/server-monitoring/common/request"
	"github.com/c/server-monitoring/common/ubzer"
	params2 "github.com/c/server-monitoring/server/internal/v1/params"
)

type ServeListView struct {
	Id            int64  `json:"id"`
	ServeName     string `json:"serve_name"`
	ServeAddress  string `json:"serve_address"`
	LastCheckTime string `json:"last_check_time"`
	ServeState    int    `json:"serve_state"`
}

// ServeList
func ServeList(params *params2.ServeListParams, admin *models.MonAdmin, url string, method string) (int64, []*ServeListView, error) {
	//t := time.Now().Format(_const.Layout)
	count, serves, err := daos.ServeList(params)
	data := make([]*ServeListView, 0)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 查看服务检测列表获取数据失败", admin.Username, time.Now().Format(_const.Layout)))
		return 0, nil, errors.New("服务检测列表获取失败")
	}

	for _, v := range serves {
		slv := &ServeListView{}
		slv.Id = v.Id
		slv.ServeName = v.ServeName
		slv.ServeAddress = v.ServeAddress
		slv.LastCheckTime = v.LastCheckTime.Format(_const.Layout)
		slv.ServeState = v.ServeState

		data = append(data, slv)
	}

	//err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 查看了服务检测列表",
	//	admin.Username, t))
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("记录 %v 在 %v 查看服务检测列表日志失败", admin.Username, t))
	//}

	return count, data, nil
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

// UpgradeServe
func UpgradeServe(param *params2.UpgradeServeParams, packageName string, admin *models.MonAdmin, url string, method string) error {
	t := time.Now().Format(_const.Layout)
	err := daos.AddUpgradeServeRecord(param, packageName)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 升级了名字为: %v 的服务 上传的包名: %v", admin.Username, t, param.ServeName,
			packageName))
		return err
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 升级了名字为: %v 的服务 上传的包名: %v", admin.Username, t, param.ServeName,
		packageName))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 升级了名字为: %v 的服务 上传的包名: %v 记录此日志失败", admin.Username, t, param.ServeName,
			packageName))
	}
	return nil
}

// CheckUpgradeVersion
func CheckUpgradeVersion(serveAddress string, serveUpgradeVersion string) error {
	m, err := daos.CheckUpgradeVersion(serveAddress)
	if err != nil {
		return err
	}

	if m != nil {
		if m.UpgradeVersion == serveUpgradeVersion {
			return errors.New("版本重复")
		}
		if m.UpgradeVersion > serveUpgradeVersion {
			return errors.New("当前版本号不得小于历史版本号")
		}
	}
	return nil
}

type UpgradeRecordList struct {
	Id             int64  `json:"id"`
	ServeName      string `json:"serve_name"`
	ServeAddress   string `json:"serve_address"`
	ServeId        int64  `json:"serve_id"`
	CreatedAt      string `json:"created_at"`
	UpgradeVersion string `json:"upgrade_version"`
	PackName       string `json:"pack_name"`
}

// UpgradeRecord
func UpgradeRecord(param *params2.UpgradeRecord) (int64, []*UpgradeRecordList, error) {
	data := make([]*UpgradeRecordList, 0)
	count, records, err := daos.UpgradeRecord(param)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("查看服务升级记录失败 服务地址: %v", param.ServeAddress), zap.Error(err))
		return 0, nil, errors.New("查看服务升级记录失败")
	}

	for _, v := range records {
		u := &UpgradeRecordList{}
		u.Id = v.Id
		u.ServeName = v.ServeName
		u.ServeAddress = v.ServeAddress
		u.ServeId = v.ServeId
		u.CreatedAt = v.CreatedAt.Format(_const.Layout)
		u.UpgradeVersion = v.UpgradeVersion
		u.PackName = v.PackageName
		data = append(data, u)
	}

	return count, data, nil
}
