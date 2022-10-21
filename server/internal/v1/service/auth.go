package service

import (
	"errors"
	"fmt"
	"time"

	_const "github.com/c/monitor-system/common/const"

	"github.com/c/monitor-system/server/config"
	"github.com/c/monitor-system/server/utils"

	jwt2 "github.com/c/monitor-system/common/jwt"
	"github.com/c/monitor-system/common/ubzer"
	"github.com/c/monitor-system/server/internal/v1/daos"
	params2 "github.com/c/monitor-system/server/internal/v1/params"
	"go.uber.org/zap"
)

// AdminLogin
func AdminLogin(params *params2.AdminLoginParams, url, method string) (string, error) {
	admin, err := daos.GetAdminInfoByUsername(params.Username)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("查询用户是否存在错误 username: %v", params.Username), zap.Error(err))
		return "", errors.New("用户不存在或者查询用户错误")
	}

	if admin.State == _const.AdminStateOff {
		ubzer.MLog.Info(fmt.Sprintf("账号: %v 已被禁用", params.Username))
		return "", errors.New("账号已被禁用")
	}

	if ok := utils.ComparePassword(admin.Password, params.Password, config.Config().Slat); !ok {
		ubzer.MLog.Info(fmt.Sprintf("账号: %v 密码: %v 错误", params.Username, params.Password))
		return "", errors.New("账号密码不正确")
	}

	cl := &jwt2.JwtClaims{
		ID:       admin.Id,
		Username: admin.Username,
		RealName: admin.RealName,
	}
	token, err := jwt2.GenerateToken(cl)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("签发token失败 账号: %v 密码: %v", params.Username, params.Password), zap.Error(err))
		return "", errors.New("签发token失败")
	}
	lastLoginTime := time.Now()
	admin.LastLoginTime = lastLoginTime
	err = daos.UpdateAdminById(admin)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("账号: %v 登陆签发Token成功, 修改其登陆时间出现异常", params.Username), zap.Error(err))
	}
	err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, url, method, fmt.Sprintf("%v 在 %v 登陆此系统",
		admin.Username, lastLoginTime))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("账号: %v 登陆时记录日志失败", params.Username), zap.Error(err))
	}

	return token, nil
}
