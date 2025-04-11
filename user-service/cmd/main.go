package main

import (
	"github.com/Wafer233/msproject-be/common"
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest"
	"github.com/go-redis/redis/v8"

	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/router"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化配置
	// 已由config包加载

	// 初始化数据库连接
	//db, err := gorm.Open(mysql.Open(config.C.MySQL.DSN), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database: " + err.Error())
	//}

	// 初始化Redis客户端
	rdb := redisClient.NewClient(&redisClient.Options{
		Addr:     config.C.Redis.Addr,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,
	})

	// 初始化缓存
	redisCache := redis.NewRedisClient(rdb)

	// 初始化仓储
	captchaRepo := repository.NewRedisCaptchaRepository(redisCache)
	userRepo := repository.NewMysqlUserRepository(db)

	// 初始化应用服务
	captchaService := service.NewCaptchaService(captchaRepo)
	userService := service.NewUserService(userRepo, captchaRepo)

	// 初始化处理器
	loginHandler := rest.NewLoginHandler(captchaService, userService)

	// 初始化路由
	r := gin.Default()
	loginRouter := router.NewLoginRouter(loginHandler)
	loginRouter.Register(r)

	// 注册其他路由...

	// 启动服务
	common.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
