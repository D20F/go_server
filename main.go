package main

import (
	"encoding/json"
	api "go_server/api"
	middleware "go_server/middleware"
	"strings"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}
func JsonDump(v interface{}, indent int) string {
	var b []byte
	var err error
	if indent > 0 {
		b, err = json.MarshalIndent(v, " ", strings.Repeat(" ", indent))
	} else {
		b, err = json.Marshal(v)
	}

	if err != nil {
		return ""
	}
	return string(b)
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
