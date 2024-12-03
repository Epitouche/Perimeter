package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type GithubRepository interface {
}

// Define a struct that embeds `*schemas.Database` and implements `GithubRepository`.
type githubRepository struct {
	db *schemas.Database
}

func NewGithubRepository(conn *gorm.DB) GithubRepository {
	return &githubRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
