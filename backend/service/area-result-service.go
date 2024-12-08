package service

import (
	"area/repository"
	"area/schemas"
)

type AreaResultService interface {
	FindAll() []schemas.AreaResult
}

type areaResultService struct {
	repository repository.AreaResultRepository
}

func NewAreaResultService(
	repository repository.AreaResultRepository,
) AreaResultService {
	newService := areaResultService{
		repository: repository,
	}
	return &newService
}

func (service *areaResultService) FindAll() []schemas.AreaResult {
	return service.repository.FindAll()
}
