package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Epitouche/Perimeter/schemas"
)

type UserRepository interface {
	Save(user schemas.User) (err error)
	Update(user schemas.User) (err error)
	Delete(user schemas.User) (err error)
	FindAll() (users []schemas.User, err error)
	FindByEmail(email string) (users []schemas.User, err error)
	FindByUserName(username string) (users []schemas.User, err error)
	FindById(id uint64) (user schemas.User, err error)
}

// Define a struct that embeds `*schemas.Database` and implements `UserRepository`.
type userRepository struct {
	db *schemas.Database
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	err := conn.AutoMigrate(&schemas.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &userRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *userRepository) Save(user schemas.User) (err error) {
	err = repo.db.Connection.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (repo *userRepository) Update(user schemas.User) (err error) {
	err = repo.db.Connection.Save(&user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (repo *userRepository) Delete(user schemas.User) (err error) {
	err = repo.db.Connection.Delete(&user).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (repo *userRepository) FindAll() (users []schemas.User, err error) {
	err = repo.db.Connection.Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("failed to find all users: %w", err)
	}
	return users, nil
}

func (repo *userRepository) FindByEmail(email string) (users []schemas.User, err error) {
	err = repo.db.Connection.Where(&schemas.User{Email: email}).Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("failed to find user by email: %w", err)
	}
	return users, nil
}

func (repo *userRepository) FindByUserName(username string) (users []schemas.User, err error) {
	err = repo.db.Connection.Where(&schemas.User{Username: username}).Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("failed to find user by username: %w", err)
	}
	return users, nil
}

func (repo *userRepository) FindById(id uint64) (user schemas.User, err error) {
	err = repo.db.Connection.Where(&schemas.User{Id: id}).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("failed to find user by id: %w", err)
	}
	return user, nil
}
