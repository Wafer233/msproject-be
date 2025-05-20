package entity

import (
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC")
	return db.AutoMigrate(&Member{}, &Organization{})
}
