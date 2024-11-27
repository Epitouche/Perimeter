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
	err := repo.db.Connection.Preload("UrlId").Find(&service)
	if err.Error != nil {
		panic(err.Error)
	}
	return service
}
