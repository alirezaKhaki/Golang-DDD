package controller

import (
	"GIN_GORM/repository/postrepo"

	"github.com/gin-gonic/gin"
)

// Schema builds a graphql schema and returns it
func Schema(r *gin.Engine, controllers *Controllers) {

	r.GET("posts/list", func(c *gin.Context) {
		posts := controllers.postController.List()
		c.JSON(200, gin.H{
			"posts": posts,
		})
	})

	r.GET("posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		post := controllers.postController.FindOne(id)
		c.JSON(200, gin.H{
			"posts": post,
		})
	})

	r.DELETE("posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		result, e := controllers.postController.Delete(id)
		if e != nil {
			c.Status(400)
		}
		c.JSON(200, gin.H{
			"post": result,
		})
	})

	r.PUT("posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var body postrepo.PostBody
		c.Bind(&body)
		result := controllers.postController.Update(id, body)

		c.JSON(200, gin.H{
			"post": result,
		})
	})

	r.POST("posts", func(c *gin.Context) {
		var body postrepo.PostBody

		c.Bind(&body)
		result, e := controllers.postController.Create(body)

		if e != nil {
			c.Status(400)
		}
		c.JSON(200, gin.H{
			"post": result,
		})
	})
}
