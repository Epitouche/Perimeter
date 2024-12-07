package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type ServiceController interface {
	AboutJSON(ctx *gin.Context) (aboutJSON schemas.AboutJSON, err error)
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
) (aboutJSON schemas.AboutJSON, err error) {
	allServicesJSON := []schemas.ServiceJSON{}
	allServices := controller.service.FindAll()
	for _, oneService := range allServices {
		allServicesJSON = append(allServicesJSON, schemas.ServiceJSON{
			Name:     schemas.ServiceName(oneService.Name),
			Action:   controller.serviceAction.GetAllServicesByServiceId(oneService.Id),
			Reaction: controller.serviceReaction.GetAllServicesByServiceId(oneService.Id),
		})
	}
	aboutJSON.Client.Host = ctx.ClientIP()
	aboutJSON.Server.CurrentTime = fmt.Sprintf("%d", time.Now().Unix())
	aboutJSON.Server.Services = allServicesJSON
	return aboutJSON, nil
}

func (controller *serviceController) GetServicesInfo() (response []schemas.Service, err error) {
	return controller.service.GetServicesInfo()
}
