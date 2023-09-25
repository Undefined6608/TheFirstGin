package routes

import "github.com/gin-gonic/gin"

// SetupRouterGroup /** 一级路由方法
func SetupRouterGroup(router *gin.RouterGroup) {
	// 用户路由
	UserRouter(router.Group("/user"))
	// 代办路由
	TodosRouter(router.Group("/todo"))
}
