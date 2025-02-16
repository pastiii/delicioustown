package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	SuccessCode = 200
	SuccessMsg  = "success"
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

// 请求响应
func ResultJson(ctx *gin.Context, code int, msg interface{}, data interface{}) {
	// 格式化时间
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	})
}

func Success(ctx *gin.Context, data interface{}) {
	ResultJson(ctx, SuccessCode, SuccessMsg, data)
}

func Error(ctx *gin.Context, code int, msg string) {
	ResultJson(ctx, code, msg, nil)
}

func ValidateError(ctx *gin.Context, msg interface{}) {
	ResultJson(ctx, http.StatusBadRequest, msg, nil)
}

func InternalServerError(ctx *gin.Context, msg string, data interface{}) {
	ResultJson(ctx, http.StatusInternalServerError, msg, data)
}
