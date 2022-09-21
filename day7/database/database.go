package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dakasakti/day2/config"
	"github.com/dakasakti/day2/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func InitMongoDB(config *config.AppConfig) *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(config.Mongo_URI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func InitCollection(client *mongo.Client, config *config.AppConfig, name string) *mongo.Collection {
	collection := client.Database(config.Mongo_Database).Collection(name)
	return collection
}
