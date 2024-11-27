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
}

func NewServiceService(repository repository.ServiceRepository) ServiceService {
	InitialSaveService(repository)
	return &serviceService{
		repository: repository,
	}
}

func InitialSaveService(repository repository.ServiceRepository) {
	repository.Save(schemas.Service{
		Name:        "Github",
		Description: "Github API",
	})

	println("Service Github created\n")
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
