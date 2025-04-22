package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainService "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
)

func ProvideCachedCaptchaService(captchaRepo repository.CaptchaRepository) service.CaptchaService {
	return service.NewCachedCaptchaService(captchaRepo)
}

func ProvideCachedAuthService(
	mr repository.MemberRepository,
	or repository.OrganizationRepository,
	ps *domainService.PasswordService,
	cr repository.CaptchaRepository,
) service.AuthService {
	return service.NewCachedAuthService(mr, or, ps, cr)
}

func ProvidePasswordService() *domainService.PasswordService {
	return domainService.NewPasswordService()
}
