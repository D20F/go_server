package main

import (
	"fmt"
	api "go_server/api"
	config "go_server/config"
	middleware "go_server/middleware"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	router.GET("/set", setting, middleware.MiddleToken)
	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

//颁发token
func setting(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &config.Claims{
		UserId:   2,
		UserName: "哇嗷",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(config.Jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, gin.H{"token": tokenString})
}
