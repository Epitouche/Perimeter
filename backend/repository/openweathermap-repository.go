package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type OpenWeatherMapRepository interface{}

type openweathermapRepository struct {
	db *schemas.Database
}

// NewOpenWeatherMapRepository creates a new instance of OpenWeatherMapRepository
// with the provided gorm.DB connection. It initializes the internal database
// connection within the openweathermapRepository struct.
//
// Parameters:
//
//	conn - a pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//
//	An instance of OpenWeatherMapRepository.
func NewOpenWeatherMapRepository(conn *gorm.DB) OpenWeatherMapRepository {
	return &openweathermapRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
