package admin

import "DeliciousTown/app/model"

func GetUserInfo(user *model.AdminUser) map[string]interface{} {
	return map[string]interface{}{
		"id":      user.ID,
		"account": user.Account,
		"status":  user.Status,
		"token":   user.Token,
	}
}
