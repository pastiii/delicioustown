package admin

import "DeliciousTown/app/model"

func AdminLoginVo(user model.AdminUser) map[string]interface{} {
	return map[string]interface{}{
		"account": user.Account,
		"token":   user.Token,
		"status":  user.Status,
	}
}

