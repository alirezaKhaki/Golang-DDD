package post

import (
	"GIN_GORM/models"
	"GIN_GORM/repository/postrepo"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type PostRepository interface {
	FindOne(id string) models.Post
	List()[]models.Post
	Update(id string, body postrepo.PostBody) models.Post
	Create(Post models.Post) (models.Post, error)
	Delete(id string) (models.Post, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	service PostRepository
}

// InitController initializes the user controller.
func InitController(postRepo *postrepo.PostRepo) *Controller {
	return &Controller{
		service: postRepo,
	}
}
