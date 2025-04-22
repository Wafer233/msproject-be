package router

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/handler"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	ch *handler.CaptchaHandler
	lh *handler.LoginHandler
	rh *handler.RegisterHandler
}

func NewAuthRouter(
	ch *handler.CaptchaHandler,
	lh *handler.LoginHandler,
	rh *handler.RegisterHandler,
) *AuthRouter {
	return &AuthRouter{
		ch: ch,
		lh: lh,
		rh: rh,
	}
}

func (ar *AuthRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")
	group.POST("/login/getCaptcha", ar.ch.GetCaptcha)
	group.POST("/login/register", ar.rh.Register)
	//group.POST("/login", ar.loginHandler.Login)
}
