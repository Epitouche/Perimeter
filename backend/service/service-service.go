package service

import (
	"area/repository"
	"area/schemas"
)

type ServiceService interface {
	FindAll() (allServices []schemas.Service)
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

func (service *serviceService) FindAll() (allServices []schemas.Service) {
	return service.repository.FindAll()
}

func (service *serviceService) GetAllServices() (allServicesJson []schemas.ServiceJson, err error) {
	allServicesJson = []schemas.ServiceJson{}
	allServices := service.repository.FindAll()
	for _, oneService := range allServices {
		println(oneService.Name)
		allServicesJson = append(allServicesJson, schemas.ServiceJson{
			Name: oneService.Name,
		})
	}
	return allServicesJson, nil
}
