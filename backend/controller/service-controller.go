package controller

import (
	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type ServiceController interface {
	AboutJSON(ctx *gin.Context) (allService []schemas.ServiceJSON, err error)
	GetServicesInfo() (response []schemas.Service, err error)
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

func (controller *serviceController) AboutJSON(
	ctx *gin.Context,
) (allServicesJSON []schemas.ServiceJSON, err error) {
	allServices := controller.service.FindAll()
	for _, oneService := range allServices {
		allServicesJSON = append(allServicesJSON, schemas.ServiceJSON{
			Name:     schemas.ServiceName(oneService.Name),
			Action:   controller.serviceAction.GetAllServicesByServiceId(oneService.Id),
			Reaction: controller.serviceReaction.GetAllServicesByServiceId(oneService.Id),
		})
	}
	return allServicesJSON, nil
}

func (controller *serviceController) GetServicesInfo() (response []schemas.Service, err error) {
	return controller.service.GetServicesInfo()
}
