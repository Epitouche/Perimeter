package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type AreaController interface {
	CreateArea(ctx *gin.Context) (string, error)
	GetUserAreas(ctx *gin.Context) (areaList []schemas.Area, err error)
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
	println("CreateArea Controller")
	return controller.service.CreateArea(ctx)
}

func (controller *areaController) GetUserAreas(ctx *gin.Context) (areaList []schemas.Area, err error) {
	areaList, err = controller.service.GetUserAreas(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get user areas: %w", err)
	}
	return areaList, nil
}
