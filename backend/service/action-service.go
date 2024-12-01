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
	repository     repository.ActionRepository
	serviceService ServiceService
	allAction      map[schemas.ServiceName][]schemas.Action
}

func NewActionService(
	repository repository.ActionRepository,
	serviceService ServiceService,
) ActionService {
	newService := actionService{
		repository:     repository,
		serviceService: serviceService,
		allAction: map[schemas.ServiceName][]schemas.Action{
			schemas.Spotify: {
				{
					Name:        "action1",
					Description: "do something",
				},
				{
					Name:        "action2",
					Description: "do something",
				},
				{
					Name:        "action3",
					Description: "do something",
				},
			},
			schemas.OpenWeatherMap: {
				{
					Name:        "action1",
					Description: "do something",
				},
				{
					Name:        "action2",
					Description: "do something",
				},
				{
					Name:        "action3",
					Description: "do something",
				},
			},
			schemas.Timer: {
				{
					Name:        "action1",
					Description: "do something",
				},
				{
					Name:        "action2",
					Description: "do something",
				},
				{
					Name:        "action3",
					Description: "do something",
				},
			},
		},
	}
	newService.InitialSaveAction()
	return &newService
}

func (service *actionService) InitialSaveAction() {
	allService := service.serviceService.FindAll()
	// Find all service and save action
	for _, oneService := range allService {
		// Find all action by service name
		for _, oneAction := range service.allAction[schemas.ServiceName(oneService.Name)] {
			existingActions := service.repository.FindByServiceByName(oneService.Id, oneAction.Name)

			if len(existingActions) == 0 {
				oneAction.Service = oneService
				service.repository.Save(oneAction)
			}
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
