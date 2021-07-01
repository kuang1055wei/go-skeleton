package errors

var (
	//refreshToken相关
	RefreshTokenError = NewError(1001, "refreshToken不合法")
	//token相关
	TokenExistError     = NewError(1004, "token不存在")
	TokenRuntimeError   = NewError(1005, "token已过期")
	TokenWrongError     = NewError(1006, "token不正确")
	TokenTypeWrongError = NewError(1007, "token格式不正确")
	CaptchaError        = NewError(1000, "验证码错误")
)
