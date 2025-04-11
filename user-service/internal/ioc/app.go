package ioc

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Server *gin.Engine
	// 可能还有其他东西
}
