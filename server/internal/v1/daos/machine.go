package daos

import (
	"bz.service.cloud.monitoring/common/db"
	"bz.service.cloud.monitoring/server/internal/v1/models"
)

// MachineList
func MachineList() ([]*models.MonMachine, error) {
	data := make([]*models.MonMachine, 0)

	err := db.Mysql.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
