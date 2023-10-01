package routes

import (
	"TheFirstGin/controller"
	"github.com/gin-gonic/gin"
)

// UserRouter /** 用户路由方法
func UserRouter(router *gin.RouterGroup) {
	// 测试
	router.GET("/hello", controller.UserHello)
	// 电话号码查重
	router.POST("/phoneOccupy", controller.PhoneOccupy)
	// 邮箱查重
	router.POST("/emailOccupy", controller.EmailOccupy)
	// 注册
	router.POST("/register", controller.Register)
	// 电话号码登录
	router.POST("/phoneLogin", controller.PhoneLogin)
	// 修改密码
	router.POST("/modifyPwd", controller.ModifyPassword)
	// 发送验证码
	router.POST("/sendEmailCode", controller.SendEmailCode)
	// 忘记密码
	router.POST("/forgetPassword", controller.ForgetPassword)
	// 获取用户信息
	router.GET("/getUserInfo", controller.GetUserInfo)
	// 退出登录
	router.POST("/logout", controller.Logout)
}
