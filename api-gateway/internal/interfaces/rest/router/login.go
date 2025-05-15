package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
	getCaptchaHandler *handler.GetCaptchaHandler
	loginHandler      *handler.LoginHandler
	registerHandler   *handler.RegisterHandler
}

func NewLoginRouter(
	getCaptchaHandler *handler.GetCaptchaHandler,
	loginHandler *handler.LoginHandler,
	registerHandler *handler.RegisterHandler,
) *LoginRouter {
	return &LoginRouter{
		getCaptchaHandler: getCaptchaHandler,
		loginHandler:      loginHandler,
		registerHandler:   registerHandler,
	}
}

func (router *LoginRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")
	// Public routes don't need middleware
	group.POST("/login/getCaptcha", router.getCaptchaHandler.GetCaptcha)
	group.POST("/login/register", router.registerHandler.Register)
	group.POST("/login", router.loginHandler.Login)
}
