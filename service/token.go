package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

var Claims = make(map[string]interface{})

// 使用当前时间 + string 组合一个签名 secret
var nowData = time.Now().Format("2006-01-02 15")
var Secret = fmt.Sprintf("%v%v", nowData, "next_admin")

// GenerateToken 生成token	调用库的 NewWithClaims (加密方式,载荷).SignedString(签名) 生成token
func GenerateToken(mapClaims jwt.MapClaims) (token string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims).SignedString([]byte(Secret))
	return
}

// Verification 校验token
func Verification(g *gin.Context, token string) bool {
	session := sessions.Default(g)
	seToken := session.Get("admin_token")
	if token != seToken {
		return false
	}
	return true
}
