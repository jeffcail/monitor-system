package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	_const "github.com/c/server-monitoring/common/const"

	"go.uber.org/zap"

	"github.com/c/server-monitoring/server/internal/v1/daos"

	"github.com/c/server-monitoring/common/jwt"

	"github.com/c/server-monitoring/common/ubzer"

	"github.com/c/server-monitoring/common/utils"
	"github.com/labstack/echo"
)

// AuthCheck
func AuthCheck() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				ubzer.MLog.Error(fmt.Sprintf("未授权 authHeader: %v", authHeader))
				return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "未授权", ""))
			}
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				ubzer.MLog.Error(fmt.Sprintf("非法Token %v", authHeader))
				return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "非法Token", ""))
			}
			cl, err := jwt.ParseToken(parts[1])
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("Token认证失败 %v", authHeader), zap.Error(err))
				return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "Token认证失败", ""))
			}
			admin, err := daos.GetAdminInfoByUsername(cl.Username)
			if err != nil {
				ubzer.MLog.Error(fmt.Sprintf("校验账号 %v 是否被禁用，获取管理员信息失败", cl.Username), zap.Error(err))
				return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "授权失败", ""))
			}
			if admin.State == _const.AdminStateOff {
				ubzer.MLog.Error(fmt.Sprintf("校验账号 %v 已被禁用", cl.Username))
				return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "账号已经禁用", ""))
			}
			return handlerFunc(c)
		}
	}
}
