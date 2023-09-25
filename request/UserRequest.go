package request

// PhoneOccupyParam /** 用户查重参数
type PhoneOccupyParam struct {
	Phone string `json:"phone" binding:"required,len=11"`
}

// EmailOccupyParam /** 邮箱查重参数
type EmailOccupyParam struct {
	Email string `json:"email" binding:"required"`
}

// RegisterParam /** 注册参数
type RegisterParam struct {
	UserName string `json:"username" binding:"required,min=1,max=15"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,len=32"`
	Email    string `json:"email" binding:"required"`
}

// LoginParam /** 电话号码登录参数
type LoginParam struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,len=32"`
}

// ModifyPasswordParam /** 修改密码参数
type ModifyPasswordParam struct {
	OldPassword string `json:"oldPassword" binding:"required,len=32"`
	NewPassword string `json:"newPassword" binding:"required,len=32"`
}

// EmailCodeParam /** 邮箱参数
type EmailCodeParam struct {
	Email string `json:"email" binging:"required"`
}

// ForgetPasswordParam /** 忘记密码参数
type ForgetPasswordParam struct {
	Email    string `json:"email" binging:"required"`
	Code     string `json:"code" binging:"required"`
	Password string `json:"password" binging:"required,len=32"`
}
