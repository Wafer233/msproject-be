package router

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/handler"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	captchaHandler  *handler.CaptchaHandler
	loginHandler    *handler.LoginHandler
	registerHandler *handler.RegisterHandler
}

func NewAuthRouter(ch *handler.CaptchaHandler,
	lr *handler.LoginHandler,
	rh *handler.RegisterHandler,
) *AuthRouter {
	return &AuthRouter{
		captchaHandler:  ch,
		loginHandler:    lr,
		registerHandler: rh,
	}
}

func (ar *AuthRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")
	group.POST("/login/getCaptcha", ar.captchaHandler.GetCaptcha)
	group.POST("/login/register", ar.registerHandler.Register)
	//group.POST("/login", ar.loginHandler.Login)
}
