package admin

import (
	AdminServices "DeliciousTown/app/service/admin-services"
	"DeliciousTown/app/util/response"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCuisine(ctx *gin.Context) {
	var addCuisine ValidateRules.AddCuisine
	err := ctx.BindJSON(&addCuisine)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, addCuisine)
	if !res {
		return
	}

	ret, msg := AdminServices.AddCuisine(addCuisine)
	if !ret && msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, true)
}

func GetCuisineInfo(ctx *gin.Context) {
	var params ValidateRules.OnlyId
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.GetCuisineInfo(params.CuisineId)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}

func GetCuisineList(ctx *gin.Context) {
	var params ValidateRules.CuisineList
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.GetCuisineList(params)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}

func EditCuisine(ctx *gin.Context) {
	var params ValidateRules.EditCuisine
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.EditCuisine(params)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}

func DelCuisine(ctx *gin.Context) {
	var params ValidateRules.OnlyId
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := validator.BaseValidate(ctx, params)
	if !res {
		return
	}

	ret, msg := AdminServices.DelCuisine(params.CuisineId)
	if msg != nil {
		response.Error(ctx, http.StatusInternalServerError, msg.Error())
		return
	}

	response.Success(ctx, ret)
}
