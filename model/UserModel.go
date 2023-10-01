package model

import (
	"TheFirstGin/entity"
)

// 获取数据库连接池
var pool = Pool()

// PhoneOccupySQL /** 按电话号码查找用户
func PhoneOccupySQL(val string) []entity.SysUser {
	var user []entity.SysUser
	pool.Where("phone=?", val).Find(&user)
	return user
}

// EmailOccupySQL /** 按电话号码查找用户
func EmailOccupySQL(val string) []entity.SysUser {
	var user []entity.SysUser
	pool.Where("email=?", val).Find(&user)
	return user
}

// RegisterSQL /** 注册
func RegisterSQL(userInfo entity.SysUser) error {
	err := pool.Create(&userInfo).Error
	return err
}

// PhoneLoginSQL /** 电话号码登录
func PhoneLoginSQL(val string) []entity.SysUser {
	var user []entity.SysUser
	pool.Where("phone=?", val).Find(&user)
	return user
}

// QueryTokenById /** 通过用户ID查询Token
func QueryTokenById(userID uint32) []entity.SysToken {
	var token []entity.SysToken
	pool.Where("user_id=?", userID).Find(&token)
	return token
}

// QueryIdByToken /** 通过用户ID查询Token
func QueryIdByToken(tokenStr string) []entity.SysToken {
	var token []entity.SysToken
	pool.Where("token=?", tokenStr).Find(&token)
	return token
}

// SaveTokenSQL /** 创建Token项
func SaveTokenSQL(token entity.SysToken) error {
	err := pool.Create(&token).Error
	return err
}

// UpdateTokenSQl /** 更新Token项
func UpdateTokenSQl(token entity.SysToken) error {
	err := pool.Model(&entity.SysToken{}).Where("user_id", token.ID).Update("token", token.Token).Error
	return err
}

// DeleteTokenById /** 删除Token
func DeleteTokenById(userID uint32) error {
	err := pool.Exec("DELETE FROM sys_token WHERE user_id = ?", userID).Error
	return err
}

// ModifyPasswordSQL /** 修改密码
func ModifyPasswordSQL(userInfo entity.SysUser) error {
	err := pool.Model(&entity.SysUser{}).Where("uid", userInfo.UID).Update("pwd", userInfo.Pwd).Error
	return err
}

// QueryUserInfoByEmailSQl /** 通过邮箱查询用户信息
func QueryUserInfoByEmailSQl(email string) []entity.SysUser {
	var userList []entity.SysUser
	pool.Where("email=?", email).Find(&userList)
	return userList
}
