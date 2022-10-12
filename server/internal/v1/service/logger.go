package service

import (
	"fmt"
	"time"

	_const "github.com/c/server-monitoring/common/const"
	"github.com/c/server-monitoring/common/ubzer"
	"github.com/c/server-monitoring/server/internal/v1/daos"
	"github.com/c/server-monitoring/server/internal/v1/models"
	"github.com/c/server-monitoring/server/internal/v1/params"
)

// log query return params list
type LogQueryReturnParam struct {
	Id            int64  `json:"id"`
	AdminId       int64  `json:"admin_id" `
	AdminUsername string `json:"admin_username" `
	AdminRealName string `json:"admin_real_name" `
	Url           string `json:"url" `
	Method        string `json:"method" `
	Content       string `json:"content" `
	CreatedAt     string `json:"created_at" `
	UpdatedAt     string `json:"updated_at" `
	Version       int64  `json:"version"`
}

func LogQueryList(params *params.LogQueryParam, admin *models.MonAdmin) (int64, []*LogQueryReturnParam, error) {
	t := time.Now().Format(_const.Layout)
	count, data, err := daos.LogQuery(params, filterLogQuery(params))
	if err != nil {
		ubzer.MLog.Error(fmt.Sprintf("%v 在 %v 查看管理员日志记录列表时获取数据失败", admin.Username, t))
		return 0, nil, err
	}
	list := make([]*LogQueryReturnParam, 0)
	for _, bean := range data {
		res := &LogQueryReturnParam{
			Id:            bean.Id,
			AdminId:       bean.AdminId,
			AdminUsername: bean.AdminUsername,
			AdminRealName: bean.AdminRealName,
			Url:           bean.Url,
			Method:        bean.Method,
			Content:       bean.Content,
			CreatedAt:     bean.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     bean.UpdatedAt.Format("2006-01-02 15:04:05"),
			Version:       bean.Version,
		}
		list = append(list, res)
	}

	return count, list, nil
}
func filterLogQuery(params *params.LogQueryParam) map[string]interface{} {

	m := make(map[string]interface{})
	if params.Username != "" {
		m["admin_username"] = params.Username
	}

	return m
}
