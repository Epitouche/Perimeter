package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type AreaController interface {
	CreateArea(ctx *gin.Context) (string, error)
	GetUserAreas(ctx *gin.Context) (areaList []schemas.Area, err error)
	UpdateUserArea(ctx *gin.Context) (newArea schemas.Area, err error)
	DeleteUserArea(ctx *gin.Context) (newArea schemas.Area, err error)
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
	var result schemas.AreaMessage

	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	fmt.Printf("result: %v\n", result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	return controller.service.CreateArea(result, token)
}

func (controller *areaController) GetUserAreas(
	ctx *gin.Context,
) (areaList []schemas.Area, err error) {
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	areaList, err = controller.service.GetUserAreas(token)
	if err != nil {
		return nil, fmt.Errorf("can't get user areas: %w", err)
	}
	return areaList, nil
}

func (controller *areaController) UpdateUserArea(
	ctx *gin.Context,
) (newArea schemas.Area, err error) {
	var result schemas.Area

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return newArea, fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	newArea, err = controller.service.UpdateUserArea(token, result)
	if err != nil {
		return newArea, fmt.Errorf("can't get user areas: %w", err)
	}
	return newArea, nil
}

func (controller *areaController) DeleteUserArea(
	ctx *gin.Context,
) (newArea schemas.Area, err error) {
	var result struct{ Id uint64 }

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return newArea, fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	newArea, err = controller.service.DeleteUserArea(token, result)
	if err != nil {
		return newArea, fmt.Errorf("can't get user areas: %w", err)
	}
	return newArea, nil
}
