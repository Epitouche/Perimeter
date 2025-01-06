package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type TimerRepository interface{}

type timerRepository struct {
	db *schemas.Database
}

func NewTimerRepository(conn *gorm.DB) TimerRepository {
	return &timerRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
