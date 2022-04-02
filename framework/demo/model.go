package demo

// Sms 验证码
type Sms struct {
	Phone string `json:"phone" binding:"required" validate:"len=11"` // 手机号
	Code  string `json:"code" binding:"required" validate:"len=6"`   // 验证码
}
