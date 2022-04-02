package demo

// 发送验证码 - 请求
type postSmsReq struct {
	Phone string `json:"phone" binding:"required" validate:"len=11"` // 手机号
}
