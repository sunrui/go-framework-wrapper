package demo

// Sms 验证码
type Sms struct {
	Phone string `json:"phone" validate:"required,len=11"` // 手机号
	Code  string `json:"code" validate:"required,len=6"`   // 验证码
}
