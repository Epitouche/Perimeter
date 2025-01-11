package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type GooglelRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `GooglelRepository`.
type gmailRepository struct {
	db *schemas.Database
}

func NewGoogleRepository(conn *gorm.DB) GooglelRepository {
	return &gmailRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
