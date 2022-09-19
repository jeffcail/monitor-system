package service

import (
	"errors"
	"fmt"
	"time"

	_const "bz.service.cloud.monitoring/common/const"
	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/config"
	"bz.service.cloud.monitoring/server/internal/v1/daos"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	params "bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/utils"
	"go.uber.org/zap"
)

func AdminRegister(params *params.AdminParam, admin *models.MonAdmin, url, method string) error {
	res, err := daos.GetAdminInfoByUsername(params.Username)
	if res != nil {
		ubzer.MLog.Error(fmt.Sprintf(" username: %v 用户已存在", params.Username), zap.Error(err))
		return errors.New("用户已存在")
	}
	vo := &models.MonAdmin{
		Username:   params.Username,
		RealName:   params.RealName,
		Password:   utils.GeneratePassword(params.Password, config.Config().Slat),
		Phone:      params.Phone,
		RoleId:     "1",
		Department: params.Department,
		OfficePost: params.OfficePost,
		State:      1,
	}
	err = daos.AddAdmin(vo)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("账号: %v 创建失败", params.Username), zap.Error(err))
		return err
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("管理员账号：%v 在 %v 成功创建",
		vo.Username, vo.CreatedAt))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("账号: %v 创建时记录日志失败", params.Username), zap.Error(err))
	}
	return nil
}

// 查询管理员信息所需要返回的字段
type SelAdminList struct {
	Id            int64  `json:"id"`
	Username      string `json:"username" `
	RealName      string `json:"real_name" `
	Phone         string `json:"phone" `
	RoleId        string `json:"role_id" `
	Department    string `json:"department" `
	OfficePost    string `json:"office_post" `
	State         int    `json:"state" `
	LastLoginTime string `json:"last_login_time" `
	CreatedAt     string `json:"created_at" `
	UpdatedAt     string `json:"updated_at" `
	Version       int64  `json:"version" `
}

// SelAdmin
func SelAdmin(params *params.SelAdminParam, admin *models.MonAdmin, url string, method string) (int64, []*SelAdminList, error) {
	t := time.Now().Format(_const.Layout)
	count, list1, err := daos.SelAdmin(params, filterSelAdmin(params))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 查看管理员信息列表时获取数据失败", admin.Username, t))
		return 0, nil, errors.New("管理员信息列表获取失败")
	}
	list := make([]*SelAdminList, 0)
	for _, bean := range list1 {
		date := &SelAdminList{
			Id:            bean.Id,
			Username:      bean.Username,
			RealName:      bean.RealName,
			Phone:         bean.Phone,
			RoleId:        bean.RoleId,
			Department:    bean.Department,
			OfficePost:    bean.OfficePost,
			State:         bean.State,
			LastLoginTime: bean.LastLoginTime.Format("2006-01-02 15:04:05"), //转为string再返回给前端
			CreatedAt:     bean.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     bean.UpdatedAt.Format("2006-01-02 15:04:05"),
			Version:       bean.Version,
		}
		list = append(list, date)
	}
	//err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 查看了管理员信息列表",
	//	admin.Username, t))
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("记录 %v 在 %v 查看管理员信息列表日志失败", admin.Username, t))
	//}

	return count, list, nil
}

func filterSelAdmin(params *params.SelAdminParam) map[string]interface{} {

	m := make(map[string]interface{})
	if params.Username != "" {
		m["username"] = params.Username
	}
	if params.RealName != "" {
		m["real_name"] = params.RealName
	}
	if params.Phone != "" {
		m["phone"] = params.Phone
	}
	if params.Department != "" {
		m["department"] = params.Department
	}
	if params.OfficePost != "" {
		m["office_post"] = params.OfficePost
	}
	return m
}

// PasswordValidate
func PasswordValidate(params *params.UpdAdminParamById) string {
	if params.Password != "" {
		if len(params.Password) < 6 || len(params.Password) > 12 {
			return "Password 字段要求不能小于6位或大于12位"
		}

	}

	return ""
}

// UpdateAdminById
func UpdAdminById(params *params.UpdAdminParamById, admin *models.MonAdmin, url string, method string) error {
	t := time.Now().Format(_const.Layout)
	data, err := daos.GetAdminInfoById(params.Id)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 变更Id:%v 管理员信息时获取数据失败", admin.Username, t, params.Id))
		return err
	}
	data = &models.MonAdmin{
		Id:            params.Id,
		Username:      params.Username,
		RealName:      params.RealName,
		Phone:         params.Phone,
		Password:      utils.GeneratePassword(params.Password, config.Config().Slat),
		RoleId:        params.RoleId,
		Department:    params.Department,
		OfficePost:    params.OfficePost,
		State:         params.State,
		UpdatedAt:     time.Now(),
		LastLoginTime: data.LastLoginTime,
		CreatedAt:     data.CreatedAt,
		Version:       data.Version,
	}

	err = daos.UpdAdmin(data)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 变更Id:%v 管理员信息失败", admin.Username, t, params.Id))
		return err
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 修改了ID:%v管理员信息列表",
		admin.Username, t, params.Id))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录 %v 在 %v 变更管理员信息日志失败", admin.Username, t))
	}
	return nil
}
func DeleteAdminById(params *params.DeleteParam, admin *models.MonAdmin, url string, method string) error {
	t := time.Now().Format(_const.Layout)
	err := daos.DeleteAdminById(params.Id)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 删除Id:%v 管理员信息记录失败", admin.Username, t, params.Id))
		return err
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 删除了ID:%v管理员信息记录",
		admin.Username, t, params.Id))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录 %v 在 %v 删除管理员信息日志失败", admin.Username, t))
	}
	return nil
}

// EnableDisableAdmin
func EnableDisableAdmin(id int64, state int, admin *models.MonAdmin, url string, method string) error {
	t := time.Now().Format(_const.Layout)
	a, err := daos.GetAdminInfoById(id)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 启用｜禁用管理员失败", admin.Username, t), zap.Error(err))
		return err
	}
	a.State = state
	err = daos.UpdAdmin(a)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 启用｜禁用管理员失败", admin.Username, t), zap.Error(err))
		return err
	}

	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 启用｜禁用了管理员 :%v",
		admin.Username, t, a.Username))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("记录 %v 在 %v 启用｜禁用 管理员 记录日志失败", admin.Username, t))
	}

	return nil
}
