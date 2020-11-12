package main

import (
	api "go_server/api"
	middleware "go_server/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func main() {
	// use注册中间件
	// router.Use(api.Middle)
	router.Use(middleware.Cors())

	router.POST("/Upload", api.Upload)
	router.GET("/GetToken", api.GetToken)

	// router.GET("/Ping", middleware.MiddleToken, api.Ping)
	// router.GET("/Pings", api.Ping)

	router.Run(":8085") // 监听并在 0.0.0.0:8080 上启动服务
}
