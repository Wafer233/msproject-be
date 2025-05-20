package main

import (
	"github.com/Wafer233/msproject-be/user-service/internal/ioc"
	"log"
)

func main() {
	app := ioc.InitApp()

	if err := app.GrpcServer.Start(); err != nil {
		log.Fatal(err)
	}
}
