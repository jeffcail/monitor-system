package params

// RunSshServerParams
type RunSshServerParams struct {
	Ip       string `json:"ip" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
