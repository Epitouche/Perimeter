package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"area/repository"
	"area/schemas"
)

// TokenService defines the interface for managing tokens.
// It provides methods to save, update, delete, and retrieve tokens.
type TokenService interface {
	// SaveToken saves a new token and returns its ID.
	// Returns an error if the operation fails.
	SaveToken(token schemas.Token) (tokenID uint64, err error)

	// Update modifies an existing token.
	// Returns an error if the operation fails.
	Update(token schemas.Token) error

	// Delete removes a token.
	// Returns an error if the operation fails.
	Delete(token schemas.Token) error

	// FindAll retrieves all tokens.
	// Returns a slice of tokens.
	FindAll() (allServices []schemas.Token)

	// GetTokenById retrieves a token by its ID.
	// Returns the token and an error if the operation fails.
	GetTokenById(id uint64) (schemas.Token, error)

	// GetTokenByUserId retrieves tokens by the user ID.
	// Returns a slice of tokens and an error if the operation fails.
	GetTokenByUserId(userID uint64) ([]schemas.Token, error)

	// DeleteUserToken deletes a user token by its string representation and ID.
	// Returns the deleted token and an error if the operation fails.
	DeleteUserToken(
		token string,
		tokenToDelete struct{ Id uint64 },
	) (deletedToken schemas.Token, err error)
}

type tokenService struct {
	repository  repository.TokenRepository
	serviceUser UserService
}

// NewTokenService creates a new instance of TokenService with the provided
// TokenRepository and UserService. It initializes the tokenService struct
// with the given repository and serviceUser, and returns a pointer to the
// newly created tokenService.
//
// Parameters:
//   - repository: an instance of TokenRepository used for token-related
//     database operations.
//   - serviceUser: an instance of UserService used for user-related operations.
//
// Returns:
//   - TokenService: a pointer to the newly created tokenService instance.
func NewTokenService(repository repository.TokenRepository, serviceUser UserService) TokenService {
	newService := tokenService{
		repository:  repository,
		serviceUser: serviceUser,
	}
	return &newService
}

// SaveToken saves a token to the repository if it does not already exist.
// It first checks if the token already exists in the repository. If it does,
// it returns the existing token ID and an error indicating that the token
// already exists. If the token does not exist, it saves the token to the
// repository and then retrieves the token ID.
//
// Parameters:
//
//	token - the token to be saved.
//
// Returns:
//
//	tokenID - the ID of the saved token.
//	err - an error if the token could not be saved or if it already exists.
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

// GetUserInfo retrieves user information from the GitHub API using the provided access token.
// It sends an HTTP GET request to the GitHub user endpoint and decodes the response into a GmailUserInfo schema.
//
// Parameters:
//   - accessToken: A string containing the access token for authorization.
//
// Returns:
//   - schemas.GmailUserInfo: A struct containing the user's information.
//   - error: An error if the request or decoding fails.
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
	token, err := service.repository.FindById(id)
	if err != nil {
		return schemas.Token{}, err
	}
	return token, nil
}

// GetTokenByUserId retrieves a list of tokens associated with the given user ID.
// It returns a slice of Token schemas and an error if any occurred during the retrieval process.
//
// Parameters:
//   - userID: The ID of the user whose tokens are to be retrieved.
//
// Returns:
//   - []schemas.Token: A slice of Token schemas associated with the user.
//   - error: An error if any occurred during the retrieval process, otherwise nil.
func (service *tokenService) GetTokenByUserId(userID uint64) ([]schemas.Token, error) {
	tokens, err := service.repository.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// Update updates the given token in the repository.
// It takes a Token schema as input and returns an error if the update operation fails.
//
// Parameters:
//
//	token (schemas.Token): The token to be updated.
//
// Returns:
//
//	error: An error object if the update operation fails, otherwise nil.
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

// containsToken checks if a given token is present in a slice of tokens.
// It returns true if the token is found, otherwise it returns false.
//
// Parameters:
// - tokens: A slice of schemas.Token where the search is performed.
// - token: The schemas.Token to search for in the tokens slice.
//
// Returns:
// - bool: true if the token is found in the slice, false otherwise.
func containsToken(tokens []schemas.Token, token schemas.Token) bool {
	for _, a := range tokens {
		if a.Id == token.Id {
			return true
		}
	}
	return false
}

// DeleteUserToken deletes a specific token associated with a user.
// It takes a token string for authentication and a struct containing the ID of the token to delete.
// It returns the deleted token and an error if any occurred during the process.
//
// Parameters:
//   - token: A string representing the user's authentication token.
//   - tokenToDelete: A struct containing the ID of the token to be deleted.
//
// Returns:
//   - deletedToken: The token that was deleted.
//   - err: An error if any occurred during the deletion process.
//
// Errors:
//   - Returns an error if the user information cannot be retrieved.
//   - Returns an error if the tokens associated with the user cannot be found.
//   - Returns an error if the token to delete cannot be found.
//   - Returns an error if the token to delete belongs to the user and cannot be deleted.
//   - Returns an error if the token to delete is not found in the user's token list.
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
