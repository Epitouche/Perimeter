package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

type TokenRepository interface {
	Save(token schemas.Token) error
	Update(token schemas.Token) error
	Delete(token schemas.Token) error
	FindAll() (tokens []schemas.Token, err error)
	FindByToken(token string) (tokens []schemas.Token, err error)
	FindById(id uint64) (token schemas.Token, err error)
	FindByUserId(userID uint64) (tokens []schemas.Token, err error)
	FindByUserIdAndServiceId(id uint64, serviceId uint64) (token schemas.Token, err error)
}

type tokenRepository struct {
	db *schemas.Database
}

func NewTokenRepository(conn *gorm.DB) TokenRepository {
	err := conn.AutoMigrate(&schemas.Token{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &tokenRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *tokenRepository) Save(token schemas.Token) error {
	err := repo.db.Connection.Create(&token).Error
	if err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}
	return nil
}

func (repo *tokenRepository) Update(token schemas.Token) error {
	err := repo.db.Connection.Save(&token).Error
	if err != nil {
		return fmt.Errorf("failed to update token: %w", err)
	}
	return nil
}

func (repo *tokenRepository) Delete(token schemas.Token) error {
	err := repo.db.Connection.Delete(&token).Error
	if err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}
	return nil
}

func (repo *tokenRepository) FindAll() (tokens []schemas.Token, err error) {
	err = repo.db.Connection.Find(&tokens).Error
	if err != nil {
		return tokens, fmt.Errorf("failed to find all tokens: %w", err)
	}
	return tokens, nil
}

func (repo *tokenRepository) FindByToken(token string) (tokens []schemas.Token, err error) {
	err = repo.db.Connection.Where(&schemas.Token{Token: token}).Find(&tokens).Error
	if err != nil {
		return tokens, fmt.Errorf("failed to find token by token: %w", err)
	}
	return tokens, nil
}

func (repo *tokenRepository) FindById(id uint64) (token schemas.Token, err error) {
	err = repo.db.Connection.Where(&schemas.Token{Id: id}).First(&token).Error
	if err != nil {
		return token, fmt.Errorf("failed to find token by id: %w", err)
	}
	return token, nil
}

func (repo *tokenRepository) FindByUserId(userID uint64) (tokens []schemas.Token, err error) {
	err = repo.db.Connection.
		Preload("User").
		Preload("Service").
		Where(&schemas.Token{UserId: userID}).Find(&tokens).Error
	if err != nil {
		return tokens, fmt.Errorf("failed to find token by user id: %w", err)
	}
	return tokens, nil
}

func (repo *tokenRepository) FindByUserIdAndServiceId(id uint64, serviceId uint64) (token schemas.Token, err error) {
	err = repo.db.Connection.Where(&schemas.Token{UserId: id, ServiceId: serviceId}).First(&token).Error
	if err != nil {
		return token, fmt.Errorf("failed to find token by user id and service id: %w", err)
	}
	return token, nil
}
