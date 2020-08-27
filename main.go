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
	router.Use(api.Middle)
	router.Use(middleware.Cors())

	router.GET("/set", api.Setting, middleware.MiddleToken)
	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

