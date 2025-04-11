package model

import "github.com/Wafer233/msproject-be/common"

const (
	// ===== 通用状态码 =====
	Success        common.BusinessCode = 200
	SystemError    common.BusinessCode = 1001 // 系统内部错误
	InvalidRequest common.BusinessCode = 1002 // 请求参数错误
	Unauthorized   common.BusinessCode = 1003 // 未授权访问
	Forbidden      common.BusinessCode = 1004 // 没权限
	NotFound       common.BusinessCode = 1005 // 数据不存在
	Timeout        common.BusinessCode = 1006 // 超时
	RateLimited    common.BusinessCode = 1007 // 频率限制

	// ===== 登录模块（2xxx）=====
	LoginMobileNotLegal common.BusinessCode = 2001 // 手机号不合法
	LoginCaptchaInvalid common.BusinessCode = 2002 // 验证码无效
	LoginCaptchaExpired common.BusinessCode = 2003 // 验证码已过期
	LoginSendCodeFail   common.BusinessCode = 2004 // 发送验证码失败
	LoginAccountLocked  common.BusinessCode = 2005 // 账号被锁定
	LoginFailed         common.BusinessCode = 2006 // 登录失败，账号密码错误

	// ===== 注册模块（3xxx）=====
	RegisterMobileExists common.BusinessCode = 3001 // 手机号已注册
	RegisterEmailExists  common.BusinessCode = 3002 // 邮箱已注册
	RegisterInvalidData  common.BusinessCode = 3003 // 提交的数据不合法
	RegisterFailed       common.BusinessCode = 3004 // 注册失败（系统原因）

	// ===== 用户模块（4xxx）=====
	UserNotFound     common.BusinessCode = 4001 // 用户不存在
	UserProfileError common.BusinessCode = 4002 // 用户资料异常
	UserUpdateFailed common.BusinessCode = 4003 // 更新用户信息失败

	// ===== Token / Auth（5xxx）=====
	TokenInvalid     common.BusinessCode = 5001 // token无效
	TokenExpired     common.BusinessCode = 5002 // token已过期
	TokenGeneration  common.BusinessCode = 5003 // token生成失败
	PermissionDenied common.BusinessCode = 5004 // 权限不足

	// ===== 外部服务（6xxx）=====
	SMSServiceError     common.BusinessCode = 6001 // 调用短信服务失败
	ThirdPartyAPIFailed common.BusinessCode = 6002 // 第三方接口失败
)
