package service

import (
	"time"

	"area/repository"
	"area/schemas"
)

type TimerService interface {
	TimerActionSpecificHour(c chan string, option string, idArea uint64)
	TimerReactionGiveTime(option string, idArea uint64)
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
}

type timerService struct {
	repository        repository.TimerRepository
	serviceRepository repository.ServiceRepository
}

func NewTimerService(
	repository repository.TimerRepository,
	serviceRepository repository.ServiceRepository,
) TimerService {
	return &timerService{
		repository:        repository,
		serviceRepository: serviceRepository,
	}
}

func (service *timerService) FindActionbyName(
	name string,
) func(c chan string, option string, idArea uint64) {
	switch name {
	case string(schemas.SpecificTime):
		return service.TimerActionSpecificHour
	default:
		return nil
	}
}

func (service *timerService) FindReactionbyName(name string) func(option string, idArea uint64) {
	switch name {
	case string(schemas.GiveTime):
		return service.TimerReactionGiveTime
	default:
		return nil
	}
}

func (service *timerService) TimerActionSpecificHour(c chan string, option string, idArea uint64) {
	dt := time.Now().Local()
	if dt.Hour() == 19 && dt.Minute() == 25 {
		println("current time is ", dt.String())
		c <- "response" // send sum to c
	}
	time.Sleep(15 * time.Second)
}

func (service *timerService) TimerReactionGiveTime(option string, idArea uint64) {
	println("give time")
}

func (service *timerService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{
		{
			Name:        string(schemas.SpecificTime),
			Description: "This action is a specific time action",
			Service:     service.serviceRepository.FindByName(schemas.Timer),
			Option:      "{hour: 0, minute: 0}",
		},
	}
}

func (service *timerService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{
		{
			Name:        string(schemas.GiveTime),
			Description: "This reaction is a give time reaction",
			Service:     service.serviceRepository.FindByName(schemas.Timer),
			Option:      "{}",
		},
	}
}
