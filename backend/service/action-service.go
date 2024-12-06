package service

import (
	"fmt"

	"area/repository"
	"area/schemas"
)

type ActionService interface {
	FindAll() []schemas.Action
	SaveAllAction()
	FindById(actionId uint64) schemas.Action
	GetActionsInfo(id uint64) (response []schemas.Action, err error)
	GetAllServicesByServiceId(serviceId uint64) (actionJson []schemas.ActionJson)
}

type ServiceAction interface {
	GetServiceActionInfo() []schemas.Action
}

type actionService struct {
	repository     repository.ActionRepository
	serviceService ServiceService
}

func NewActionService(
	repository repository.ActionRepository,
	serviceService ServiceService,
) ActionService {
	newActionService := &actionService{
		repository:     repository,
		serviceService: serviceService,
	}
	newActionService.SaveAllAction()
	return newActionService
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

func (service *actionService) SaveAllAction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceAction, ok := services.(ServiceAction); ok {
			actions := serviceAction.GetServiceActionInfo()
			for _, action := range actions {
				actionByName := service.repository.FindByName(action.Name)
				if len(actionByName) == 0 {
					service.repository.Save(action)
				}
			}
		} else {
			fmt.Println("Service is not ServiceAction")
		}
	}
}

func (service *actionService) FindById(actionId uint64) schemas.Action {
	return service.repository.FindById(actionId)
}

func (service *actionService) GetActionsInfo(id uint64) (response []schemas.Action, err error) {
	return service.repository.FindByServiceId(id), nil
}
