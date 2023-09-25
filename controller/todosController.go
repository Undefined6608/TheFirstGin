package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddToDo /** 添加代办事项
func AddToDo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "添加成功！"})
}
