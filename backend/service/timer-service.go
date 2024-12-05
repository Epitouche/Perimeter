package service

import (
	"time"

	"area/repository"
	"area/schemas"
)

type TimerService interface {
	TimerActionSpecificHour(c chan string, hour int, minute int)
	TimerReactionGiveTime()
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
}

type timerService struct {
	repository repository.TimerRepository
}

func NewTimerService(repository repository.TimerRepository) TimerService {
	return &timerService{
		repository: repository,
	}
}

func (service *timerService) TimerActionSpecificHour(c chan string, hour int, minute int) {
	dt := time.Now().Local()
	if dt.Hour() == hour && dt.Minute() == minute {
		println("current time is ", dt.String())
		c <- "ok"
	}
}

func (service *timerService) TimerReactionGiveTime() {
	println("give time")
}

func (service *timerService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{
		{
			Name:        string(schemas.SpecificTime),
			Description: "This action is a specific time action",
			Service:     schemas.Service{Name: schemas.Timer},
			Option:      "{hour: 0, minute: 0}",
		},
	}
}

func (service *timerService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{
		{
			Name:        string(schemas.GiveTime),
			Description: "This reaction is a give time reaction",
			Service:     schemas.Service{Name: schemas.Timer},
			Option:      "{}",
		},
	}
}