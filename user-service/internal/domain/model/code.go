package model

import "github.com/Wafer233/msproject-be/common"

const (
	// 通用错误码
	Success     common.BusinessCode = 200
	SystemError common.BusinessCode = 1001 // 系统错误

	// 登录相关错误码
	LoginMobileNotLegal common.BusinessCode = 2001 // 手机号不合法
	LoginCaptchaInvalid common.BusinessCode = 2002 // 验证码无效
	// 其他错误码...
)
