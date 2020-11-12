package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {}

//全局中间件
func Middle(c *gin.Context) {
	fmt.Println("暂无需求")
}

// 实验数据接口
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "成功",
	})
}
// 实验数据接口2
func Pings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "成功2",
	})
}

// get请求
func Welcome(c *gin.Context) {
	name := c.Query("name")
	lastname := c.Query("lastname")
	fmt.Println("Hello %s", name)
	fmt.Println("Hello %s", lastname)
}

// 内部路由重定向
func Redirect1(c *gin.Context) {
	c.Request.URL.Path = "/Redirect2"
	// router.HandleContext(c)
}
// 外链重定向
func Redirect2(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}