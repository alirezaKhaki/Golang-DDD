package main

import (
	"GIN_GORM/internal"
	"GIN_GORM/models"
)

func init() {
	internal.LoadEnv()
	internal.ConnectToDb()
}

func main() {
	internal.DB.AutoMigrate((&models.Post{}))
}
