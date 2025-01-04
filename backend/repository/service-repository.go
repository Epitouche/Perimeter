package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type ServiceRepository interface {
	Save(service schemas.Service)
	Update(service schemas.Service)
	Delete(service schemas.Service)
	FindAll() []schemas.Service
	FindAllByName(name schemas.ServiceName) []schemas.Service
	FindByName(name schemas.ServiceName) schemas.Service
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

func (repo *serviceRepository) Save(service schemas.Service) {
	err := repo.db.Connection.Create(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) Update(service schemas.Service) {
	err := repo.db.Connection.Save(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) Delete(service schemas.Service) {
	err := repo.db.Connection.Delete(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) FindAll() []schemas.Service {
	var service []schemas.Service
	err := repo.db.Connection.Find(&service)
	if err.Error != nil {
		panic(err.Error)
	}
	return service
}

func (repo *serviceRepository) FindAllByName(name schemas.ServiceName) []schemas.Service {
	var services []schemas.Service
	err := repo.db.Connection.Where(&schemas.Service{Name: name}).Find(&services)
	if err.Error != nil {
		panic(err.Error)
	}
	return services
}

func (repo *serviceRepository) FindByName(name schemas.ServiceName) schemas.Service {
	var services schemas.Service
	err := repo.db.Connection.Where(&schemas.Service{Name: name}).First(&services)
	if err.Error != nil {
		panic(err.Error)
	}
	return services
}
