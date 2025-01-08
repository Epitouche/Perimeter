package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/service"
)

type ServiceController interface {
	AboutJSON(ctx *gin.Context) (aboutJSON schemas.AboutJSON, err error)
	GetServicesInfo() (response []schemas.Service, err error)
	GetServiceInfoById(id uint64) (response schemas.Service, err error)
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
	aboutJSON.Server.CurrentTime = strconv.FormatInt(time.Now().Unix(), 10)
	aboutJSON.Server.Services = allServicesJSON
	return aboutJSON, nil
}

func (controller *serviceController) GetServicesInfo() (response []schemas.Service, err error) {
	response, err = controller.service.GetServicesInfo()
	if err != nil {
		return nil, fmt.Errorf("can't get services info: %w", err)
	}
	return response, nil
}

func (controller *serviceController) GetServiceInfoById(
	id uint64,
) (response schemas.Service, err error) {
	response = controller.service.GetServiceById(id)
	return response, nil
}
