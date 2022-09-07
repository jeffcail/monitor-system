package params

// log query params
type LogQueryParam struct {
	Username string `json:"username" ` //用户昵称
	Page     int    `json:"page"  validate:"required"`
	PageSize int    `json:"page_size"  validate:"required"`
}
