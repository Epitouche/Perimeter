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

// NewMicrosoftRepository creates a new instance of MicrosoftRepository with the provided database connection.
// It initializes the microsoftRepository struct with a schemas.Database that holds the given gorm.DB connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - MicrosoftRepository: A new instance of MicrosoftRepository.
func NewMicrosoftRepository(conn *gorm.DB) MicrosoftRepository {
	return &microsoftRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
