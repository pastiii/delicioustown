package model

type AdminUser struct {
	BaseFiled
	Account  string `json:"account,omitempty" gorm:"varchar(20),not null"`
	Password string `json:"password,omitempty" gorm:"varchar(100),not null"`
	Status   string `json:"status,omitempty" gorm:"default:on"`
	Token    string `json:"token,omitempty" gorm:"varchar(50),null"`
}
