package service

import (
	"area/repository"
	"area/schemas"
)

type ServiceService interface {
	Save(schemas.Link) error
	Update(schemas.Link) error
	Delete(schemas.Link) error
	FindAll() []schemas.Link
}

type serviceService struct {
	repository repository.LinkRepository
}

func NewServiceService(videoRepository repository.LinkRepository) ServiceService {
	return &serviceService{
		repository: videoRepository,
	}
}

func (service *serviceService) Save(link schemas.Link) error {
	service.repository.Save(link)
	return nil
}

func (service *serviceService) Update(link schemas.Link) error {
	service.repository.Update(link)
	return nil
}

func (service *serviceService) Delete(link schemas.Link) error {
	service.repository.Delete(link)
	return nil
}

func (service *serviceService) FindAll() []schemas.Link {
	return service.repository.FindAll()
}
