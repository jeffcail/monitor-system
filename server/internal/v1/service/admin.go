package service

import (
	"bz.service.cloud.monitoring/common/ubzer"
	"bz.service.cloud.monitoring/server/config"
	"bz.service.cloud.monitoring/server/internal/v1/daos"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	params "bz.service.cloud.monitoring/server/internal/v1/params"
	"bz.service.cloud.monitoring/server/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func AdminRegister(params *params.AdminParam, url, method string) (bool, error) {
	admin, err := daos.GetAdminInfoByUsername(params.Username)
	if admin != nil {
		ubzer.MLog.Error(fmt.Sprintf(" username: %v 用户已存在", params.Username), zap.Error(err))
		return false, errors.New("用户已存在")
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
		return false, err
	}
	err = daos.RecordOperateLog(vo.Id, vo.Username, vo.RealName, url, method, fmt.Sprintf("管理员账号：%v 在 %v 成功创建",
		vo.Username, vo.CreatedAt))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("账号: %v 创建时记录日志失败", params.Username), zap.Error(err))
	}
	return true, nil
}
