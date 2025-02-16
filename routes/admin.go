package routes

import (
	"DeliciousTown/app/controller/admin"
	"DeliciousTown/app/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	AdminRoute := r.Group("admin")
	{
		//admin用户路由
		UserRoute := AdminRoute.Group("user").Use(middleware.AdminAuth())
		{
			UserRoute.POST("/CreateUser", admin.CreateUser)
			UserRoute.GET("/GetUserInfo", admin.GetUserInfo)
			UserRoute.POST("/EditPass", admin.EditPass)
			UserRoute.POST("/EditUserStatus", admin.EditUserStatus)
			UserRoute.GET("/UserList", admin.UserList)
			UserRoute.GET("/DelUser", admin.DelUser)
		}

		//登录与退出登录
		AdminRoute.POST("login/login", admin.Login)
		AdminRoute.Use(middleware.AdminAuth()).GET("login/logout", admin.Logout)

		//菜系相关
		CuisineRoute := AdminRoute.Group("cuisine").Use(middleware.AdminAuth())
		{
			CuisineRoute.POST("AddCuisine", admin.AddCuisine)
			CuisineRoute.GET("GetCuisineInfo", admin.GetCuisineInfo)
			CuisineRoute.GET("GetCuisineList", admin.GetCuisineList)
			CuisineRoute.POST("EditCuisine", admin.EditCuisine)
			CuisineRoute.GET("DelCuisine", admin.DelCuisine)
		}

		//菜谱相关
		MenuRoute := AdminRoute.Group("menu").Use(middleware.AdminAuth())
		{
			MenuRoute.POST("CreateMenu", admin.CreateMenu)
			MenuRoute.POST("EditMenu", admin.EditMenu)
		}

	}
}
