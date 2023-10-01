package request

// SysUserResponse /** sys_user 表结构
type SysUserResponse struct {
	UID           uint32 `json:"uid"`            // 主键 用户ID
	Name          string `json:"name"`           // 用户名
	Phone         string `json:"phone"`          // 电话号码 不允许为空 唯一
	HeadSculpture string `json:"head_sculpture"` // 头像 默认值
	Email         string `json:"email"`          // 邮箱 不允许为空 唯一
	Power         uint8  `json:"power"`          // 权限 不允许为空 默认值
}
