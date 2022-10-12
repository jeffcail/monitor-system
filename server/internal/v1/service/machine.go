package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/c/server-monitoring/server/internal/v1/models"

	params2 "github.com/c/server-monitoring/server/internal/v1/params"

	_const "github.com/c/server-monitoring/common/const"

	"go.uber.org/zap"

	"github.com/c/server-monitoring/common/ubzer"
	"github.com/c/server-monitoring/server/internal/v1/daos"
)

type MachineListResult struct {
	ID          int64  `json:"id"`
	MachineCode string `json:"machine_code"`
	Ip          string `json:"ip"`
	HostName    string `json:"host_name"`
	Remark      string `json:"remark"`
	CreatedAt   string `json:"created_at"`
}

// MachineList
func MachineList(params *params2.MachineListParams, admin *models.MonAdmin, url string, method string) (int64, []*MachineListResult, error) {
	count, list, err := daos.MachineList(params)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取机器列表失败"), zap.Error(err))
		return 0, nil, errors.New("获取机器列表失败")
	}

	data := make([]*MachineListResult, 0)

	for _, v := range list {
		mlr := &MachineListResult{}
		mlr.ID = v.Id
		mlr.MachineCode = v.MachineCode
		mlr.Ip = v.Ip
		mlr.HostName = v.Hostname
		mlr.Remark = v.Remark
		mlr.CreatedAt = v.CreatedAt.Format(_const.Layout)
		data = append(data, mlr)
	}

	//err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf(""+
	//	"%v 在 %v 查看了机器列表", admin.Username, time.Now().Format(_const.Layout)))
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 查看了机器列表数据，记录此操作失败", admin.Username, time.Now().Format(_const.Layout)))
	//}

	return count, data, nil
}

type AllMachineRes struct {
	Id          int64  `json:"id"`
	MachineCode string `json:"machine_code"`
	Ip          string `json:"ip"`
	Content     string `json:"content"`
}

// AllMachine
func AllMachine() ([]*AllMachineRes, error) {
	data := make([]*AllMachineRes, 0)
	machine, err := daos.AllMachine()
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("获取所有服务器失败"), zap.Error(err))
		return nil, err
	}

	for _, v := range machine {
		amr := &AllMachineRes{}
		amr.Id = v.Id
		amr.MachineCode = v.MachineCode
		amr.Ip = v.Ip
		if v.Remark != "" {
			amr.Content = v.Remark
		} else {
			amr.Content = v.Ip + " | " + v.Hostname
		}
		data = append(data, amr)
	}

	return data, nil
}

// UpdateMachineRemark
func UpdateMachineRemark(param *params2.UpdateMachineRemarkParams, admin *models.MonAdmin, url string, method string) error {
	err := daos.UpdateMachineRemark(param)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 时间 修改机器: %v 备注失败", admin.Username, time.Now().Format(_const.Layout),
			param.MachineCode), zap.Error(err))
		return err
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method,
		fmt.Sprintf("%v 在 %v 修改机器: %v 备注: %v", admin.Username, time.Now().Format(_const.Layout),
			param.MachineCode, param.Remark))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 修改机器: %v 备注记录此操作失败", admin.Username, time.Now().Format(_const.Layout),
			param.MachineCode), zap.Error(err))
	}
	return nil
}

// UpgradeClientServe
func UpgradeClientServe(param *params2.UpgradeClientMachineParams, packageName string, md5Sum string, admin *models.MonAdmin, url string, method string) error {
	t := time.Now().Format(_const.Layout)
	err := daos.AddUpgradeClientServeRecord(param, packageName, md5Sum)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 升级了机器为: %v 的服务 上传的包名: %v", admin.Username, t, param.MachineIp,
			packageName))
		return err
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 升级了"+
		"机器为: %v 的服务 上传的包名: %v", admin.Username, t, param.MachineIp,
		packageName))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 升级了机器为: %v 的服务 上传的包名: %v 记录此日志失败", admin.Username, t,
			param.MachineIp, packageName))
	}
	return nil
}

// CheckClientUpgradeVersion
func CheckClientUpgradeVersion(machineIp string, upgradeVersion string) error {
	m, err := daos.CheckClientUpgradeVersion(machineIp)
	if err != nil {
		return err
	}

	if m != nil {
		if m.UpgradeVersion == upgradeVersion {
			return errors.New("版本重复")
		}
		if m.UpgradeVersion > upgradeVersion {
			return errors.New("当前版本号不得小于历史版本号")
		}
	}
	return nil
}

type UpgradeClientRecordList struct {
	Id              int64  `json:"id"`
	MachineCode     string `json:"machine_code"`
	MachineIp       string `json:"machine_ip"`
	MachineHostname string `json:"machine_hostname"`
	MachineRemark   string `json:"machine_remark"`
	PackageName     string `json:"package_name"`
	UpgradeVersion  string `json:"upgrade_version"`
	CreatedAt       string `json:"created_at"`
}

// UpgradeClientMachineRecord
func UpgradeClientMachineRecord(param *params2.UpgradeClientMachineRecord) (int64, []*UpgradeClientRecordList, error) {
	data := make([]*UpgradeClientRecordList, 0)
	count, records, err := daos.UpgradeClientMachineRecord(param)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("查看客户端升级记录失败 服务地址: %v", param.MachineIp), zap.Error(err))
		return 0, nil, errors.New("查看客户端升级记录失败")
	}

	for _, v := range records {
		ucrl := &UpgradeClientRecordList{
			Id:              v.Id,
			MachineCode:     v.MachineCode,
			MachineIp:       v.MachineIp,
			MachineHostname: v.MachineHostname,
			MachineRemark:   v.MachineRemark,
			PackageName:     v.PackageName,
			UpgradeVersion:  v.UpgradeVersion,
			CreatedAt:       v.CreatedAt.Format(_const.Layout),
		}
		data = append(data, ucrl)
	}
	return count, data, err
}

//// DeleteMachine
//func DeleteMachine() {
//
//}
