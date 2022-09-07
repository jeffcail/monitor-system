package daos

import (
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
	"bz.service.cloud.monitoring/server/internal/v1/params"
)

// log query
func LogQuery(params *params.LogQueryParam, filterLogQuery map[string]interface{}) (int64, []*models.MonOperateRecord, error) {
	list := make([]*models.MonOperateRecord, 0)
	query := db.Mysql.NewSession().Limit(params.PageSize, (params.Page-1)*params.PageSize)
	query2 := db.Mysql.NewSession()
	for k, v := range filterLogQuery {
		query.Where(k+" = ? ", v)
		query2.Where(k+" = ? ", v)
	}
	if params.StateTime != "" {
		query.Where("updated_at >= ?", params.StateTime)
		query2.Where("updated_at >= ?", params.StateTime)
	}
	if params.StopTime != "" {
		query.Where("updated_at <=?", params.StopTime)
		query2.Where("updated_at <= ?", params.StopTime)
	}
	count, err := query.FindAndCount(&list)
	query.Desc("updated_at")
	return count, list, err
}
