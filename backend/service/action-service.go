package service

import (
	"fmt"

	"area/repository"
	"area/schemas"
)

type ActionService interface {
	FindAll() (actions []schemas.Action, err error)
	SaveAllAction()
	FindById(actionId uint64) (action schemas.Action, err error)
	GetActionsInfo(id uint64) (response []schemas.Action, err error)
	GetAllServicesByServiceId(serviceId uint64) (actionJSON []schemas.ActionJSON)
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

func (service *actionService) FindAll() (actions []schemas.Action, err error) {
	actions, err = service.repository.FindAll()
	if err != nil {
		return actions, fmt.Errorf("error when get all actions: %w", err)
	}
	return actions, nil
}

func (service *actionService) GetAllServicesByServiceId(
	serviceId uint64,
) (actionJSON []schemas.ActionJSON) {
	allActionForService, err := service.repository.FindByServiceId(serviceId)
	if err != nil {
		fmt.Println("Error when get all actions by service id")
	}
	for _, oneAction := range allActionForService {
		actionJSON = append(actionJSON, schemas.ActionJSON{
			Name:        oneAction.Name,
			Description: oneAction.Description,
		})
	}

	return actionJSON
}

func (service *actionService) SaveAllAction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceAction, ok := services.(ServiceAction); ok {
			actions := serviceAction.GetServiceActionInfo()
			for _, action := range actions {
				actionByName, err := service.repository.FindByName(action.Name)
				if err != nil {
					fmt.Println("Error when get action by name")
				}
				if len(actionByName) == 0 {
					service.repository.Save(action)
				}
			}
		} else {
			fmt.Println("Service is not ServiceAction")
		}
	}
}

func (service *actionService) FindById(actionId uint64) (action schemas.Action, err error) {
	action, err = service.repository.FindById(actionId)
	if err != nil {
		return action, fmt.Errorf("error when get action by id: %w", err)
	}
	return action, nil
}

func (service *actionService) GetActionsInfo(id uint64) (response []schemas.Action, err error) {
	response, err = service.repository.FindByServiceId(id)
	if err != nil {
		return response, fmt.Errorf("error when get actions info: %w", err)
	}
	return response, nil
}
