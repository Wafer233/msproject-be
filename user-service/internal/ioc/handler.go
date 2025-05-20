package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	domainSvc "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/grpc/handler"
)

func ProvideLoginGRPCHandler(
	captchaSvc service.CaptchaService,
	loginSvc service.LoginService,
	registerSvc service.RegisterService,
	tokenSvc domainSvc.TokenService,
	organizationSvc service.OrganizationService,
) *handler.LoginGRPCHandler {
	return handler.NewLoginGRPCHandler(captchaSvc, loginSvc, registerSvc, tokenSvc, organizationSvc)
}
