package params

// CreateServeParams
type CreateServeParams struct {
	ServeName    string `json:"serve_name" validate:"required"`
	ServeAddress string `json:"serve_address" validate:"required"`
}

// DeleteServeParams
type DeleteServeParams struct {
	ID int64 `json:"id" validate:"required"`
}

// ServeListParams
type ServeListParams struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page_size" validate:"required"`
}

// UpgradeServeParams
type UpgradeServeParams struct {
	UpgradeVersion string `json:"upgrade_version" form:"upgrade_version" validate:"required"`
	ServeId        int64  `json:"serve_id" form:"serve_id" validate:"required"`
	ServeName      string `json:"serve_name" form:"serve_name" validate:"required"`
	ServeAddress   string `json:"serve_address" form:"serve_address" validate:"required"`
	ServeState     int    `json:"serve_state" form:"serve_state" validate:"required"`
}

// UpgradeRecord
type UpgradeRecord struct {
	Page         int    `json:"page" validate:"required"`
	PageSize     int    `json:"page_size" validate:"required"`
	ServeAddress string `json:"serve_address" validate:"required"`
}
