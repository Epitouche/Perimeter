package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type AreaRepository interface {
	SaveArea(area schemas.Area) (areaID uint64, err error)
	Save(area schemas.Area) error
	Update(area schemas.Area) error
	Delete(area schemas.Area) error
	FindAll() (areas []schemas.Area, err error)
	FindByUserId(userID uint64) (areas []schemas.Area, err error)
	FindById(id uint64) (area schemas.Area, err error)
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
	err = repo.Save(action)
	if err != nil {
		return 0, fmt.Errorf("failed to save area: %w", err)
	}
	result := repo.db.Connection.Last(&action)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to save area: %w", err)
	}
	return action.Id, nil
}

func (repo *areaRepository) Save(action schemas.Area) error {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		return fmt.Errorf("failed to save area: %w", err.Error)
	}
	return nil
}

func (repo *areaRepository) Update(action schemas.Area) error {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		return fmt.Errorf("failed to update area: %w", err.Error)
	}
	return nil
}

func (repo *areaRepository) Delete(action schemas.Area) error {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		return fmt.Errorf("failed to delete area: %w", err.Error)
	}
	return nil
}

func (repo *areaRepository) FindAll() (areas []schemas.Area, err error) {
	err = repo.db.Connection.Preload("Service").Find(&areas).Error
	if err != nil {
		return areas, fmt.Errorf("failed to find all areas: %w", err)
	}
	return areas, nil
}

func (repo *areaRepository) FindByUserId(userID uint64) (areas []schemas.Area, err error) {
	err = repo.db.Connection.
		Preload("User").
		Preload("Action.Service").
		Preload("Reaction.Service").
		Where(&schemas.Area{UserId: userID}).
		Find(&areas).Error
	if err != nil {
		panic(fmt.Errorf("failed to find areas by user id: %w", err))
	}

	return areas, nil
}

func (repo *areaRepository) FindById(id uint64) (area schemas.Area, err error) {
	err = repo.db.Connection.Where(&schemas.Area{Id: id}).First(&area).Error
	var actionResult schemas.Action
	repo.db.Connection.Where(&schemas.Action{Id: area.ActionId}).First(&actionResult)
	area.Action = actionResult
	var reactionResult schemas.Reaction
	repo.db.Connection.Where(&schemas.Reaction{Id: area.ReactionId}).First(&reactionResult)
	area.Reaction = reactionResult

	if err != nil {
		return area, fmt.Errorf("failed to find action by id: %w", err)
	}
	return area, nil
}
