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
	ServeIp     string `json:"serve_ip" validate:"required"`
	PackageName string `json:"package_name" validate:"required"`
	PackagePath string `json:"package_path" validate:"required"`
}
