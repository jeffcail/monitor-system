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
