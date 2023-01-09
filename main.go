package main

import (
	"GIN_GORM/controller"
	"GIN_GORM/internal"

	"github.com/gin-gonic/gin"
)

func init() {
	internal.LoadEnv()
	internal.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.han
	r.POST("/posts", controller.Create)
	r.GET("/posts/list", controller.List)
	r.GET("/posts/:id", controller.FindOne)
	r.PUT("/posts/:id", controller.Update)
	r.DELETE("/posts/:id", controller.Delete)
	r.Run()

}
