package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Epitouche/Perimeter/repository"
	"github.com/Epitouche/Perimeter/schemas"
)

// Constructor

type DropboxService interface {
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
	GetUserAllFolderAndFileList(
		userDropboxToken string,
	) (fileList []schemas.DropboxEntry, err error)
	GetUserFileList(
		folderAndFileList []schemas.DropboxEntry,
	) (fileList []schemas.DropboxEntry)
	GetUserFolderList(
		folderAndFileList []schemas.DropboxEntry,
	) (fileList []schemas.DropboxEntry)
	CountDropboxEntry(
		folderAndFileList []schemas.DropboxEntry,
	) uint64
	GetPathDisplayDropboxEntry(
		folderAndFileList []schemas.DropboxEntry,
	) (pathDisplay []string)
	// Actions functions
	// Reactions functions
}

type dropboxService struct {
	repository        repository.DropboxRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	actionName        []string
	reactionName      []string
	serviceInfo       schemas.Service
}

func NewDropboxService(
	githubTokenRepository repository.DropboxRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) DropboxService {
	return &dropboxService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Dropbox,
			Description: "This service is a file storage service",
			Oauth:       true,
			Color:       "#001DDA",
			Icon:        "https://api.iconify.design/mdi:dropbox.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *dropboxService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *dropboxService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{}
}

func (service *dropboxService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{}
}

func (service *dropboxService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *dropboxService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	default:
		return nil
	}
}

func (service *dropboxService) GetActionsName() []string {
	return service.actionName
}

func (service *dropboxService) GetReactionsName() []string {
	return service.reactionName
}

// Service specific functions

func (service *dropboxService) AuthGetServiceAccessToken(
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

func (service *dropboxService) GetUserInfo(
	accessToken string,
) (user schemas.User, err error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.dropboxapi.com/2/users/get_current_account",
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

func (service *dropboxService) GetUserAllFolderAndFileList(
	userDropboxToken string,
) (folderAndFileList []schemas.DropboxEntry, err error) {
	ctx := context.Background()

	// Prepare the request body
	reqBody := `{"path": "","recursive": true}`

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.dropboxapi.com/2/files/list_folder",
		strings.NewReader(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+userDropboxToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to avoid resource leaks

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	println("Response status code: ", resp.StatusCode)

	// Decode the JSON response into the result struct
	result := schemas.DropboxListFolderResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	// Append the retrieved files to the file list
	folderAndFileList = result.Entries

	return folderAndFileList, nil
}

func (service *dropboxService) GetUserFileList(
	folderAndFileList []schemas.DropboxEntry,
) (fileList []schemas.DropboxEntry) {
	for _, entry := range folderAndFileList {
		if entry.Tag == "file" {
			fileList = append(fileList, entry)
		}
	}

	return fileList
}

func (service *dropboxService) GetUserFolderList(
	folderAndFileList []schemas.DropboxEntry,
) (fileList []schemas.DropboxEntry) {
	for _, entry := range folderAndFileList {
		if entry.Tag == "folder" {
			fileList = append(fileList, entry)
		}
	}

	return fileList
}

func (service *dropboxService) CountDropboxEntry(
	folderAndFileList []schemas.DropboxEntry,
) uint64 {
	return uint64(len(folderAndFileList))
}

func (service *dropboxService) GetPathDisplayDropboxEntry(
	folderAndFileList []schemas.DropboxEntry,
) (pathDisplay []string) {
	for _, entry := range folderAndFileList {
		pathDisplay = append(pathDisplay, entry.PathDisplay)
	}
	return pathDisplay
}

// Actions functions

// Reactions functions
