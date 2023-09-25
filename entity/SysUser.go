package entity

// SysUser SysUser/** sys_user 表结构
type SysUser struct {
	UID           uint32 `json:"uid" gorm:"primaryKey"`                                                                      // 主键 用户ID
	Name          string `json:"name"`                                                                                       // 用户名
	Phone         string `json:"phone" gorm:"not null;unique"`                                                               // 电话号码 不允许为空 唯一
	HeadSculpture string `json:"head_sculpture" gorm:"default:'http://39.101.72.168:81/image/icon.jpg'"`                     // 头像 默认值
	Pwd           string `json:"pwd" gorm:"not null;default:'$2a$15$ByxeOHfBdTYkyPZq5Ytr7e2TX7MpnWT3cOA.kwNMK13tUM9AlKym.'"` // 密码 默认值
	Email         string `json:"email" gorm:"not null;unique"`                                                               // 邮箱 不允许为空 唯一
	Available     bool   `json:"available" gorm:"not null;default:0"`                                                        // 是否注销 不允许为空 默认值
	Power         uint8  `json:"power" gorm:"not null;default:2"`                                                            // 权限 不允许为空 默认值
}

// TableName /** 复写默认方法，设置表名
func (SysUser) TableName() string {
	return "sys_user"
}
