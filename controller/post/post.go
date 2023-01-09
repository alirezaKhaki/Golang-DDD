package post

import (
	"GIN_GORM/models"
	"GIN_GORM/repository/postrepo"
)

func (co *Controller) Create(body postrepo.PostBody) (models.Post, error) {
	return co.service.Create(models.Post{Title: body.Title, Body: body.Body})
}

func (co *Controller) List() []models.Post {
	return co.service.List()
}

func (co *Controller) FindOne(id string) models.Post {
	return co.service.FindOne(id)
}

func (co *Controller) Update(id string, body postrepo.PostBody) models.Post {
	return co.service.Update(id, body)
}

func (co *Controller) Delete(id string) (models.Post, error) {
	return co.service.Delete(id)
}
