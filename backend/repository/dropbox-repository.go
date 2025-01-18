package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type DropboxRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `DropboxRepository`.
type dropboxRepository struct {
	db *schemas.Database
}

// NewDropboxRepository creates a new instance of DropboxRepository with the provided gorm.DB connection.
// It initializes the internal database connection using the given gorm.DB instance.
//
// Parameters:
//
//	conn - a gorm.DB instance representing the database connection.
//
// Returns:
//
//	A new instance of DropboxRepository.
func NewDropboxRepository(conn *gorm.DB) DropboxRepository {
	return &dropboxRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
