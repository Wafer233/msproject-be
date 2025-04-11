package router

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest"
	"github.com/gin-gonic/gin"
)

// LoginRouter 登录相关路由
type LoginRouter struct {
	loginHandler *rest.LoginHandler
}

// NewLoginRouter 创建登录路由
func NewLoginRouter(loginHandler *rest.LoginHandler) *LoginRouter {
	return &LoginRouter{
		loginHandler: loginHandler,
	}
}

// Register 注册路由
func (r *LoginRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project/login")

	group.POST("/getCaptcha", r.loginHandler.GetCaptcha)
	// 其他登录相关路由...
}
