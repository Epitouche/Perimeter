package controller

import (
	"github.com/go-playground/validator/v10"

	"area/service"
)

type ServiceController interface {
}

type serviceController struct {
	service service.ServiceService
}

var validateService *validator.Validate

func NewServiceController(service service.ServiceService) ServiceController {
	validateService = validator.New()
	return &serviceController{
		service: service,
	}
}
