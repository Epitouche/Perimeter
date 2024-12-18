package service

import (
	"area/repository"
	"area/schemas"
)

// Constructor

type OpenweathermapService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
	GetActionsName() []string
	GetReactionsName() []string
	// Service specific functions
	// Actions functions
	// Reactions functions
}

type openweathermapService struct {
	repository        repository.OpenweathermapRepository
	serviceRepository repository.ServiceRepository
	actionsName       []string
	reactionsName     []string
	serviceInfo       schemas.Service
}

func NewOpenweathermapService(
	repository repository.OpenweathermapRepository,
	serviceRepository repository.ServiceRepository,
) OpenweathermapService {
	return &openweathermapService{
		repository:        repository,
		serviceRepository: serviceRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Openweathermap,
			Description: "This service is a weather service",
		},
	}
}

// Service interface functions

func (service *openweathermapService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *openweathermapService) FindActionbyName(
	name string,
) func(c chan string, option string, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *openweathermapService) FindReactionbyName(name string) func(option string, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *openweathermapService) GetServiceActionInfo() []schemas.Action {
	service.actionsName = append(service.actionsName, string(schemas.SpecificTime))
	return []schemas.Action{
		{
			Name:        string(schemas.SpecificTime),
			Description: "This action is a specific time action",
			Service:     service.serviceRepository.FindByName(schemas.Openweathermap),
			Option:      "{\"hour\": 0, \"minute\": 0}",
		},
	}
}

func (service *openweathermapService) GetServiceReactionInfo() []schemas.Reaction {
	service.reactionsName = append(service.reactionsName, string(schemas.GiveTime))
	return []schemas.Reaction{
		{
			Name:        string(schemas.GiveTime),
			Description: "This reaction is a give time reaction",
			Service:     service.serviceRepository.FindByName(schemas.Openweathermap),
			Option:      "{}",
		},
	}
}

func (service *openweathermapService) GetActionsName() []string {
	return service.actionsName
}

func (service *openweathermapService) GetReactionsName() []string {
	return service.reactionsName
}

// Service specific functions

// Actions functions

// Reactions functions
