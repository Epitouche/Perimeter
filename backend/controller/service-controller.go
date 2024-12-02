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
	service         service.ServiceService
	serviceAction   service.ActionService
	serviceReaction service.ReactionService
}

func NewServiceController(
	service service.ServiceService,
	serviceAction service.ActionService,
	serviceReaction service.ReactionService,
) ServiceController {
	return &serviceController{
		service:         service,
		serviceAction:   serviceAction,
		serviceReaction: serviceReaction,
	}
}

func (controller *serviceController) AboutJson(
	ctx *gin.Context,
) (allServicesJson []schemas.ServiceJson, err error) {
	allServices := controller.service.FindAll()
	for _, oneService := range allServices {
		allServicesJson = append(allServicesJson, schemas.ServiceJson{
			Name:     schemas.ServiceName(oneService.Name),
			Action:   controller.serviceAction.GetAllServicesByServiceId(oneService.Id),
			Reaction: controller.serviceReaction.GetAllServicesByServiceId(oneService.Id),
		})
	}
	return allServicesJson, nil
}
