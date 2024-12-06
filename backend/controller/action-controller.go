package controller

import (
	"area/schemas"
	"area/service"
)

type ActionController interface {
	GetActionsInfo(id uint64) (response []schemas.Action, err error)
}

type actionController struct {
	service service.ActionService
}

func NewActionController(service service.ActionService) ActionController {
	return &actionController{
		service: service,
	}
}

func (controller *actionController) GetActionsInfo(id uint64) (response []schemas.Action, err error) {
	return controller.service.GetActionsInfo(id)
}