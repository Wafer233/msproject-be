package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.Engine)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	for _, route := range routers {
		route.Register(r)
	}
}

func RegisterRouter(r ...Router) {
	routers = append(routers, r...)
}
