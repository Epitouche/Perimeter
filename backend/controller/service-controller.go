package controller

import (
	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type ServiceController interface {
	AboutJson(ctx *gin.Context) (allService []schemas.ServiceJson, err error)
}

type serviceController struct {
	service service.ServiceService
}

func NewServiceController(service service.ServiceService) ServiceController {
	return &serviceController{
		service: service,
	}
}

func (controller *serviceController) AboutJson(ctx *gin.Context) (allService []schemas.ServiceJson, err error) {
	allServices, err := controller.service.GetAllServices()
	return allServices, nil
}
