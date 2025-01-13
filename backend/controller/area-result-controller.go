package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type AreaResultController interface {
	GetUserAreaResultsByAreaID(ctx *gin.Context, areaID uint64) (areaList []schemas.AreaResult, err error)
}

type areaResultController struct {
	service     service.AreaResultService
	serviceArea service.AreaService
}

func NewAreaResultController(service service.AreaResultService, serviceArea service.AreaService) AreaResultController {
	return &areaResultController{
		service:     service,
		serviceArea: serviceArea,
	}
}

func (controller *areaResultController) GetUserAreaResultsByAreaID(
	ctx *gin.Context,
	areaID uint64,
) (areaResultList []schemas.AreaResult, err error) {
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]

	areaList, err := controller.serviceArea.GetUserAreas(token)
	if err != nil {
		return nil, fmt.Errorf("can't get user areas: %w", err)
	}

	for _, area := range areaList {
		if area.Id == areaID {
			areaResultList = controller.service.FindByAreaID(areaID)
			return areaResultList, nil
		}
	}
	return areaResultList, fmt.Errorf("area not found")
}
