package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type TimerRepository interface{}

type timerRepository struct {
	db *schemas.Database
}

// NewTimerRepsository creates a new instance of TimerRepository with the provided gorm.DB connection.
// It initializes the timerRepository struct with a Database schema containing the given connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - TimerRepository: An interface representing the timer repository.
func NewTimerRepository(conn *gorm.DB) TimerRepository {
	return &timerRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}
