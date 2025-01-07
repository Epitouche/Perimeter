package repository

import (
	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
)

type GmailRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `GmailRepository`.
type gmailRepository struct {
	db *schemas.Database
}

func NewGmailRepository(conn *gorm.DB) GmailRepository {
	return &gmailRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
