package controller

import (
	"TheFirstGin/config"
	"TheFirstGin/entity"
	"TheFirstGin/model"
	"TheFirstGin/request"
	"TheFirstGin/temp"
	"TheFirstGin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// UserHello /** 测试接口方法
func UserHello(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "HelloWorld！", nil))
}

// PhoneOccupy /** 电话号码查重
func PhoneOccupy(c *gin.Context) {
	// 实例化参数类型
	var param request.PhoneOccupyParam
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 验证参数
	// 参数错误
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	// 电话号码格式错误
	if utils.VerPhoneReg(param.Phone) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "电话号码格式错误", nil))
		return
	}
	// 通过电话号码查询用户
	userList := model.PhoneOccupySQL(param.Phone)
	// 判断是否查到
	if len(userList) > 0 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "此电话号码已被注册！", nil))
		return
	}
	// 未查到
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "可以使用！", nil))
}

// EmailOccupy /** 邮箱查重
func EmailOccupy(c *gin.Context) {
	// 实例化参数类型
	var param request.EmailOccupyParam
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 验证参数
	// 参数错误
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	// 电话号码格式错误
	if utils.VerEmailReg(param.Email) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "邮箱格式错误", nil))
		return
	}
	// 通过电话号码查询用户
	userList := model.EmailOccupySQL(param.Email)
	// 判断是否查到
	if len(userList) > 0 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "此邮箱已被注册！", nil))
		return
	}
	// 未查到
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "可以使用！", nil))
}

// Register /** 注册接口
func Register(c *gin.Context) {
	// 实例化参数类型
	var param request.RegisterParam
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 绑定失败
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	// 判断数据
	// 判断参数是否为空
	if len(param.UserName) == 0 || len(param.Password) == 0 || len(param.Phone) == 0 || len(param.Email) == 0 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数含有空值！", nil))
		return
	}
	// 判断数据是否安全
	// 判断电话号码格式
	if utils.VerPhoneReg(param.Phone) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "电话号码格式错误！", nil))
		return
	}
	// 判断邮箱格式
	if utils.VerEmailReg(param.Email) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "邮箱格式错误！", nil))
		return
	}
	// 通过电话号码查询用户
	phoneList := model.PhoneOccupySQL(param.Phone)
	// 判断是否查到
	if len(phoneList) > 0 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "此电话号码已被注册！", nil))
		return
	}

	// 通过邮箱查询用户
	emailList := model.EmailOccupySQL(param.Email)
	// 判断是否查到
	if len(emailList) > 0 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "此邮箱已被注册！", nil))
		return
	}

	// 对密码进行加密
	bytePwd, err := bcrypt.GenerateFromPassword([]byte(param.Password+config.PwdPrivateKey), config.PwdSalt)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "注册失败！", nil))
		panic("加密失败！")
		return
	}
	userInfo := entity.SysUser{Name: param.UserName, Pwd: string(bytePwd), Phone: param.Phone, Email: param.Email}
	err = model.RegisterSQL(userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "注册失败！", nil))
		logrus.Error("注册失败！")
		return
	}
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "注册成功！", nil))
}

// PhoneLogin /** 电话号码登录
func PhoneLogin(c *gin.Context) {
	// 定义错误
	var err error
	// 获取参数类型实例
	var param request.LoginParam
	// 绑定参数
	err = c.ShouldBindJSON(&param)
	// 判断是否绑定成功
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	// 验证电话号码格式
	if utils.VerPhoneReg(param.Phone) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "电话号码格式错误！", nil))
		return
	}
	// 通过电话号码查询用户
	userList := model.PhoneLoginSQL(param.Phone)
	// 通过判断用户列表长度，判断是否有此用户
	if len(userList) < 1 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "此用户暂未注册！", nil))
	}
	if len(userList) > 1 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "账号异常！", nil))
	}
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userList[0].Pwd), []byte(param.Password+config.PwdPrivateKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "账号/密码错误！", nil))
		return
	}
	// 登录成功
	// 生成Token
	token := utils.GenerateToken(&config.TokenParam{
		UserInfo:       userList[0],
		StandardClaims: jwt.StandardClaims{},
	})
	// 将Token存入数据库
	if len(model.QueryTokenById(userList[0].UID)) != 0 {
		err = model.UpdateTokenSQl(entity.SysToken{
			UserId: userList[0].UID,
			Token:  token,
		})
	} else {
		err = model.SaveTokenSQL(entity.SysToken{
			UserId: userList[0].UID,
			Token:  token,
		})
	}
	// 判断Token是否更新/添加成功
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "登录失败！", nil))
		return
	}
	// 成功
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "登录成功！", map[string]interface{}{"token": token}))
}

