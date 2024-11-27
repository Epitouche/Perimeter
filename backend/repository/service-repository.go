package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type ServiceRepository interface {
	Save(link schemas.Link)
	Update(link schemas.Link)
	Delete(link schemas.Link)
	FindAll() []schemas.Link
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

func (repo *serviceRepository) Save(video schemas.Link) {
	err := repo.db.Connection.Create(&video)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) Update(video schemas.Link) {
	err := repo.db.Connection.Save(&video)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) Delete(video schemas.Link) {
	err := repo.db.Connection.Delete(&video)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) FindAll() []schemas.Link {
	var links []schemas.Link
	err := repo.db.Connection.Preload("UrlId").Find(&links)
	if err.Error != nil {
		panic(err.Error)
	}
	return links
}
