package service

import (
	"time"

	"area/repository"
	"area/schemas"
)

type TimerService interface {
	TimerActionSpecificHour(c chan string, option string)
	TimerReactionGiveTime(option string)
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option string)
	FindReactionbyName(name string) func(option string)
}

type timerService struct {
	repository repository.TimerRepository
}

func NewTimerService(repository repository.TimerRepository) TimerService {
	return &timerService{
		repository: repository,
	}
}

func (service *timerService) FindActionbyName(name string) func(c chan string, option string) {
	switch name {
	case string(schemas.SpecificTime):
		return service.TimerActionSpecificHour
	default:
		return nil
	}
}

func (service *timerService) FindReactionbyName(name string) func(option string) {
	switch name {
	case string(schemas.GiveTime):
		return service.TimerReactionGiveTime
	default:
		return nil
	}
}

func (service *timerService) TimerActionSpecificHour(c chan string, option string) {
	dt := time.Now().Local()
	if dt.Hour() == 19 && dt.Minute() == 15 {
		println("current time is ", dt.String())
		c <- "response" // send sum to c
	}
	time.Sleep(15 * time.Second)
}

func (service *timerService) TimerReactionGiveTime(option string) {
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
