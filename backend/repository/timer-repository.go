package repository

import (
	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
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
