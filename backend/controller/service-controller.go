package controller

import (
	"area/service"
)

type ServiceController interface{}

type serviceController struct {
	service service.ServiceService
}

func NewServiceController(service service.ServiceService) ServiceController {
	return &serviceController{
		service: service,
	}
}
