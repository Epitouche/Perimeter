package repository

import (
	"area/schemas"

	"gorm.io/gorm"
)

type ActionRepository interface {
	Save(action schemas.Action)
	Update(action schemas.Action)
	Delete(action schemas.Action)
	FindAll() []schemas.Action
	FindByName(actionName string) []schemas.Action
	FindByServiceByName(serviceId uint64, actionName string) []schemas.Action
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
	err := repo.db.Connection.Preload("Service").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

func (repo *actionRepository) FindByName(actionName string) []schemas.Action {
	var actions []schemas.Action
	err := repo.db.Connection.Where(&schemas.Action{Name: actionName}).Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}

func (repo *actionRepository) FindByServiceByName(
	serviceId uint64,
	actionName string,
) []schemas.Action {
	var actions []schemas.Action
	err := repo.db.Connection.Where(&schemas.Action{ServiceId: serviceId, Name: actionName}).
		Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}
