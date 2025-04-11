package main

import (
	"github.com/Wafer233/msproject-be/user-service/internal/ioc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	app := ioc.InitApp()

	server := app.Server

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello webook！")
	})

	server.Run(":80")
}
