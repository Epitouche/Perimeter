package service

import (
	"area/repository"
	"area/schemas"
)

type ServiceService interface {
	FindAll() (allServices []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	GetAllServices() (allServicesJson []schemas.ServiceJson, err error)
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
				Name:        schemas.Spotify,
				Description: "This service is a music service",
			},
			{
				Name:        schemas.OpenWeatherMap,
				Description: "This service is a weather service",
			},
			{
				Name:        schemas.Timer,
				Description: "This service is a time service",
			},
		},
	}
	newService.InitialSaveService()
	return &newService
}

func (service *serviceService) InitialSaveService() {
	for _, oneService := range service.allService {
		serviceByName := service.repository.FindAllByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
		}
	}
}

func (service *serviceService) FindAll() (allServices []schemas.Service) {
	return service.repository.FindAll()
}

func (service *serviceService) GetAllServices() (allServicesJson []schemas.ServiceJson, err error) {
	allServices := service.repository.FindAll()
	for _, oneService := range allServices {
		println(oneService.Name)
		allServicesJson = append(allServicesJson, schemas.ServiceJson{
			Name: schemas.ServiceName(oneService.Name),
		})
	}
	return allServicesJson, nil
}

func (service *serviceService) FindByName(serviceName schemas.ServiceName) schemas.Service {
	return service.repository.FindByName(serviceName)
}
