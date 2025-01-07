package repository

import (
	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
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
