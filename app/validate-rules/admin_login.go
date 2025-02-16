package validate_rules

type AdminLogin struct {
	Account  string `json:"account" from:"account" validate:"required,max=20,min=3"`
	Password string `json:"password" from:"password" validate:"required,max=20,min=6"`
	Status   string `json:"status" validate:"omitempty"`
}
