package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	api "go_server/api"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	fmt.Println("初始化")

}

func main() {

	// use注册中间件
	router.Use(api.Middle)

	router.GET("/ping", api.Ping)
	router.GET("/welcome", api.Welcome)
	router.GET("/form", api.Form)
	router.POST("/upload", api.Upload)

	router.GET("/Redirect1", api.Redirect1)
	router.GET("/Redirect2", api.Redirect2)

	// 托管静态目录
	// router.StaticFS("/more_static", http.Dir("./"))
	// 托管单个文件
	// router.StaticFile("/logo.png", ".././express/dist/logo.png")

	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
