package service

import (
	"area/repository"
	"area/schemas"
)

type TokenService interface {
	FindAll() (allServices []schemas.Token)
}

type tokenService struct {
	repository repository.TokenRepository
}

func NewTokenService(repository repository.TokenRepository) TokenService {
	newService := tokenService{
		repository: repository,
	}
	return &newService
}

func (service *tokenService) FindAll() (allServices []schemas.Token) {
	return service.repository.FindAll()
}
