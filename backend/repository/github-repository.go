package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type GithubRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `GithubRepository`.
type githubRepository struct {
	db *schemas.Database
}

// NewGithubRepository creates a new instance of GithubRepository with the provided gorm.DB connection.
// It initializes the database connection within the githubRepository struct.
//
// Parameters:
//
//	conn - a gorm.DB connection to be used by the repository.
//
// Returns:
//
//	A new instance of GithubRepository.
func NewGithubRepository(conn *gorm.DB) GithubRepository {
	return &githubRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
