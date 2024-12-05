package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/repository"
	"area/schemas"
)

type AreaService interface {
	FindAll() []schemas.Area
	CreateArea(ctx *gin.Context) (string, error)
}

type areaService struct {
	repository repository.AreaRepository
}

func NewAreaService(
	repository repository.AreaRepository,
	serviceService ServiceService,
) AreaService {
	newService := areaService{
		repository: repository,
	}
	return &newService
}

func (service *areaService) FindAll() []schemas.Area {
	return service.repository.FindAll()
}

func (service *areaService) CreateArea(ctx *gin.Context) (string, error) {
	var result schemas.Area
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	service.repository.Save(result)
	return "Area created successfully", nil
}
