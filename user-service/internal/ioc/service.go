package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainSvc "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
)

func ProvideDefaultCaptchaService(captchaRepo repo.CaptchaRepo) service.CaptchaService {
	return service.NewDefaultCaptchaService(captchaRepo)
}

func ProvideDefaultTokenService(cfg *config.Config, tokenRepo repo.TokenRepo) domainSvc.TokenService {

	return domainSvc.NewJWTTokenService(cfg, tokenRepo)

}

func ProvideDefaultLoginService(memberRepo repo.MemberRepo, organizationRepo repo.OrganizationRepo) service.LoginService {
	return service.NewDefaultLoginService(memberRepo, organizationRepo)
}

func ProvideDefaultOrganizationService(organizationRepo repo.OrganizationRepo) service.OrganizationService {
	return service.NewDefaultOrganizationService(organizationRepo)
}

func ProvideDefaultRegisterService(memberRepo repo.MemberRepo, organizationRepo repo.OrganizationRepo) service.RegisterService {
	return service.NewDefaultRegisterService(memberRepo, organizationRepo)
}
