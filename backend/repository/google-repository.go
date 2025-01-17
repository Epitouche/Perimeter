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

// NewGoogleRepository creates a new instance of GoogleRepository with the provided gorm.DB connection.
// It initializes a gmailRepository with a schemas.Database containing the given connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - GoogleRepository: A new instance of GoogleRepository.
func NewGoogleRepository(conn *gorm.DB) GoogleRepository {
	return &gmailRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
