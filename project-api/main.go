package main

import (
	"github.com/Wafer233/msproject-be/project-api/config"
	"github.com/Wafer233/msproject-be/project-api/router"
	srv "github.com/Wafer233/msproject-be/project-common"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, config.AppConf.AppConfig.Name, config.AppConf.AppConfig.Addr, nil)
}
