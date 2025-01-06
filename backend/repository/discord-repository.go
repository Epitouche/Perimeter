package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type DiscordRepository interface{}

// Define a struct that embeds `*schemas.Database` and implements `DiscordRepository`.
type discordRepository struct {
	db *schemas.Database
}

func NewDiscordRepository(conn *gorm.DB) DiscordRepository {
	return &discordRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
