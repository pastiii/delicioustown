package validate_rules

type AdminUser struct {
	Account  string `form:"account" json:"account" validate:"required,max=20,min=3"`
	Password string `form:"password" json:"password" validate:"required,max=20,min=6"`
}

type EditPass struct {
	UserId      int    `form:"id" json:"id" validate:"required"`
	OldPassword string `form:"old_password" json:"old_password" validate:"required,max=20,min=6"`
	NewPassword string `form:"new_password" json:"new_password" validate:"required,max=20,min=6"`
}

type EditUserStatus struct {
	UserId int    `form:"id" json:"id" validate:"required"`
	Status string `form:"status" json:"status" validate:"required,adminUserStatus"`
}

type AdminUserList struct {
	Status   string `form:"status" json:"status" validate:"omitempty,adminUserStatus"`
	Account  string `form:"account" json:"account" validate:"omitempty"`
	Page     int    `form:"page" json:"page" validate:"omitempty"`
	PageSize int    `form:"page_size" json:"page_size" validate:"omitempty"`
}
