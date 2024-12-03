package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type SpotifyRepository interface {
	Save(token schemas.SpotifyToken)
	Update(token schemas.SpotifyToken)
	Delete(token schemas.SpotifyToken)
	FindAll() []schemas.SpotifyToken
	FindByAccessToken(accessToken string) []schemas.SpotifyToken
	FindById(id uint64) schemas.SpotifyToken
}

// Define a struct that embeds `*schemas.Database` and implements `SpotifyRepository`.
type spotifyRepository struct {
	db *schemas.Database
}

func NewSpotifyRepository(conn *gorm.DB) SpotifyRepository {
	err := conn.AutoMigrate(&schemas.SpotifyToken{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &spotifyRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *spotifyRepository) Save(token schemas.SpotifyToken) {
	err := repo.db.Connection.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *spotifyRepository) Update(token schemas.SpotifyToken) {
	err := repo.db.Connection.Save(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *spotifyRepository) Delete(token schemas.SpotifyToken) {
	err := repo.db.Connection.Delete(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *spotifyRepository) FindAll() []schemas.SpotifyToken {
	var tokens []schemas.SpotifyToken
	err := repo.db.Connection.Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *spotifyRepository) FindByAccessToken(accessToken string) []schemas.SpotifyToken {
	var tokens []schemas.SpotifyToken
	err := repo.db.Connection.Where(&schemas.SpotifyToken{AccessToken: accessToken}).Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *spotifyRepository) FindById(id uint64) schemas.SpotifyToken {
	var tokens schemas.SpotifyToken
	err := repo.db.Connection.Where(&schemas.SpotifyToken{Id: id}).First(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}
