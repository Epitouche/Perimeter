package repository

import (
	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
)

type SpotifyRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `SpotifyRepository`.
type spotifyRepository struct {
	db *schemas.Database
}

func NewSpotifyRepository(conn *gorm.DB) SpotifyRepository {
	return &spotifyRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
