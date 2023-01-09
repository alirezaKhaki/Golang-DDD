package internal

import (
	"GIN_GORM/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("faild to connect to database")
	}

	env := os.Getenv("ENV")
	if env != "PRODUCTION" {
		DB.AutoMigrate(&models.Post{})
	}
}
