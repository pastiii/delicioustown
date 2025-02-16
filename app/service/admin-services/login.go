package admin_services

import (
	"DeliciousTown/app/model"
	"DeliciousTown/app/repositories"
	"DeliciousTown/app/util/common"
	"DeliciousTown/app/util/redis"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/vo/admin"
	"DeliciousTown/global"
	"context"
	"encoding/json"
	"time"
)

func Login(params ValidateRules.AdminLogin) (bool, map[string]interface{}) {
	// 设置状态为默认开启
	dao := repositories.UserDao{
		Status:   "on",
		Password: common.GetPassword(params.Password),
		Account:  params.Account,
	}
	user, err := dao.GetUserByParams()
	if err != nil {
		return false, nil
	}

	//生成token 缓存用户信息
	user.Token = common.GetToken(user.Account, user.Password)
	AddAuthCache(user)

	return true, admin.AdminLoginVo(user)
}

func AddAuthCache(user model.AdminUser) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	userStr,_ := json.Marshal(user)
	global.GvaRedis.SetEX(ctx, user.Token, userStr, time.Hour*24)
	return true
}

func Logout(token string) bool {
	return redis.ForgetValue(token)
}
