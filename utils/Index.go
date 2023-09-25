// Package utils /** 工具类
package utils

import (
	"TheFirstGin/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gomail.v2"
	"math/rand"
	"regexp"
	"time"
)

// ResultType /** 返回值方法
func ResultType(code int, msg string, data interface{}) config.Response {
	return config.Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

// VerPhoneReg /** 验证电话号码格式
func VerPhoneReg(phone string) bool {
	phoneReg := regexp.MustCompile(config.PhoneReg)
	return !phoneReg.MatchString(phone)
}

// VerEmailReg /** 验证电话号码格式
func VerEmailReg(email string) bool {
	emailReg := regexp.MustCompile(config.EmailReg)
	return !emailReg.MatchString(email)
}

// IsContainArr /**Token URL过滤
func IsContainArr(noVerify []string, requestUrl string) bool {
	for _, str := range noVerify {
		if str == requestUrl {
			return true
		}
	}
	return false
}

// GenerateToken /**生成token
func GenerateToken(claims *config.TokenParam) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(config.TokenEffectAge).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.TokenPrivateKey))

	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		panic(err)
	}
	return sign
}

// Refresh /**更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &config.TokenParam{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenPrivateKey), nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*config.TokenParam)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(config.TokenEffectAge).Unix()
	return GenerateToken(claims)
}

// generateVerificationCode /** 生成验证码
func generateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", r.Intn(10))
	}
	return code
}

func SendEmail(email string) string {
	// 生成验证码
	code := generateVerificationCode()
	// 创建消息
	m := gomail.NewMessage()
	// 设置发件地址和发件人
	m.SetAddressHeader("From", config.EmailConfig.EmailAddress, config.EmailConfig.EmailName)
	// 发送地址
	m.SetHeader("To", email)
	// 设置标题
	m.SetHeader("Subject", "验证码")
	// 设置内容
	m.SetBody("text/html", `
            <p>您好！</p>
            <p>您的验证码是：<strong style="color:orangered;">`+code+`</strong></p>
            <p>如果不是您本人操作，请无视此邮件</p>
        `)
	// 使用 smtp发送邮件
	s := gomail.NewDialer(config.EmailConfig.SmtpServer, config.EmailConfig.SmtpPort, config.EmailConfig.EmailAddress, config.EmailConfig.Password)

	if err := s.DialAndSend(m); err != nil {
		panic("发送失败！")
	}
	return code
}
