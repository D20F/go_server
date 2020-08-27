package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	config "go_server/config"

)



func init() {}


// token 验证中间件
func MiddleToken(ctx *gin.Context) {
	tokenString := ctx.GetHeader("token")
	
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	token, claims, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	fmt.Println(claims.UserId)
	ctx.JSON(200, gin.H{
		"messeage": "成功",
		"UserId":claims.UserId,
		"UserName":claims.UserName,
	})
}


func parseToken(tokenString string) (*jwt.Token, *config.Claims, error) {
	var jwtkey string
	Claims := &config.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}