package admin

import (
	AdminServices "DeliciousTown/app/service/admin-services"
	"DeliciousTown/app/util/response"
	ValidateVules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var loginAccount ValidateVules.AdminLogin
	err := ctx.BindJSON(&loginAccount)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	res := validator.BaseValidate(ctx, loginAccount)
	if !res {
		return
	}

	ret, user := AdminServices.Login(loginAccount)
	if !ret {
		response.Error(ctx, http.StatusInternalServerError, "登录失败!请核对账户信息")
		return
	}

	response.Success(ctx, user)
}

func Logout(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	ret := AdminServices.Logout(token)
	if !ret {
		response.Error(ctx, http.StatusInternalServerError, "退出登录失败!请联系管理员")
		return
	}

	response.Success(ctx, ret)
}
