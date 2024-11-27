package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type ActionRepository interface {
	Save(action schemas.Action)
	Update(action schemas.Action)
	Delete(action schemas.Action)
	FindAll() []schemas.Action
	FindByName(name string) []schemas.Action
}

type actionRepository struct {
	db *schemas.Database
}

func NewActionRepository(conn *gorm.DB) ActionRepository {
	err := conn.AutoMigrate(&schemas.Action{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &actionRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *actionRepository) Save(action schemas.Action) {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *actionRepository) Update(action schemas.Action) {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *actionRepository) Delete(action schemas.Action) {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *actionRepository) FindAll() []schemas.Action {
	var action []schemas.Action
	err := repo.db.Connection.Preload("UrlId").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

func (repo *actionRepository) FindByName(name string) []schemas.Action {
	var actions []schemas.Action
	err := repo.db.Connection.Where(&schemas.Action{Name: name}).Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}
