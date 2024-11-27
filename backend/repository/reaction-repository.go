package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type ReactionRepository interface {
	Save(action schemas.Reaction)
	Update(action schemas.Reaction)
	Delete(action schemas.Reaction)
	FindAll() []schemas.Reaction
	FindByName(actionName string) []schemas.Reaction
	FindByServiceByName(serviceId uint64, actionName string) []schemas.Reaction
}

type reactionRepository struct {
	db *schemas.Database
}

func NewReactionRepository(conn *gorm.DB) ReactionRepository {
	err := conn.AutoMigrate(&schemas.Reaction{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &reactionRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *reactionRepository) Save(action schemas.Reaction) {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Update(action schemas.Reaction) {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Delete(action schemas.Reaction) {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) FindAll() []schemas.Reaction {
	var action []schemas.Reaction
	err := repo.db.Connection.Preload("Service").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

func (repo *reactionRepository) FindByName(actionName string) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{Name: actionName}).Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}

func (repo *reactionRepository) FindByServiceByName(serviceId uint64, actionName string) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{ServiceId: serviceId, Name: actionName}).Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}
