package main

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/ioc"
	"go.uber.org/zap"
)

func main() {
	app, err := ioc.InitApp()
	if err != nil {
		zap.L().Fatal("Failed to initialize application", zap.Error(err))
	}

	// Start HTTP server
	if err := app.Server.Run(":80"); err != nil {
		zap.L().Fatal("Failed to start HTTP server", zap.Error(err))
	}
}
