package main

import (
	"github.com/Wafer233/msproject-be/user-service/internal/ioc"
)

func main() {

	app := ioc.InitApp()

	server := app.Server

	server.Run(":80")
}
