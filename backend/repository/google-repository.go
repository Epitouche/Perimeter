package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type GoogleRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `GoogleRepository`.
type gmailRepository struct {
	db *schemas.Database
}

func NewGoogleRepository(conn *gorm.DB) GoogleRepository {
	return &gmailRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
