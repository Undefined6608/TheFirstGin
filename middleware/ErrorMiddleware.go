package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ErrorMiddleware /** 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 在这里处理错误，例如记录日志、返回特定的错误响应等
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				log.Fatal(err)
			}
		}()
		c.Next()
	}
}
