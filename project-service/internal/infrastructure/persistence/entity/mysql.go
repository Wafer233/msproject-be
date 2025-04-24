package entity

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC")
	err := db.AutoMigrate(&ProjectMenuEntity{})
	if err != nil {
		zap.L().Error("自动迁移数据库失败", zap.Error(err))
		return err
	}
	return nil
}
