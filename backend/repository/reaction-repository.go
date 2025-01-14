package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type ReactionRepository interface {
	Save(reaction schemas.Reaction) error
	Update(reaction schemas.Reaction) error
	Delete(reaction schemas.Reaction) error
	FindAll() (reactions []schemas.Reaction, err error)
	FindByName(actionName string) (reactions []schemas.Reaction, err error)
	FindByServiceId(serviceId uint64) (reactions []schemas.Reaction, err error)
	FindByServiceByName(
		serviceID uint64,
		actionName string,
	) (reactions []schemas.Reaction, err error)
	FindById(actionId uint64) (reaction schemas.Reaction, err error)
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

func (repo *reactionRepository) Save(reaction schemas.Reaction) error {
	err := repo.db.Connection.Create(&reaction)
	if err.Error != nil {
		return fmt.Errorf("failed to save reaction: %w", err.Error)
	}
	return nil
}

func (repo *reactionRepository) Update(reaction schemas.Reaction) error {
	err := repo.db.Connection.Save(&reaction)
	if err.Error != nil {
		return fmt.Errorf("failed to update reaction: %w", err.Error)
	}
	return nil
}

func (repo *reactionRepository) Delete(reaction schemas.Reaction) error {
	err := repo.db.Connection.Delete(&reaction).Error
	if err != nil {
		return fmt.Errorf("failed to delete reaction: %w", err)
	}
	return nil
}

func (repo *reactionRepository) FindAll() (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").Find(&reactions).Error
	if err != nil {
		return reactions, fmt.Errorf("failed to find all reactions: %w", err)
	}
	return reactions, nil
}

func (repo *reactionRepository) FindByName(
	actionName string,
) (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").
		Where(&schemas.Reaction{Name: actionName}).
		Find(&reactions).
		Error
	if err != nil {
		return reactions, fmt.Errorf("failed to find reaction by name: %w", err)
	}
	return reactions, nil
}

func (repo *reactionRepository) FindByServiceId(
	serviceId uint64,
) (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").Where(&schemas.Reaction{ServiceId: serviceId}).
		Find(&reactions).Error
	if err != nil {
		panic(fmt.Errorf("failed to find reaction by service id: %w", err))
	}
	return reactions, nil
}

func (repo *reactionRepository) FindByServiceByName(
	serviceID uint64,
	actionName string,
) (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").
		Where(&schemas.Reaction{ServiceId: serviceID, Name: actionName}).
		Find(&reactions).
		Error
	if err != nil {
		panic(fmt.Errorf("failed to find reaction by service name: %w", err))
	}
	return reactions, nil
}

func (repo *reactionRepository) FindById(actionId uint64) (reaction schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").
		Where(&schemas.Reaction{Id: actionId}).
		First(&reaction).
		Error
	if err != nil {
		return reaction, fmt.Errorf("failed to find reaction by id: %w", err)
	}
	return reaction, nil
}
