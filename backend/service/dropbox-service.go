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

	"area/repository"
	"area/schemas"
)

// Constructor

type DropboxService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	GetUserAllFolderAndFileList(
		userDropboxToken string,
	) (fileList []schemas.DropboxEntry, err error)
	GetUserFolderAndFileList(
		userDropboxToken string, path string,
	) (folderAndFileList []schemas.DropboxEntry, err error)
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
	defaultValue := schemas.DropboxActionUpdateInFolder{
		Path: "",
	}
	actionUpdateInFolder, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Dropbox,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:               string(schemas.UpdateInFolder),
			Description:        "This reaction save content from a URL to a file in Dropbox",
			Service:            service.serviceInfo,
			Option:             actionUpdateInFolder,
			MinimumRefreshRate: 10,
		},
	}
}

func (service *dropboxService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.DropboxSaveUrlReactionOption{
		Path: "",
		URL:  "",
	}
	saveUrlReactionOption, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Dropbox,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.SaveUrl),
			Description: "This reaction save content from a URL to a file in Dropbox",
			Service:     service.serviceInfo,
			Option:      saveUrlReactionOption,
		},
	}
}

func (service *dropboxService) FindActionByName(
	name string,
) func(c chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.UpdateInFolder):
		return service.DropboxActionUpdateInFolder
	default:
		return nil
	}
}

func (service *dropboxService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {
	case string(schemas.SaveUrl):
		return service.DropboxReactionSaveUrl
	default:
		return nil
	}
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
	return service.GetUserFolderAndFileList(userDropboxToken, "")
}

func (service *dropboxService) GetUserFolderAndFileList(
	userDropboxToken string, path string,
) (folderAndFileList []schemas.DropboxEntry, err error) {
	ctx := context.Background()

	// Prepare the request body
	reqBody := `{"path": "` + path + `","recursive": true}`

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

func (service *dropboxService) SaveUrl(
	userDropboxToken string, path string, url string,
) (fileJobStatus schemas.DropboxSaveUrlResult, err error) {
	ctx := context.Background()

	if path == "" {
		return fileJobStatus, fmt.Errorf("path cannot be empty")
	}

	if url == "" {
		return fileJobStatus, fmt.Errorf("url cannot be empty")
	}

	if path[0] != '/' {
		path = "/" + path
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Prepare the request body
	reqBody := `{"path": "` + path + `","url": "` + url + `"}`

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.dropboxapi.com/2/files/save_url",
		strings.NewReader(reqBody),
	)
	if err != nil {
		return fileJobStatus, fmt.Errorf("unable to create request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+userDropboxToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fileJobStatus, fmt.Errorf("unable to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to avoid resource leaks

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, _ := io.ReadAll(resp.Body)
		return fileJobStatus, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	// Decode the JSON response into the result struct
	err = json.NewDecoder(resp.Body).Decode(&fileJobStatus)
	if err != nil {
		return fileJobStatus, fmt.Errorf("unable to decode response: %w", err)
	}

	return fileJobStatus, nil
}

func (service *dropboxService) SaveUrlCheckJobStatus(
	userDropboxToken string, saveUrlResult schemas.DropboxSaveUrlResult,
) (saveUrlFile schemas.DropboxEntry, err error) {
	ctx := context.Background()

	// Prepare the request body
	reqBody := `{"async_job_id":"` + saveUrlResult.AsyncJobID + `"}`

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.dropboxapi.com/2/files/save_url/check_job_status",
		strings.NewReader(reqBody),
	)
	if err != nil {
		return saveUrlFile, fmt.Errorf("unable to create request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+userDropboxToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return saveUrlFile, fmt.Errorf("unable to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to avoid resource leaks

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, _ := io.ReadAll(resp.Body)
		return saveUrlFile, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	// Decode the JSON response into the result struct
	err = json.NewDecoder(resp.Body).Decode(&saveUrlFile)
	if err != nil {
		return saveUrlFile, fmt.Errorf("unable to decode response: %w", err)
	}

	return saveUrlFile, nil
}

func (service *dropboxService) IsEntryUpdate(
	folderAndFileList []schemas.DropboxEntry,
	date time.Time,
) bool {
	for _, entry := range folderAndFileList {
		if entry.ClientModified.After(date) {
			return true
		}
	}

	return false
}

// Actions functions

func (service *dropboxService) DropboxActionUpdateInFolder(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	// Find the token of the user
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Action.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return
	}

	databaseStored := schemas.DropboxActionUpdateInFolderStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.DropboxActionUpdateInFolderStorage{
				Time: time.Now(),
			}
			area.StorageVariable, err = json.Marshal(databaseStored)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return
			}
		}
	}

	if databaseStored.Time.IsZero() {
		println("initializing storage variable")
		databaseStored = schemas.DropboxActionUpdateInFolderStorage{
			Time: time.Now(),
		}
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
	}

	// Unmarshal the option
	optionJSON := schemas.DropboxActionUpdateInFolder{}

	err = json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	if optionJSON.Path[0] == '/' {
		optionJSON.Path = optionJSON.Path[1:]
	}

	fileAndFolder, err := service.GetUserFolderAndFileList(token.Token, optionJSON.Path)
	if err != nil {
		fmt.Println(err)
	}

	if service.IsEntryUpdate(fileAndFolder, databaseStored.Time) {
		response := "new update in " + optionJSON.Path + " folder"
		databaseStored.Time = time.Now()
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
		println(response)
		channel <- response
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions

func (service *dropboxService) DropboxReactionSaveUrl(
	option json.RawMessage,
	area schemas.Area,
) string {
	optionJSON := schemas.DropboxSaveUrlReactionOption{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal temperature option: " + err.Error())
		time.Sleep(time.Second)
		return "error unmarshal temperature option: " + err.Error()
	}

	// Find the token of the user
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return "Error finding token" + err.Error()
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return "Error: Token not found"
	}

	fileJobStatus, err := service.SaveUrl(token.Token, optionJSON.Path, optionJSON.URL)
	if err != nil {
		fmt.Println(err)
	}

	saveFile, err := service.SaveUrlCheckJobStatus(token.Token, fileJobStatus)
	if err != nil {
		fmt.Println(err)
	}

	return "create file " + saveFile.PathDisplay + " that save content from " + optionJSON.URL
}
