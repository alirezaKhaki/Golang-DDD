package controller

import (
	"GIN_GORM/controller/post"
	"GIN_GORM/repository"
)

type Controllers struct {
	postController *post.Controller
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		postController: post.InitController(repositories.PostRepo),
	}
}

