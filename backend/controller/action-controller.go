package controller

import (
	"fmt"

	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/service"
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

func (controller *actionController) GetActionsInfo(
	id uint64,
) (response []schemas.Action, err error) {
	response, err = controller.service.GetActionsInfo(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get actions info because %w", err)
	}
	return response, nil
}
