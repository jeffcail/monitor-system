package handler

import (
	"fmt"
	"strings"

	"github.com/c/monitor-system/server/internal/v1/models"

	"github.com/c/monitor-system/server/internal/v1/daos"

	"github.com/c/monitor-system/common/ubzer"

	"github.com/c/monitor-system/common/jwt"

	"github.com/labstack/echo"
)

// GetAdminInfoFromParseToken
func GetAdminInfoFromParseToken(c echo.Context) *models.MonAdmin {
	authHeader := c.Request().Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	cl, err := jwt.ParseToken(parts[1])
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("解析token失败"))
		return nil
	}
	admin, err := daos.GetAdminInfoByUsername(cl.Username)
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("解析完token获取用户信息失败"))
		return nil
	}
	return admin
}
