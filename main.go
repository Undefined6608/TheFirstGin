/**
 * @projectName:    TheFirstGin
 * @package:        TheFirstGin
 * @className:      main
 * @author:     张杰
 * @description:  TODO
 * @date:    2023/9/18 8:53
 * @version:    1.0
 */
package main

import (
	"TheFirstGin/config"
	"TheFirstGin/middleware"
	"TheFirstGin/model"
	routes "TheFirstGin/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

// main /** 项目主入口
func main() {
	// 获取路由
	router := gin.Default()
	// 挂载中间件
	router.Use(gin.Logger(), gin.Recovery(), middleware.LoggerToFile(), middleware.CorsMiddleware(), middleware.JwtVerifyMiddle(), middleware.ErrorMiddleware())
	// 加载代理中间件
	err := router.SetTrustedProxies([]string{"192.168.1.0/24"})
	if err != nil {
		fmt.Println("代理失败！")
		return
	}
	// 验证数据库/表是否存在
	model.VerDataBase()
	// 挂载主路由
	routes.SetupRouterGroup(router.Group("/api"))
	// 获取端口配置
	port := config.Default().Port
	// 开启端口监听
	err = router.Run(":" + port)
	// 启动失败
	if err != nil {
		panic(err)
		return
	}
}
