package model

import (
	"TheFirstGin/config"
	"TheFirstGin/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Pool /** 创建数据库连接池
func Pool() *gorm.DB {
	// 获取数据库配置
	dbConfig := config.Default().DataBase
	fmt.Println(dbConfig.Host)
	// 创建数据库连接
	db, err := gorm.Open(mysql.Open(dbConfig.UserName+":"+dbConfig.Password+"@tcp("+dbConfig.Host+")/"+dbConfig.Schema+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// 连接失败
	if err != nil {
		panic("连接数据库失败")
	}
	// 抛出数据库连接池
	return db
}

// VerDataBase /** 验证数据库/表是否已经创建
func VerDataBase() {
	// 获取数据库连接池
	pool := Pool()
	// 验证sys_user表是否存在
	err := pool.Table("sys_user").AutoMigrate(&entity.SysUser{})
	err = pool.Table("sys_token").AutoMigrate(&entity.SysToken{})
	// 表创建失败
	if err != nil {
		panic(err)
	}
}
