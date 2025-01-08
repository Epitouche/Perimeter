package repository

import (
	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
)

type DropboxRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `DropboxRepository`.
type dropboxRepository struct {
	db *schemas.Database
}

func NewDropboxRepository(conn *gorm.DB) DropboxRepository {
	return &dropboxRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
