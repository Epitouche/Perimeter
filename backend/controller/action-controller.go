package controller

import (
	"area/service"
)

type ActionController interface {
}

type actionController struct {
	service service.ActionService
}

func NewActionController(service service.ActionService) ActionController {
	return &actionController{
		service: service,
	}
}
