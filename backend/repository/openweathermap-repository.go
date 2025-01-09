package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type OpenweathermapRepository interface{}

type openweathermapRepository struct {
	db *schemas.Database
}

func NewOpenweathermapRepository(conn *gorm.DB) OpenweathermapRepository {
	return &openweathermapRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
