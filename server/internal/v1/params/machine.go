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
