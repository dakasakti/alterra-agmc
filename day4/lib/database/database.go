package database

import (
	"fmt"
	"log"

	"github.com/dakasakti/day2/config"
	"github.com/dakasakti/day2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(config *config.AppConfig) *gorm.DB {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Address,
		config.DB_Port,
		config.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func InitMigrate() {
	db := InitMySQL(config.GetConfig())
	db.AutoMigrate(&models.User{})
}
