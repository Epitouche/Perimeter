package service

import (
	"area/repository"
	"area/schemas"
)

allService := []schemas.Service{
	{
		Name:        "Spotify",
		Description: "This service is a music service",
	},
	{
		Name:        "OpenWeatherMap",
		Description: "This service is a weather service",
	},
	{
		Name:        "Time",
		Description: "This service is a time service",
	}
}

type ServiceService interface {
	Save(newService schemas.Service) error
	Update(newService schemas.Service) error
	Delete(newService schemas.Service) error
	FindAll() []schemas.Service
}

type serviceService struct {
	repository repository.ServiceRepository
}

func NewServiceService(repository repository.ServiceRepository) ServiceService {
	InitialSaveService(repository)
	return &serviceService{
		repository: repository,
	}
}

func InitialSaveService(repository repository.ServiceRepository) {
	for _, service := range allService {
		repository.Save(service)
	}

	println("All Service created\n")
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
