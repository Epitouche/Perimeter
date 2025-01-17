package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type OpenWeatherMapRepository interface{}

type openweathermapRepository struct {
	db *schemas.Database
}

func NewOpenWeatherMapRepository(conn *gorm.DB) OpenWeatherMapRepository {
	return &openweathermapRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
