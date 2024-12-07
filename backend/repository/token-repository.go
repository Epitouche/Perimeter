package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type TokenRepository interface {
	Save(token schemas.Token)
	Update(token schemas.Token)
	Delete(token schemas.Token)
	FindAll() []schemas.Token
	FindByToken(token string) []schemas.Token
	FindById(id uint64) schemas.Token
	FindByUserIdAndServiceId(id uint64, serviceId uint64) schemas.Token
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

func (repo *tokenRepository) Save(token schemas.Token) {
	err := repo.db.Connection.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) Update(token schemas.Token) {
	err := repo.db.Connection.Save(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) Delete(token schemas.Token) {
	err := repo.db.Connection.Delete(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) FindAll() []schemas.Token {
	var token []schemas.Token
	err := repo.db.Connection.Find(&token)
	if err.Error != nil {
		panic(err.Error)
	}
	return token
}

func (repo *tokenRepository) FindByToken(token string) []schemas.Token {
	var tokens []schemas.Token
	err := repo.db.Connection.Where(&schemas.Token{Token: token}).Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *tokenRepository) FindById(id uint64) schemas.Token {
	var tokens schemas.Token
	err := repo.db.Connection.Where(&schemas.Token{Id: id}).First(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *tokenRepository) FindByUserIdAndServiceId(id uint64, serviceId uint64) schemas.Token {
	var tokens schemas.Token
	err := repo.db.Connection.Where(&schemas.Token{UserId: id, ServiceId: serviceId}).First(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}
