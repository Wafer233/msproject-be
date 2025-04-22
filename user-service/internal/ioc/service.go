package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainService "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	"time"
)

func ProvideDefaultCaptchaService(cr repository.CaptchaRepository) service.CaptchaService {
	return service.NewDefaultCaptchaService(cr)
}

func ProvideDefaultAuthService(
	mr repository.MemberRepository,
	or repository.OrganizationRepository,
	ps *domainService.PasswordService,
	cr repository.CaptchaRepository,
	ts domainService.TokenService,
) service.AuthService {
	return service.NewDefaultAuthService(mr, or, ps, cr, ts)
}

func ProvidePasswordService() *domainService.PasswordService {
	return domainService.NewPasswordService()
}

func ProvideJWTTokenService(cfg *config.Config) domainService.TokenService {
	accessDur, _ := time.ParseDuration(cfg.JWT.AccessTokenDuration)
	refreshDur, _ := time.ParseDuration(cfg.JWT.RefreshTokenDuration)
	return domainService.NewJWTTokenService(cfg.JWT.SecretKey, accessDur, refreshDur)
}
