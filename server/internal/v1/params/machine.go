package params

// MachineListParams
type MachineListParams struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page_size" validate:"required"`
}

// DeleteMachineParams
type DeleteMachineParams struct {
	ID int64 `json:"id" validate:"required"`
}

// SendComParams
type SendComParams struct {
	Ip      string `json:"ip" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// UpdateMachineRemarkParams
type UpdateMachineRemarkParams struct {
	MachineCode string `json:"machine_code" validate:"required"`
	Ip          string `json:"ip" validate:"required"`
	Remark      string `json:"remark"`
}

// UpgradeClientMachine
type UpgradeClientMachineParams struct {
	MachineCode     string `json:"machine_code" form:"machine_code" validate:"required"`
	MachineIp       string `json:"machine_ip" form:"machine_ip" validate:"required"`
	MachineHostname string `json:"machine_hostname" form:"machine_hostname" validate:"required"`
	MachineRemark   string `json:"machine_remark" form:"machine_remark" validate:"required"`
	UpgradeVersion  string `json:"upgrade_version" form:"upgrade_version" validate:"required"`
}

// UpgradeClientMachineRecord
type UpgradeClientMachineRecord struct {
	Page      int    `json:"page" validate:"required"`
	PageSize  int    `json:"page_size" validate:"required"`
	MachineIp string `json:"machine_ip" validate:"required"`
}
