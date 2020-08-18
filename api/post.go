package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)


func init() {}

// 表单
func Form(c *gin.Context) {
	typ := c.PostForm("typ")
	msg := c.PostForm("msg")
	title := c.PostForm("title")
	fmt.Println("typ is %s, msg is %s, title is %s", typ, msg, title)
}
