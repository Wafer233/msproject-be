package router

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/handler"
	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct {
	ch *handler.CaptchaHandler
}

func NewCaptchaRouter(ch *handler.CaptchaHandler) *CaptchaRouter {
	return &CaptchaRouter{ch: ch}
}

func (cr *CaptchaRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project/login")
	group.POST("/getCaptcha", cr.ch.GetCaptcha)
}
