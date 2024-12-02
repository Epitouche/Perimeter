package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type GmailRepository interface {
	Save(token schemas.GmailToken)
	Update(token schemas.GmailToken)
	Delete(token schemas.GmailToken)
	FindAll() []schemas.GmailToken
	FindByAccessToken(accessToken string) []schemas.GmailToken
	FindById(id uint64) schemas.GmailToken
}

// Define a struct that embeds `*schemas.Database` and implements `GmailRepository`.
type gmailRepository struct {
	db *schemas.Database
}

func NewGmailRepository(conn *gorm.DB) GmailRepository {
	err := conn.AutoMigrate(&schemas.GmailToken{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &gmailRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *gmailRepository) Save(token schemas.GmailToken) {
	err := repo.db.Connection.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *gmailRepository) Update(token schemas.GmailToken) {
	err := repo.db.Connection.Save(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *gmailRepository) Delete(token schemas.GmailToken) {
	err := repo.db.Connection.Delete(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *gmailRepository) FindAll() []schemas.GmailToken {
	var tokens []schemas.GmailToken
	err := repo.db.Connection.Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *gmailRepository) FindByAccessToken(accessToken string) []schemas.GmailToken {
	var tokens []schemas.GmailToken
	err := repo.db.Connection.Where(&schemas.GmailToken{AccessToken: accessToken}).Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *gmailRepository) FindById(id uint64) schemas.GmailToken {
	var tokens schemas.GmailToken
	err := repo.db.Connection.Where(&schemas.GmailToken{Id: id}).First(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}
