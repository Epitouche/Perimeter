package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
)

type ServiceRepository interface {
	Save(service schemas.Service) error
	Update(service schemas.Service) error
	Delete(service schemas.Service) error
	FindAll() (services []schemas.Service, err error)
	FindAllByName(name schemas.ServiceName) (services []schemas.Service, err error)
	FindByName(name schemas.ServiceName) (service schemas.Service, err error)
	FindById(id uint64) (service schemas.Service, err error)
}

type serviceRepository struct {
	db *schemas.Database
}

func NewServiceRepository(conn *gorm.DB) ServiceRepository {
	err := conn.AutoMigrate(&schemas.Service{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &serviceRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *serviceRepository) Save(service schemas.Service) error {
	err := repo.db.Connection.Create(&service).Error
	if err != nil {
		return fmt.Errorf("failed to save service: %w", err)
	}
	return nil
}

func (repo *serviceRepository) Update(service schemas.Service) error {
	err := repo.db.Connection.Save(&service).Error
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	return nil
}

func (repo *serviceRepository) Delete(service schemas.Service) error {
	err := repo.db.Connection.Delete(&service).Error
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	return nil
}

func (repo *serviceRepository) FindAll() (services []schemas.Service, err error) {
	err = repo.db.Connection.Find(&services).Error
	if err != nil {
		return services, fmt.Errorf("failed to get all services: %w", err)
	}
	return services, nil
}

func (repo *serviceRepository) FindAllByName(
	name schemas.ServiceName,
) (services []schemas.Service, err error) {
	err = repo.db.Connection.Where(&schemas.Service{Name: name}).Find(&services).Error
	if err != nil {
		return services, fmt.Errorf("failed to get all services by name: %w", err)
	}
	return services, nil
}

func (repo *serviceRepository) FindByName(
	name schemas.ServiceName,
) (service schemas.Service, err error) {
	err = repo.db.Connection.Where(&schemas.Service{Name: name}).First(&service).Error
	if err != nil {
		return service, fmt.Errorf("failed to get service by name: %w", err)
	}
	return service, nil
}

func (repo *serviceRepository) FindById(id uint64) (service schemas.Service, err error) {
	err = repo.db.Connection.Where(&schemas.Service{Id: id}).First(&service).Error
	if err != nil {
		return service, fmt.Errorf("failed to get service by id: %w", err)
	}
	return service, nil
}
