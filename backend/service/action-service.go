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
	serviceTimer   TimerService
}

func NewActionService(
	repository repository.ActionRepository,
	serviceService ServiceService,
	serviceTimer TimerService,
) ActionService {
	return &actionService{
		repository:     repository,
		serviceService: serviceService,
		serviceTimer:   serviceTimer,
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
