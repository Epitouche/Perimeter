package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type AreaRepository interface {
	SaveArea(action schemas.Area) (areaID uint64, err error)
	Save(action schemas.Area)
	Update(action schemas.Area)
	Delete(action schemas.Area)
	FindAll() []schemas.Area
	FindByUserId(userID uint64) []schemas.Area
	FindById(id uint64) (schemas.Area, error)
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

func (repo *areaRepository) SaveArea(action schemas.Area) (areaID uint64, err error) {
	repo.Save(action)
	result := repo.db.Connection.Last(&action)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to save area: %w", err)
	}
	return action.Id, nil
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

func (repo *areaRepository) FindByUserId(userID uint64) []schemas.Area {
	var areas []schemas.Area

	err := repo.db.Connection.
		Preload("User").
		Preload("Action.Service").
		Preload("Reaction.Service").
		Where(&schemas.Area{UserId: userID}).
		Find(&areas)

	if err.Error != nil {
		panic(fmt.Errorf("failed to find areas by user id: %w", err.Error))
	}

	return areas
}

func (repo *areaRepository) FindById(id uint64) (schemas.Area, error) {
	var area schemas.Area
	err := repo.db.Connection.Where(&schemas.Area{Id: id}).First(&area)
	var actionResult schemas.Action
	repo.db.Connection.Where(&schemas.Action{Id: area.ActionId}).First(&actionResult)
	area.Action = actionResult
	var reactionResult schemas.Reaction
	repo.db.Connection.Where(&schemas.Reaction{Id: area.ReactionId}).First(&reactionResult)
	area.Reaction = reactionResult

	if err.Error != nil {
		println(err.Error)
		return schemas.Area{}, fmt.Errorf("failed to find action by id: %w", err.Error)
	}
	return area, nil
}
