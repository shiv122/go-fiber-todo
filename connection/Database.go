package connection

import (
	"github.com/shiv122/go-todo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := config.DB.Username + ":" + config.DB.Password + "@tcp(" +
		config.DB.Host + ":" + config.DB.Port + ")/" +
		config.DB.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
}
