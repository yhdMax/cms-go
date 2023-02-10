package router

import (
	"github.com/gin-gonic/gin"
	"vue-next-admin-go/service"
)

func Routes(g *gin.Engine) {
	// 路由分组
	router := g.Group("/api")
	{
		router.POST("/login", service.Login)
		router.GET("/captcha", service.SetCaptcha)
		router.POST("/test", service.SignIn)
	}
}
