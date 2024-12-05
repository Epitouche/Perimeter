package service

import (
	"area/repository"
	"area/schemas"
)

type ServiceService interface {
	FindAll() (allServices []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	GetAllServices() (allServicesJson []schemas.ServiceJson, err error)
	GetServices() []interface{}
	FindActionbyName(name string) func(c chan string, option string)
}

type serviceService struct {
	repository repository.ServiceRepository
	allService []interface{}
	allServiceSchemas []schemas.Service
}

func NewServiceService(repository repository.ServiceRepository, timerService TimerService) ServiceService {
	newService := serviceService{
		repository: repository,
		allServiceSchemas: []schemas.Service{
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
			{
				Name:        schemas.Gmail,
				Description: "This service is a mail service",
			},
		},
		allService: []interface{}{ timerService },
	}
	newService.InitialSaveService()
	return &newService
}

func (service *serviceService) InitialSaveService() {
	for _, oneService := range service.allServiceSchemas {
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

func (service *serviceService) GetServices() []interface{} {
	return service.allService
}

func (service *serviceService) FindActionbyName(name string) func(c chan string, option string) {
	for _, service := range service.allService {
		if timerService, ok := service.(TimerService); ok {
			return timerService.FindActionbyName(name)
		}
	}
	return nil
}