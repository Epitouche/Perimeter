package service

import (
	"area/repository"
	"area/schemas"
)

type ServiceService interface {
	Save(newService schemas.Service) error
	Update(newService schemas.Service) error
	Delete(newService schemas.Service) error
	FindAll() []schemas.Service
}

type serviceService struct {
	repository repository.ServiceRepository
	allService []schemas.Service
}

func NewServiceService(repository repository.ServiceRepository) ServiceService {
	newService := serviceService{
		repository: repository,
		allService: []schemas.Service{
			{
				Name:        string(schemas.Spotify),
				Description: "This service is a music service",
			},
			{
				Name:        string(schemas.OpenWeatherMap),
				Description: "This service is a weather service",
			},
			{
				Name:        string(schemas.Timer),
				Description: "This service is a time service",
			},
		},
	}
	newService.InitialSaveService()
	return &newService
}

func (service *serviceService) InitialSaveService() {
	for _, oneService := range service.allService {
		serviceByName := service.repository.FindByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
		}
	}
}

func (service *serviceService) Save(newService schemas.Service) error {
	service.repository.Save(newService)
	return nil
}

func (service *serviceService) Update(newService schemas.Service) error {
	service.repository.Update(newService)
	return nil
}

func (service *serviceService) Delete(newService schemas.Service) error {
	service.repository.Delete(newService)
	return nil
}

func (service *serviceService) FindAll() []schemas.Service {
	return service.repository.FindAll()
}
