package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type ReactionRepository interface {
	Save(reaction schemas.Reaction)
	Update(reaction schemas.Reaction)
	Delete(reaction schemas.Reaction)
	FindAll() []schemas.Reaction
	FindByName(actionName string) []schemas.Reaction
	FindByServiceId(serviceId uint64) []schemas.Reaction
	FindByServiceByName(serviceID uint64, actionName string) []schemas.Reaction
	FindById(actionId uint64) schemas.Reaction
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

func (repo *reactionRepository) Save(reaction schemas.Reaction) {
	err := repo.db.Connection.Create(&reaction)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Update(reaction schemas.Reaction) {
	err := repo.db.Connection.Save(&reaction)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Delete(reaction schemas.Reaction) {
	err := repo.db.Connection.Delete(&reaction)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) FindAll() []schemas.Reaction {
	var reaction []schemas.Reaction
	err := repo.db.Connection.Preload("Service").Find(&reaction)
	if err.Error != nil {
		panic(err.Error)
	}
	return reaction
}

func (repo *reactionRepository) FindByName(actionName string) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{Name: actionName}).Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}

func (repo *reactionRepository) FindByServiceId(serviceId uint64) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{ServiceId: serviceId}).
		Find(&actions)
	if err.Error != nil {
		panic(fmt.Errorf("failed to find reaction by service id: %w", err.Error))
	}
	return actions
}

func (repo *reactionRepository) FindByServiceByName(
	serviceID uint64,
	actionName string,
) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{ServiceId: serviceID, Name: actionName}).
		Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}

func (repo *reactionRepository) FindById(actionId uint64) schemas.Reaction {
	var reaction schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{Id: actionId}).First(&reaction)
	if err.Error != nil {
		panic(err.Error)
	}
	return reaction
}
