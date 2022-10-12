package handler

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/c/server-monitoring/server/config"

	"github.com/c/server-monitoring/common/ubzer"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"github.com/c/server-monitoring/server/internal/v1/service"

	_const "github.com/c/server-monitoring/common/const"
	"github.com/c/server-monitoring/common/utils"
	"github.com/c/server-monitoring/server/internal/v1/params"
	"github.com/labstack/echo"
)

// ServeList
func ServeList(c echo.Context) error {
	params := &params.ServeListParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	count, data, err := service.ServeList(params, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "查看列表失败", ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功",
		utils.Resp.ResponsePagination(count, data)))
}

// CreateServe
func CreateServe(c echo.Context) error {
	serveParams := &params.CreateServeParams{}
	_ = c.Bind(serveParams)
	msg := utils.ValidateParam(serveParams)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.CreateServe(serveParams, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "创建成功", ""))
}

// DeleteServe
func DeleteServe(c echo.Context) error {
	params := &params.DeleteServeParams{}
	_ = c.Bind(params)
	msg := utils.ValidateParam(params)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}
	err := service.DeleteServe(params, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "删除成功", ""))
}

// UpgradeServe
func UpgradeServe(c echo.Context) error {
	param := &params.UpgradeServeParams{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, msg, ""))
	}

	if strings.Contains(param.UpgradeVersion, ".") {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "版本号只能为整数版本", ""))
	}

	err := service.CheckUpgradeVersion(param.ServeAddress, param.UpgradeVersion)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	file, err := c.FormFile("go_file")
	if err != nil {
		ubzer.MLog.Error("文件接收失败", zap.Error(err))
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "文件接收失败", ""))
	}
	src, err := fileCheck(file)
	defer src.Close()
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	packageName := file.Filename + ":" + param.UpgradeVersion
	err = checkFileIsExists(packageName)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	dst, err := os.Create(config.Config().UpPkgPath + packageName)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	err = service.UpgradeServe(param, file.Filename, GetAdminInfoFromParseToken(c), c.Request().URL.Path, c.Request().Method)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}

	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "上传成功，60秒内自动升级此服务", ""))

	////lastIndex := strings.LastIndex(param.ServeAddress, ":")
	////s := param.ServeAddress[:lastIndex]
	////index := strings.LastIndex(s, ":")
	////s2 := s[index+3:]
	//
	//header := make(map[string]string)
	//p := make(map[string]interface{})
	//p["package_name"] = param.PackageName
	//p["package_path"] = param.PackagePath
	//res, err := request.GetParams("http://"+param.ServeIp+":8888/c/serve/upgrade", header, p)
	//ubzer.MLog.Info(fmt.Sprintf("升级=========== res: %v", string(res)))
	//if err != nil {
	//	ubzer.MLog.Error(fmt.Sprintf("升级服务请求失败 ServeAddress: %v", param.ServeIp), zap.Error(err))
	//	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, "升级服务请求失败", ""))
	//}
	//type FRes struct {
	//	str string
	//}
	//
	////if f.str != "success" {
	////	ubzer.MLog.Error(fmt.Sprintf("升级服务失败 ServeAddress: %v", param.ServeIp), zap.Error(err))
	////	return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, string(res), ""))
	////}
	//admin := GetAdminInfoFromParseToken(c)
	//err = daos.RecordOperateLog(admin.Id, admin.Username, admin.RealName, c.Request().URL.Path, c.Request().Method,
	//	fmt.Sprintf("%v 在 %v 时间升级了服务地址为 %v 的服务", admin.Username, time.Now(), param.ServeIp))
	//if err != nil {
	//	ubzer.MLog.Error("记录操作日志失败", zap.Error(err))
	//}
	//return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, string(res), ""))
}

// fileCheck
func fileCheck(file *multipart.FileHeader) (multipart.File, error) {
	// 字节 100 * 1024 * 1024
	if file.Size > 1024*1024*100 {
		return nil, errors.New("文件过大,不能超过100M")
	}

	if strings.Contains(file.Filename, ".") {
		return nil, errors.New("非法文件,文件不允许带后缀")
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	return src, nil
}

// checkFileIsExists
func checkFileIsExists(packageName string) error {
	_, err := os.Stat(config.Config().UpPkgPath)
	if err != nil {
		err = os.Mkdir(config.Config().UpPkgPath, 0777)
		return err
	}
	_, err = os.Stat(config.Config().UpPkgPath + packageName)
	if !os.IsNotExist(err) {
		return errors.New("已经存在此版本发布的包，请检查输入的版本号，或重新输入版本号")
	}
	return nil
}

// UpgradeRecord
func UpgradeRecord(c echo.Context) error {
	param := &params.UpgradeRecord{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(msg), ""))
	}

	count, records, err := service.UpgradeRecord(param)
	if err != nil {
		return c.JSON(http.StatusOK, utils.Res.ResponseJson(false, _const.Fail, cast.ToString(err), ""))
	}
	return c.JSON(http.StatusOK, utils.Res.ResponseJson(true, _const.Success, "成功", utils.Resp.ResponsePagination(count, records)))
}
