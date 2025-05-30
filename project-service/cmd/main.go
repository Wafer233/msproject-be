package main

import (
	"github.com/Wafer233/msproject-be/project-service/internal/ioc"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// grpc
func main() {
	app := ioc.InitApp()

	// 启动gRPC服务器
	go func() {
		if err := app.GrpcServer.Start(); err != nil {
			zap.L().Fatal("Failed to start gRPC server", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutting down gRPC server...")
	app.GrpcServer.Stop()
	zap.L().Info("Server exited")
}
