package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		// todo 校验token有效性
		oldToken := "dddd"
		if token != oldToken {
			context.JSON(500, gin.H{"msg": "请先登录111!", "token":token})
			context.Abort()
		}

		context.Next()
	}
}
