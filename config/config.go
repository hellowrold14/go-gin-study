package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

var config Config
var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	// Load environment variables from .env file
	config = Config{
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
	}

	db = InitDB(config)
}

func GetDB() *gorm.DB {
	return db
}

func GetConfig() Config {
	return config
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}
