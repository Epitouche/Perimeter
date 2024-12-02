package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type TokenRepository interface {
	Save(service schemas.Token)
	Update(service schemas.Token)
	Delete(service schemas.Token)
	FindAll() []schemas.Token
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

func (repo *tokenRepository) Save(service schemas.Token) {
	err := repo.db.Connection.Create(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) Update(service schemas.Token) {
	err := repo.db.Connection.Save(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) Delete(service schemas.Token) {
	err := repo.db.Connection.Delete(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) FindAll() []schemas.Token {
	var service []schemas.Token
	err := repo.db.Connection.Find(&service)
	if err.Error != nil {
		panic(err.Error)
	}
	return service
}
