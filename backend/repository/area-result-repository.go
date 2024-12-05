package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type AreaResultRepository interface {
	Save(action schemas.AreaResult)
	Update(action schemas.AreaResult)
	Delete(action schemas.AreaResult)
	FindAll() []schemas.AreaResult
	FindByAreaId(userId uint64) []schemas.AreaResult
}

type areaResultRepository struct {
	db *schemas.Database
}

func NewAreaResultRepository(conn *gorm.DB) AreaResultRepository {
	err := conn.AutoMigrate(&schemas.AreaResult{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &areaResultRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *areaResultRepository) Save(action schemas.AreaResult) {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *areaResultRepository) Update(action schemas.AreaResult) {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *areaResultRepository) Delete(action schemas.AreaResult) {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *areaResultRepository) FindAll() []schemas.AreaResult {
	var action []schemas.AreaResult
	err := repo.db.Connection.Preload("Service").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

func (repo *areaResultRepository) FindByAreaId(areaId uint64) []schemas.AreaResult {
	var actions []schemas.AreaResult
	err := repo.db.Connection.Where(&schemas.AreaResult{AreaId: areaId}).
		Find(&actions)
	if err.Error != nil {
		panic(fmt.Errorf("failed to find action by service id: %v", err.Error))
	}
	return actions
}
