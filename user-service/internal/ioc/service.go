package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
)

func ProvideCachedCaptchaService(captchaRepo repository.CaptchaRepository) service.CaptchaService {
	return service.NewCachedCaptchaService(captchaRepo)
}
