package routes

import (
	"TheFirstGin/controller"
	"github.com/gin-gonic/gin"
)

// TodosRouter /** 代办路由方法
func TodosRouter(router *gin.RouterGroup) {
	router.GET("/addToDos", controller.AddToDo)
}
