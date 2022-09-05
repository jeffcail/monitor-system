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
	ServeAddress string `json:"serve_address" validate:"required"`
}
