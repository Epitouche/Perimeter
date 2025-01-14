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
	DeleteUserToken(
		token string,
		tokenToDelete struct{ Id uint64 },
	) (deletedToken schemas.Token, err error)
}

type tokenService struct {
	repository  repository.TokenRepository
	serviceUser UserService
}

func NewTokenService(repository repository.TokenRepository, serviceUser UserService) TokenService {
	newService := tokenService{
		repository:  repository,
		serviceUser: serviceUser,
	}
	return &newService
}

func (service *tokenService) SaveToken(
	token schemas.Token,
) (tokenID uint64, err error) {
	tokens, err := service.repository.FindByToken(token.Token)
	if err != nil {
		return 0, err
	}
	for _, t := range tokens {
		if t.Token == token.Token {
			return t.Id, schemas.ErrTokenAlreadyExists
		}
	}

	err = service.repository.Save(token)
	if err != nil {
		return 0, err
	}
	tokens, err = service.repository.FindByToken(token.Token)
	if err != nil {
		return 0, err
	}

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

	defer resp.Body.Close()

	result := schemas.GmailUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to decode response because %w", err)
	}

	return result, nil
}

func (service *tokenService) GetTokenById(id uint64) (schemas.Token, error) {
	token, err := service.repository.FindById(id)
	if err != nil {
		return schemas.Token{}, err
	}
	return token, nil
}

func (service *tokenService) GetTokenByUserId(userID uint64) ([]schemas.Token, error) {
	tokens, err := service.repository.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (service *tokenService) Update(token schemas.Token) error {
	return service.repository.Update(token)
}

func (service *tokenService) Delete(token schemas.Token) error {
	return service.repository.Delete(token)
}

func (service *tokenService) FindAll() []schemas.Token {
	tokens, err := service.repository.FindAll()
	if err != nil {
		return nil
	}
	return tokens
}

// containsArea checks if a slice of areas contains a specific area
func containsToken(tokens []schemas.Token, token schemas.Token) bool {
	for _, a := range tokens {
		if a.Id == token.Id {
			return true
		}
	}
	return false
}

func (service *tokenService) DeleteUserToken(
	token string,
	tokenToDelete struct{ Id uint64 },
) (deletedToken schemas.Token, err error) {
	user, err := service.serviceUser.GetUserInfo(token)
	if err != nil {
		return deletedToken, fmt.Errorf("can't get user info: %w", err)
	}
	userTokenList, err := service.repository.FindByUserId(user.Id)
	if err != nil {
		return deletedToken, fmt.Errorf("can't find tokens by user id: %w", err)
	}
	tokenToDeleteDatabase, err := service.repository.FindById(tokenToDelete.Id)
	if err != nil {
		return deletedToken, fmt.Errorf("can't find token by id: %w", err)
	}
	if containsToken(userTokenList, tokenToDeleteDatabase) {
		// can't delete oauth login token
		if tokenToDeleteDatabase.Id != user.TokenId {
			err = service.repository.Delete(tokenToDeleteDatabase)
			if err != nil {
				return deletedToken, fmt.Errorf("can't delete token: %w", err)
			}
			return tokenToDeleteDatabase, nil
		} else {
			return deletedToken, fmt.Errorf("can't delete token: %w", schemas.ErrTokenBelongToUser)
		}
	} else {
		return deletedToken, fmt.Errorf("token not found")
	}
}
