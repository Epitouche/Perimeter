package service

import (
	"area/repository"
	"area/schemas"
)

type ActionService interface {
	Save(newService schemas.Action) error
	Update(newService schemas.Action) error
	Delete(newService schemas.Action) error
	FindAll() []schemas.Action
}

type actionService struct {
	repository repository.ActionRepository
	action     []schemas.Action
}

func NewActionService(repository repository.ActionRepository) ActionService {
	newService := actionService{
		repository: repository,
		action: []schemas.Action{
			{
				Name:        "Spotify",
				Description: "This service is a music service",
			},
			{
				Name:        "OpenWeatherMap",
				Description: "This service is a weather service",
			},
			{
				Name:        "Time",
				Description: "This service is a time service",
			},
		},
	}
	newService.InitialSaveService()
	return &newService
}

func (service *actionService) InitialSaveService() {
	for _, oneService := range service.action {
		serviceByName := service.repository.FindByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
		}
	}
}

func (service *actionService) Save(newService schemas.Action) error {
	service.repository.Save(newService)
	return nil
}

func (service *actionService) Update(newService schemas.Action) error {
	service.repository.Update(newService)
	return nil
}

func (service *actionService) Delete(newService schemas.Action) error {
	service.repository.Delete(newService)
	return nil
}

func (service *actionService) FindAll() []schemas.Action {
	return service.repository.FindAll()
}
