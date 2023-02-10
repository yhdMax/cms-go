package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() (g *gin.Engine) {
	// 初始化全局 store
	store := cookie.NewStore([]byte("secret"))
	g = gin.Default()

	// 使用session中间件
	g.Use(sessions.Sessions("admin_token", store))

	// router
	Routes(g)

	return
}
