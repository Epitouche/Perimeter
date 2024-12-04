package service

import (
	"time"

	"area/repository"
)

type TimerService interface {
	timerActionSpecificHour(c chan string, hour int, minute int)
}

type timerService struct {
	repository repository.TimerRepository
}

func NewTimerService(repository repository.TimerRepository) TimerService {
	return &timerService{
		repository: repository,
	}
}

func (service *timerService) timerActionSpecificHour(c chan string, hour int, minute int) {
	dt := time.Now().Local()
	if dt.Hour() == hour && dt.Minute() == minute {
		println("current time is ", dt.String())
		c <- "ok" // send sum to c
	}
}
