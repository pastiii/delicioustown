package user

import (
	"DeliciousTown/app/util/response"
	ValidateRules "DeliciousTown/app/validate-rules"
	"DeliciousTown/app/validator"
	"DeliciousTown/global"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Create(context *gin.Context) {
	var user ValidateRules.User
	err := context.BindJSON(&user)
	if err != nil {
		response.Error(context, http.StatusInternalServerError, err.Error())
		return
	}

	validateRes := validator.BaseValidate(context, user)
	if !validateRes {
		return
	}
	str, _ := json.Marshal(global.GvaConfig)
	global.UserLogger.Sugar().Info("user", zap.Any("config", global.GvaConfig))
	global.DefaultLogger.Sugar().Info(string(str))
	response.Success(context, global.GvaConfig)
}

func GetUserInfo(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "xiaoMing!!!!1"})
}

func TestZapLog(context *gin.Context) {
	global.UserLogger.Info("xiaoming")
	context.JSON(200, gin.H{"msg": true})
}