// ModifyPassword /** 修改密码
func ModifyPassword(c *gin.Context) {
	// 定义参数
	var param request.ModifyPasswordParam
	err := c.ShouldBindJSON(&param)
	// 获取用户信息
	user, _ := c.Get("user")
	// 判断用户信息是否存在
	if user == nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "登陆失效，请重新登录！", nil))
		return
	}
	tokenParam, ok := user.(*config.TokenParam)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "登陆失效，请重新登录！", nil))
		return
	}
	// 判断参数是否错误
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(tokenParam.UserInfo.Pwd), []byte(param.OldPassword+config.PwdPrivateKey))
	// 未通过验证
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "密码错误！", nil))
		return
	}
	// 对新密码进行加密
	bytePwd, err := bcrypt.GenerateFromPassword([]byte(param.NewPassword+config.PwdPrivateKey), config.PwdSalt)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "修改失败！", nil))
		panic("加密失败！")
		return
	}
	// 修改userInfo里的数据
	tokenParam.UserInfo.Pwd = string(bytePwd)
	// 通过验证，修改密码
	err = model.ModifyPasswordSQL(tokenParam.UserInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "修改失败！", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "修改成功！", nil))
}

// SendEmailCode /** 获取邮箱验证码
func SendEmailCode(c *gin.Context) {
	var param request.EmailCodeParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	if utils.VerEmailReg(param.Email) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "邮箱格式错误！", nil))
	}
	code := utils.SendEmail(param.Email)
	// 将邮箱和验证码一一对应存入缓存中
	temp.SetTempCode(param.Email, code)
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "发送成功！", nil))
}

// ForgetPassword /** 忘记密码
func ForgetPassword(c *gin.Context) {
	// 获取参数类型
	var param request.ForgetPasswordParam
	// 绑定参数
	err := c.ShouldBindJSON(&param)
	// 是否绑定成功
	if err != nil {
		// 绑定失败
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "参数错误！", nil))
		return
	}
	// 判断邮箱格式
	if utils.VerEmailReg(param.Email) {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "邮箱格式错误！", nil))
		return
	}
	// 判断验证码
	code := temp.GetTempCode(param.Email)
	// 验证验证码
	if param.Code != code {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "验证码错误！", nil))
		return
	}
	// 删除验证码缓存
	temp.DeleteTempCode(param.Email)
	// 通过邮箱查询用户信息
	userList := model.QueryUserInfoByEmailSQl(param.Email)
	// 判断用户列表长度
	if len(userList) < 1 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "此账号不存在！", nil))
		return
	}
	if len(userList) > 1 {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "账号异常！", nil))
		return
	}

	// 对密码进行加密
	bytePwd, err := bcrypt.GenerateFromPassword([]byte(param.Password+config.PwdPrivateKey), config.PwdSalt)
	// 加密失败
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "注册失败！", nil))
		panic("加密失败！")
		return
	}
	// 修改原密码
	userList[0].Pwd = string(bytePwd)
	// 执行修改密码SQL
	err = model.ModifyPasswordSQL(userList[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResultType(http.StatusBadRequest, "找回失败！", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResultType(http.StatusOK, "找回成功！", nil))
}
