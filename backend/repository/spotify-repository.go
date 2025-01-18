package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type SpotifyRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `SpotifyRepository`.
type spotifyRepository struct {
	db *schemas.Database
}

// NewSpotifyRepository creates a new instance of SpotifyRepository with the provided database connection.
// It initializes the spotifyRepository struct with a Database schema containing the given gorm.DB connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - SpotifyRepository: An interface representing the Spotify repository.
func NewSpotifyRepository(conn *gorm.DB) SpotifyRepository {
	return &spotifyRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
