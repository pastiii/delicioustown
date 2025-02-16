package admin

import (
	AdminServices "DeliciousTown/app/service/admin-services"
	"DeliciousTown/app/util/response"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMenu(ctx *gin.Context) {
	var params ValidateRules.CreateMenu
	err := ctx.BindJSON(&params)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.CreateMenu(params)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}

func EditMenu(ctx *gin.Context) {
	var params ValidateRules.EditMenu
	err := ctx.BindJSON(&params)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.EditMenu(params)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}

func DelMenu(ctx *gin.Context) {
	var params ValidateRules.EditMenu
	err := ctx.BindJSON(&params)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.EditMenu(params)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}
