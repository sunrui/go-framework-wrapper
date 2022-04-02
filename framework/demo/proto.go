package demo

// 发送验证码 - 请求
type postSmsReq struct {
	Phone string `json:"phone" validate:"required,len=11"` // 手机号
}
