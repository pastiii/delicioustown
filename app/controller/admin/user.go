package admin

import (
	AdminServices "DeliciousTown/app/service/admin-services"
	"DeliciousTown/app/util/response"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/validator"
	"DeliciousTown/app/vo/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var AdminUser ValidateRules.AdminUser
	err := ctx.BindJSON(&AdminUser)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, AdminUser)
	if !res {
		return
	}

	checkRes := AdminServices.CheckAccount(AdminUser)
	if !checkRes {
		response.Error(ctx, http.StatusInternalServerError, "用户已存在请勿重复注册")
		return
	}

	// 逻辑处理
	ret := AdminServices.CreateAdminUser(AdminUser)
	if !ret {
		response.Error(ctx, http.StatusInternalServerError, "用户添加失败")
		return
	}

	response.Success(ctx, nil)
}

func GetUserInfo(ctx *gin.Context) {
	userId := ctx.Query("id")
	Id, err := strconv.Atoi(userId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "传参类型错误")
		return
	}

	ret, res := AdminServices.GetUserInfo(Id)
	if !res {
		response.Error(ctx, http.StatusInternalServerError, "获取用户信息失败")
		return
	}

	response.Success(ctx, admin.GetUserInfo(ret))
}

func EditPass(ctx *gin.Context) {
	var editPass ValidateRules.EditPass
	err := ctx.BindJSON(&editPass)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, editPass)
	if !res {
		return
	}

	ret, errs := AdminServices.EditPass(editPass)
	if !ret && errs != nil{
		response.Error(ctx, http.StatusInternalServerError, errs.Error())
		return
	}

	response.Success(ctx, ret)
}

func EditUserStatus(ctx *gin.Context)  {
	var userInfo ValidateRules.EditUserStatus
	err :=ctx.BindJSON(&userInfo)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, userInfo)
	if !res {
		return
	}

	ret, errs := AdminServices.EditUserStatus(userInfo)
	if !ret && errs != nil{
		response.Error(ctx, http.StatusInternalServerError, errs.Error())
		return
	}

	response.Success(ctx, ret)
}

func UserList(ctx *gin.Context)  {
	var userList ValidateRules.AdminUserList
	err := ctx.ShouldBindQuery(&userList)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, userList)
	if !res {
		return
	}

	total, list := AdminServices.UserList(userList)
	response.Success(ctx, admin.GetList(total, list))
}

func DelUser(ctx *gin.Context)  {
	userId := ctx.Query("id")
	Id, err := strconv.Atoi(userId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "传参类型错误")
		return
	}

	ret, res := AdminServices.DelUser(Id)
	if res != nil {
		response.Error(ctx, http.StatusInternalServerError, res.Error())
		return
	}

	response.Success(ctx, ret)
}
