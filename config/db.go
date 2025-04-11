package config

import (
	"github.com/cody/go-demo1/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config Config) *gorm.DB {
	dsn := config.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to auto migrate")
	}
	return db
}
