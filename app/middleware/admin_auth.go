package middleware

import (
	"DeliciousTown/app/util/redis"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		// 获取用户信息
		res := redis.GetValue(token)
		if res == nil {
			context.JSON(500, gin.H{"msg": "请先登录111!", "token": token})
			context.Abort()
		}

		context.Next()
	}
}
