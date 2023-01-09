package repository

import (
	"GIN_GORM/repository/postrepo"

	"gorm.io/gorm"
)

// Repositories contains all the repo structs
type Repositories struct {
	PostRepo *postrepo.PostRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB) *Repositories {
	postRepo := postrepo.NewPostRepo(db)
	return &Repositories{PostRepo: postRepo}
}
