package controller

import (
	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type AreaController interface {
	CreateArea(ctx *gin.Context) (string, error)
	GetUserAreas(ctx *gin.Context) ([]schemas.Area, error)
}

type areaController struct {
	service service.AreaService
}

func NewAreaController(service service.AreaService) AreaController {
	return &areaController{
		service: service,
	}
}

func (controller *areaController) CreateArea(ctx *gin.Context) (string, error) {
	return controller.service.CreateArea(ctx)
}

func (controller *areaController) GetUserAreas(ctx *gin.Context) ([]schemas.Area, error) {
	return controller.service.GetUserAreas(ctx)
}
