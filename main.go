package main

import (
	"DeliciousTown/global"
	"DeliciousTown/initialize"
	"DeliciousTown/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.InitViperConfig()
	initialize.InitLogger()
	initialize.InitRedis()
	initialize.InitGorm()
}

func main() {
	r := gin.Default()
	routes.LoadRouter(r)
	routes.AdminRouter(r)
	_ = r.Run(global.GvaConfig.App.Addr)
}
