package postrepo

import (
	"GIN_GORM/models"

	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

type PostBody struct {
	Body  string
	Title string
}


func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

// GetExistingUser fetches a user by the username from the db and returns it.
func (repo *PostRepo) FindOne(id string) models.Post {
	var Post models.Post
	repo.db.Order("id desc").Where("id = ?", id).First(&Post)
	return Post
}

func (repo *PostRepo) List() []models.Post {
	var Post []models.Post
	repo.db.Order("id desc").Find(&Post)
	return Post
}

func (repo *PostRepo) Update(id string, body PostBody) models.Post {

	var post models.Post
	repo.FindOne(id)
	repo.db.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	return post
}

func (repo *PostRepo) Create(Post models.Post) (models.Post, error) {
	// TODO handle the potential error below.

	result := repo.db.Create(&Post)

	if result.Error != nil {
		return Post, result.Error
	}
	return Post, nil
}

func (repo *PostRepo) Delete(id string) (models.Post, error) {
	// TODO handle the potential error below.
	var post models.Post
	result := repo.db.Delete(&post, id)

	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}
