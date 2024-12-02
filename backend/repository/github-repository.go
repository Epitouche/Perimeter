package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

type GithubRepository interface {
	Save(token schemas.GithubToken)
	Update(token schemas.GithubToken)
	Delete(token schemas.GithubToken)
	FindAll() []schemas.GithubToken
	FindByAccessToken(accessToken string) []schemas.GithubToken
	FindById(id uint64) schemas.GithubToken
}

// Define a struct that embeds `*schemas.Database` and implements `GithubRepository`.
type githubRepository struct {
	db *schemas.Database
}

func NewGithubRepository(conn *gorm.DB) GithubRepository {
	err := conn.AutoMigrate(&schemas.GithubToken{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &githubRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *githubRepository) Save(token schemas.GithubToken) {
	err := repo.db.Connection.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *githubRepository) Update(token schemas.GithubToken) {
	err := repo.db.Connection.Save(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *githubRepository) Delete(token schemas.GithubToken) {
	err := repo.db.Connection.Delete(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *githubRepository) FindAll() []schemas.GithubToken {
	var tokens []schemas.GithubToken
	err := repo.db.Connection.Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *githubRepository) FindByAccessToken(accessToken string) []schemas.GithubToken {
	var tokens []schemas.GithubToken
	err := repo.db.Connection.Where(&schemas.GithubToken{AccessToken: accessToken}).Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *githubRepository) FindById(id uint64) schemas.GithubToken {
	var tokens schemas.GithubToken
	err := repo.db.Connection.Where(&schemas.GithubToken{Id: id}).First(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}
