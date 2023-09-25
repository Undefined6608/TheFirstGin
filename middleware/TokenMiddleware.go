package middleware

import (
	"TheFirstGin/config"
	"TheFirstGin/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JwtVerifyMiddle /**验证token
func JwtVerifyMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//过滤是否验证token
		if utils.IsContainArr(config.TokenNoVerify, c.Request.RequestURI) {
			return
		}
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.ResultType(http.StatusUnauthorized, "权限验证失败,无法访问系统资源！", nil))
			// 终止请求
			c.Abort()
			return
		}
		// 判断是否错误
		claims := parseToken(c, token)
		fmt.Println(claims)
		//验证token，并存储在请求中
		c.Set("user", claims)
	}
}

// 解析Token
func parseToken(c *gin.Context, tokenString string) *config.TokenParam {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &config.TokenParam{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenPrivateKey), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ResultType(http.StatusUnauthorized, "权限验证失败,无法访问系统资源！", nil))
		c.Abort()
		return nil
	}
	claims, ok := token.Claims.(*config.TokenParam)
	if !ok {
		c.JSON(http.StatusUnauthorized, utils.ResultType(http.StatusUnauthorized, "登录失效！", nil))
		c.Abort()
		return nil
	}
	return claims
}
