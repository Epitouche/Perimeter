package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"area/repository"
	"area/schemas"
)

type TokenService interface {
	SaveToken(token schemas.Token) (tokenID uint64, err error)
	Update(token schemas.Token) error
	Delete(token schemas.Token) error
	FindAll() (allServices []schemas.Token)
	GetTokenById(id uint64) (schemas.Token, error)
	GetTokenByUserId(userID uint64) ([]schemas.Token, error)
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

func (service *tokenService) SaveToken(
	token schemas.Token,
) (tokenID uint64, err error) {
	tokens := service.repository.FindByToken(token.Token)
	for _, t := range tokens {
		if t.Token == token.Token {
			return t.Id, schemas.ErrTokenAlreadyExists
		}
	}

	service.repository.Save(token)
	tokens = service.repository.FindByToken(token.Token)

	for _, t := range tokens {
		if t.Token == token.Token {
			return t.Id, nil
		}
	}
	return 0, schemas.ErrUnableToSaveToken
}

func (service *tokenService) GetUserInfo(accessToken string) (schemas.GmailUserInfo, error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.GmailUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()
	return result, nil
}

func (service *tokenService) GetTokenById(id uint64) (schemas.Token, error) {
	return service.repository.FindById(id), nil
}

func (service *tokenService) GetTokenByUserId(userID uint64) ([]schemas.Token, error) {
	return service.repository.FindByUserId(userID), nil
}

func (service *tokenService) Update(token schemas.Token) error {
	service.repository.Update(token)
	return nil
}

func (service *tokenService) Delete(token schemas.Token) error {
	service.repository.Delete(token)
	return nil
}

func (service *tokenService) FindAll() []schemas.Token {
	return service.repository.FindAll()
}
