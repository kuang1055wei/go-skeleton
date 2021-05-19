package common

import "gin-test/utils"

var (
	//refreshToken相关
	RefreshTokenError = utils.NewError(1001, "refreshToken不合法")
	//token相关
	TokenExistError     = utils.NewError(1004, "token不存在")
	TokenRuntimeError   = utils.NewError(1005, "token已过期")
	TokenWrongError     = utils.NewError(1006, "token不正确")
	TokenTypeWrongError = utils.NewError(1007, "token格式不正确")

	CaptchaError = utils.NewError(1000, "验证码错误")
)
