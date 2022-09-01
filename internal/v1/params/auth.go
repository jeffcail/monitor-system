package params

// AdminLoginParams
type AdminLoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=6,lte=12"`
}
