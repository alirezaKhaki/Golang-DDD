package main

import (
	"GIN_GORM/controller"
	"GIN_GORM/internal"
	"GIN_GORM/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	internal.LoadEnv()
	internal.ConnectToDb()
	repository.InitRepositories(internal.DB)
}

func main() {
	r := gin.Default()

	repos := repository.InitRepositories(internal.DB)
	controllers := controller.InitControllers((repos))
	controller.Schema(r, controllers)
	r.Run()
}
