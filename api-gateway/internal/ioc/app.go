package ioc

import (
	"github.com/gin-gonic/gin"
)

// no need to change since in api-gateway only need rest interfaces

type App struct {
	Server *gin.Engine
	// Could add more fields if needed
}
