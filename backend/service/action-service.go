package service

import (
	"area/repository"
	"area/schemas"
)

type ActionService interface {
	FindAll() []schemas.Action
	GetAllServicesByServiceId(serviceId uint64) (actionJson []schemas.ActionJson)
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

func (service *actionService) FindAll() []schemas.Action {
	return service.repository.FindAll()
}

func (service *actionService) GetAllServicesByServiceId(
	serviceId uint64,
) (actionJson []schemas.ActionJson) {
	allActionForService := service.repository.FindByServiceId(serviceId)
	for _, oneAction := range allActionForService {
		actionJson = append(actionJson, schemas.ActionJson{
			Name:        oneAction.Name,
			Description: oneAction.Description,
		})
	}
	return actionJson
}
