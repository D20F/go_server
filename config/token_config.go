package config

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   uint
	UserName string
	jwt.StandardClaims
}

var Jwtkey = []byte("D20F")

// stus := []student{
// 	{name: "pprof.cn", age: 18},
// 	{name: "测试", age: 23},
// 	{name: "博客", age: 28},
// }
