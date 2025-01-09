package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type MicrosoftRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `MicrosoftRepository`.
type microsoftRepository struct {
	db *schemas.Database
}

func NewMicrosoftRepository(conn *gorm.DB) MicrosoftRepository {
	return &microsoftRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
