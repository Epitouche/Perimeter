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
	service       service.ServiceService
	serviceAction service.ActionService
}

func NewServiceController(service service.ServiceService, serviceAction service.ActionService) ServiceController {
	return &serviceController{
		service:       service,
		serviceAction: serviceAction,
	}
}

func (controller *serviceController) AboutJson(ctx *gin.Context) (allServicesJson []schemas.ServiceJson, err error) {
	allServices := controller.service.FindAll()
	for _, oneService := range allServices {
		allServicesJson = append(allServicesJson, schemas.ServiceJson{
			Name:   schemas.ServiceName(oneService.Name),
			Action: controller.serviceAction.GetAllServicesByServiceId(oneService.Id),
		})
		println(oneService.Id)
	}
	return allServicesJson, nil
}
