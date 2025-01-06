package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type DiscordService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	GetActionsName() []string
	GetReactionsName() []string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	GetUserFileList(userDropboxToken string) (fileList []schemas.DropboxFile, err error)
	// Actions functions
	// Reactions functions
}

type discordService struct {
	repository        repository.DiscordRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	actionName        []string
	reactionName      []string
	serviceInfo       schemas.Service
}

func NewDiscordService(
	githubTokenRepository repository.DiscordRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) DiscordService {
	return &discordService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Discord,
			Description: "This service is a messaging platform.",
			Oauth:       true,
			Color:       "#001DDA",
			Icon:        "https://api.iconify.design/mdi:discord.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *discordService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *discordService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{}
}

func (service *discordService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{}
}

func (service *discordService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *discordService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	default:
		return nil
	}
}

func (service *discordService) GetActionsName() []string {
	return service.actionName
}

func (service *discordService) GetReactionsName() []string {
	return service.reactionName
}

// Service specific functions

func (service *discordService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("DROPBOX_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrDropboxClientIdNotSet
	}

	clientSecret := os.Getenv("DROPBOX_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrDropboxSecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/dropbox"

	apiURL := "https://api.dropboxapi.com/oauth2/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to make request because %w", err)
	}

	var result schemas.DropboxTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	if (result.AccessToken == "") || (result.TokenType == "") {
		return schemas.Token{}, schemas.ErrAccessTokenNotFoundInResponse
	}

	resp.Body.Close()

	token = schemas.Token{
		Token:        result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpireAt:     time.Now().Add(time.Duration(result.ExpiresIn) * time.Second),
	}
	return token, nil
}

func (service *discordService) GetUserInfo(
	accessToken string,
) (user schemas.User, err error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://discord.com/api/v8/users/@me",
		nil,
	)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.DropboxUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	user = schemas.User{
		Email:    result.Email,
		Username: result.Name.DisplayName,
	}

	return user, nil
}

func (service *discordService) GetUserFileList(
	userDropboxToken string,
) (fileList []schemas.DropboxFile, err error) {
	ctx := context.Background()

	reqBody := `{"limit": 100}`

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.dropboxapi.com/2/file_requests/list_v2",
		strings.NewReader(reqBody),
	)
	if err != nil {
		return fileList, fmt.Errorf("unable to create request because %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+userDropboxToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fileList, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.DropboxListFileRequestsV2Result{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fileList, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	fileList = append(fileList, result.FileRequests...)

	return fileList, nil
}

func (service *discordService) GetUserFileCount(
	userDropboxToken string,
) (numberFile uint64, err error) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.dropboxapi.com/2/file_requests/count", nil,
	)
	if err != nil {
		return numberFile, fmt.Errorf("unable to create request because %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+userDropboxToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return numberFile, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.DropboxCountFileRequestsResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return numberFile, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	numberFile = result.FileRequestCount
	return numberFile, nil
}

// Actions functions

// Reactions functions
