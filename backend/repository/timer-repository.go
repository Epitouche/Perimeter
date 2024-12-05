package repository

type TimerRepository interface{}

type timerRepository struct{}

func NewTimerRepository() TimerRepository {
	return &timerRepository{}
}
