package out_field

import (
	"DeliciousTown/app/model"
)

type AdminUserListField struct {
	Id        int64       `json:"id"`
	Account   string      `json:"account"`
	Status    string      `json:"status"`
	Token     string      `json:"token"`
	CreatedAt model.XTime `json:"created_at"`
}
