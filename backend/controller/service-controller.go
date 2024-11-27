package controller

import (
	"github.com/go-playground/validator/v10"

	"area/service"
)

type ServiceController interface {
}

type serviceController struct {
	service     service.ServiceService
	serviceUser service.UserService
}

var validateService *validator.Validate

func NewServiceController(service service.ServiceService, serviceUser service.UserService) ServiceController {
	validateService = validator.New()
	return &serviceController{
		service:     service,
		serviceUser: serviceUser,
	}
}
