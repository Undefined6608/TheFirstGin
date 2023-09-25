package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	allowAccess := []string{"http://127.0.0.1", "http://39.101.72.168", "http://localhost"}
	config.AllowOrigins = allowAccess // 允许访问的域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	return cors.New(config)
}
