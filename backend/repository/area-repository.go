package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type AreaRepository interface {
	Save(action schemas.Area)
	Update(action schemas.Area)
	Delete(action schemas.Area)
	FindAll() []schemas.Area
	FindByUserId(userId uint64) []schemas.Area
}

type areaRepository struct {
	db *schemas.Database
}

func NewAreaRepository(conn *gorm.DB) AreaRepository {
	err := conn.AutoMigrate(&schemas.Area{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &areaRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *areaRepository) Save(action schemas.Area) {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *areaRepository) Update(action schemas.Area) {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *areaRepository) Delete(action schemas.Area) {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *areaRepository) FindAll() []schemas.Area {
	var action []schemas.Area
	err := repo.db.Connection.Preload("Service").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

func (repo *areaRepository) FindByUserId(userId uint64) []schemas.Area {
	var actions []schemas.Area
	err := repo.db.Connection.Where(&schemas.Area{UserId: userId}).
		Find(&actions)
	if err.Error != nil {
		panic(fmt.Errorf("failed to find action by service id: %v", err.Error))
	}
	return actions
}
