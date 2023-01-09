package controller

import (
	"GIN_GORM/internal"
	"GIN_GORM/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// get data from body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//create data
	post := models.Post{Title: body.Title, Body: body.Body}

	result := internal.DB.Create((&post))

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return response
	c.JSON(200, gin.H{
		"result": body,
	})
}

func List(c *gin.Context) {
	var posts []models.Post
	internal.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func FindOne(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	internal.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	internal.DB.First(&post, id)
	internal.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	internal.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}
