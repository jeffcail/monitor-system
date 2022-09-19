package daos

import "bz.service.cloud.monitoring/common/db"

// SaveCpuPercent
func SaveCpuPercent(key, value string) (err error) {
	err = db.Rc.Set(key, value, -1).Err()
	return
}

// SaveMemPercent
func SaveMemPercent(key, value string) (err error) {
	err = db.Rc.Set(key, value, -1).Err()
	return
}

// SaveDiskPercent
func SaveDiskPercent(key, value string) (err error) {
	err = db.Rc.Set(key, value, -1).Err()
	return
}
