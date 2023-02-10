/*
 * @Author       : Eug
 * @Date         : 2022-12-18 18:04:16
 * @LastEditTime : 2022-12-18 18:07:14
 * @LastEditors  : Eug
 * @Description : Description
 * @FilePath     : /vue-next-admin-go/main.go
 */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vue-next-admin-go/config"
	"vue-next-admin-go/db"
	"vue-next-admin-go/router"
)

func main() {
	// 设置gin的运行模式 默认为debug模式 可设置为release模式
	gin.SetMode(gin.ReleaseMode)

	// 读取 config.json配置
	config.InitConfiguration()

	// 链接mysql
	db.InitMysql(config.Config)

	/* 使用gin的router开启端口或使用自定义 http 配置 监听端口 */
	server := &http.Server{
		Addr:    ":9090",
		Handler: router.InitRouter(),
	}
	// 设置监听端口
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
