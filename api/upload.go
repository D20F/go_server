package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)


func init() {}

// 单文件上传
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)
	// 上传文件到指定的路径
	err := c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	}
}