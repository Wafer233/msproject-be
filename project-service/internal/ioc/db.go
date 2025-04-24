package ioc

import (
	"fmt"
	"github.com/Wafer233/msproject-be/project-service/config"
	"github.com/Wafer233/msproject-be/project-service/internal/infrastructure/persistence/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to MySQL: " + err.Error())
	}

	// 可选自动迁移
	if err := entity.InitTable(db); err != nil {
		panic("failed to auto-migrate tables: " + err.Error())
	}

	return db
}
