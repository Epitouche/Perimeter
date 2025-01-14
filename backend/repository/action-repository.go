package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type ActionRepository interface {
	Save(action schemas.Action) error
	Update(action schemas.Action) error
	Delete(action schemas.Action) error
	FindAll() (action []schemas.Action, err error)
	FindByName(actionName string) (action []schemas.Action, err error)
	FindByServiceId(serviceId uint64) (action []schemas.Action, err error)
	FindById(actionId uint64) (action schemas.Action, err error)
	FindByServiceByName(serviceId uint64, actionName string) (action []schemas.Action, err error)
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

func (repo *actionRepository) Save(action schemas.Action) error {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *actionRepository) Update(action schemas.Action) error {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *actionRepository) Delete(action schemas.Action) error {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *actionRepository) FindAll() (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

func (repo *actionRepository) FindByName(actionName string) (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{Name: actionName}).
		Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

func (repo *actionRepository) FindByServiceId(
	serviceId uint64,
) (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{ServiceId: serviceId}).
		Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

func (repo *actionRepository) FindByServiceByName(
	serviceId uint64,
	actionName string,
) (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{ServiceId: serviceId, Name: actionName}).
		Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

func (repo *actionRepository) FindById(actionId uint64) (action schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{Id: actionId}).
		First(&action)

	if errDatabase.Error != nil {
		return action, errDatabase.Error
	}
	return action, nil
}
