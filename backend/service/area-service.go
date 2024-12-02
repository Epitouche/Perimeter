package service

import (
	"area/repository"
	"area/schemas"
)

type AreaService interface {
	FindAll() []schemas.Area
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
