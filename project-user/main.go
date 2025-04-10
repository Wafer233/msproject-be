package main

import (
	srv "github.com/Wafer233/msproject-be/project-common"
	"github.com/Wafer233/msproject-be/project-user/config"
	"github.com/Wafer233/msproject-be/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	gc := router.RegisterGrpc()

	stop := func() {
		gc.Stop()
	}

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
