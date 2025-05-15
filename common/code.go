package common

const (

	// ===== LoginService（2xxx）=====
	LoginMobileNotLegal                   BusinessCode = 2001 // 手机号不合法
	LoginGetCaptchaFail                   BusinessCode = 2002 // 生成验证码失败
	LoginRegisterBindFail                 BusinessCode = 2003 // 注册绑定失败
	LoginRegisterReqFormatError           BusinessCode = 2004 // 注册请求参数格式错误
	LoginRegisterServiceError             BusinessCode = 2005 // 注册服务调用失败
	LoginLoginReqBindFail                 BusinessCode = 2006 // 登录请求参数绑定失败
	LoginTokenVerifyMiddlewareServiceFail BusinessCode = 2007 // 登录请求参数绑定失败

	// ===== OrganizationService（3xxx）=====
	OganizationServiceFail BusinessCode = 3001 // 调用GetOrgList服务失败

	// ===== IndexService（4xxx）=====
	IndexIndexServiceFail BusinessCode = 4002 // 调用Index服务失败

	// ===== ProjectService（5xxx）=====
	ProjectSelfListFail BusinessCode = 5001 // 调用Project服务失败

)
