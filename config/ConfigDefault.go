package config

import (
	"TheFirstGin/entity"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

// Config /** yaml 配置文件格式
type Config struct {
	// 端口号
	Port string `yaml:"port"`
	// 数据库配置
	DataBase struct {
		Host     string `yaml:"host"`     // 数据库地址
		Schema   string `yaml:"schema"`   // 数据库名
		UserName string `yaml:"userName"` // 用户名
		Password string `yaml:"password"` // 密码
	} `yaml:"dataBase"`
	// 日志配置
	LogConfig struct {
		Path string `yaml:"path"` // 日志存放路径
		Name string `yaml:"name"` // 日志名称
	} `yaml:"logConfig"`
	// 私钥
	PrivateKey struct {
		Token   string `yaml:"token"`   // Token私钥
		Pwd     string `yaml:"pwd"`     // 密码私钥
		PwdSalt int    `yaml:"pwdSalt"` // 密码盐值
	} `yaml:"privateKey"`
	// 正则
	Regular struct {
		Phone string `yaml:"phone"` // 电话号码正则
		Email string `yaml:"email"` // 邮箱正则
	} `yaml:"regular"`
	// token配置
	TokenConfig struct {
		NoVerify []string `yaml:"noVerify"` // 过滤验证路由
	} `yaml:"tokenConfig"`
	// 邮箱配置
	EmailConfig struct {
		EmailAddress string `yaml:"emailAddress"` // 邮箱地址
		EmailName    string `yaml:"emailName"`    // 邮箱名
		Password     string `yaml:"password"`     // 密码
		SmtpServer   string `yaml:"smtpServer"`   // 邮箱发送服务
		SmtpPort     int    `yaml:"smtpPort"`     // 邮箱发送接口
	} `yaml:"emailConfig"`
}

// Default /** 获取 yaml 配置
func Default() Config {
	// 实例化配置对象
	var configObj Config
	// 读取配置文件
	yamlFil, err := os.ReadFile("Application.yaml")
	// 读取失败
	if err != nil {
		panic(err)
	}
	// 将读到的文件解析为配置对象
	err = yaml.Unmarshal(yamlFil, &configObj)
	// 解析失败
	if err != nil {
		panic(err)
	}
	// 抛出配置文件对象
	return configObj
}

// Response /**返回值类型
type Response struct {
	Code    int         `json:"code"` // 响应值
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 数据
}

// TokenParam /** token存储数据类型
type TokenParam struct {
	UserInfo           entity.SysUser // 用户信息
	jwt.StandardClaims                // token配置
}

// TokenEffectAge /**Token过期时间
const TokenEffectAge = 7 * 24 * time.Hour

// PhoneReg /** 定义电话号码正则
var PhoneReg = Default().Regular.Phone

// PwdPrivateKey /**密码私钥
var PwdPrivateKey = Default().PrivateKey.Pwd

// PwdSalt /** 密码盐值
var PwdSalt = Default().PrivateKey.PwdSalt

// TokenPrivateKey /**定义 Token 私钥
var TokenPrivateKey = Default().PrivateKey.Token

// TokenNoVerify /**定义不经过验证的接口地址
var TokenNoVerify = Default().TokenConfig.NoVerify

// EmailReg /** 定义邮箱正则
var EmailReg = Default().Regular.Email

// LogFilePath /**日志存放地址
var LogFilePath = Default().LogConfig.Path

// LogFileName /**日志文件名
var LogFileName = Default().LogConfig.Name

var EmailConfig = Default().EmailConfig
