package main

import (
	srv "github.com/Wafer233/msproject-be/project-common"
	"github.com/Wafer233/msproject-be/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, "webcenter", ":80")
}
