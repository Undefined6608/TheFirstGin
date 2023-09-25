package entity

// SysToken /** sys_user 表结构
type SysToken struct {
	ID     uint32 `json:"id" gorm:"primaryKey"`            // 主键 用户ID
	UserId uint32 `json:"user_id" gorm:"not null;unique"`  // 用户ID
	Token  string `json:"token" gorm:"not null;type:text"` // 用户Token
}

// TableName /** 复写默认方法，设置表名
func (SysToken) TableName() string {
	return "sys_token"
}
