package common

const (
	LoginGetCaptchaFail        = 10000
	LoginGetCaptchaMobileEmpty = 10001
	LoginGetCaptchaServiceFail = 10002

	LoginLoginFail            = 10010
	LoginLoginRequestBindFail = 10011
	LoginLoginCopyFail        = 10012
	LoginLoginServiceFail     = 10013

	LoginRegisterFail              = 10020
	LoginRegisterRequestBindFail   = 10021
	LoginRegisterRequestVerifyFail = 10022
	LoginRegisterCopyFail          = 10023
	LoginRegisterServiceFail       = 10024

	LoginGetOrgListFail           = 10030
	LoginGetOrgListMemberNotExist = 10031
	LoginGetOrgListServiceFail    = 10032
	LoginGetOrgListCopyFail       = 10033

	ProjectIndexFail         = 20000
	ProjectIndexServiceFail  = 20001
	ProjectIndexCopyFail     = 20002
	ProjectIndexPageBindFail = 20003

	ProjectSelfProjectFail        = 20010
	ProjectSelfProjectServiceFail = 20011
	ProjectSelfProjectCopyFail    = 20012

	MiddlewareTokenVerifyFail        = 30000
	MiddlewareTokenVerifyServiceFail = 30001
)
