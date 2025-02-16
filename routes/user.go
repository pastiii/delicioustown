package routes

import (
	"DeliciousTown/app/controller/user"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	UserRoute := r.Group("user")
	{
		UserRoute.POST("user/create", user.Create)
		UserRoute.GET("user/getUserInfo/:id", user.GetUserInfo)
		UserRoute.GET("test", user.TestZapLog)

	}
}
