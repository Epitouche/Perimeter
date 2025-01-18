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

// DropboxService defines the interface for interacting with Dropbox services.
// It includes methods for authentication, user information retrieval, file and folder management,
// and specific actions and reactions related to Dropbox.
//
// Methods:
//
// - GetServiceActionInfo: Returns a list of available actions for the service.
// - GetServiceReactionInfo: Returns a list of available reactions for the service.
// - FindActionByName: Finds an action by its name and returns a function to execute the action.
// - FindReactionByName: Finds a reaction by its name and returns a function to execute the reaction.
// - AuthGetServiceAccessToken: Retrieves an access token using an authorization code.
// - GetUserInfo: Retrieves user information using an access token.
// - GetUserAllFolderAndFileList: Retrieves a list of all folders and files for a user.
// - GetUserFolderAndFileList: Retrieves a list of folders and files for a user at a specific path.
// - GetUserFileList: Filters and returns only files from a list of Dropbox entries.
// - GetUserFolderList: Filters and returns only folders from a list of Dropbox entries.
// - CountDropboxEntry: Counts the number of Dropbox entries in a list.
// - GetPathDisplayDropboxEntry: Retrieves the display paths for a list of Dropbox entries.
// - DropboxActionUpdateInFolder: Executes an action to update content in a folder.
// - DropboxReactionSaveUrl: Executes a reaction to save a URL.
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
	DropboxActionUpdateInFolder(channel chan string, option json.RawMessage, area schemas.Area)
	// Reactions functions
	DropboxReactionSaveUrl(option json.RawMessage, area schemas.Area) string
}

// dropboxService is a struct that provides methods to interact with Dropbox services.
// It contains repositories for Dropbox, Service, Area, and Token, as well as service information.
//
// Fields:
// - repository: Interface for Dropbox repository operations.
// - serviceRepository: Interface for service repository operations.
// - areaRepository: Interface for area repository operations.
// - tokenRepository: Interface for token repository operations.
// - serviceInfo: Information about the service, represented by the Service schema.
type dropboxService struct {
	repository        repository.DropboxRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	serviceInfo       schemas.Service
}

// NewDropboxService creates a new instance of DropboxService with the provided repositories.
// It initializes the service with the necessary repositories and service information.
//
// Parameters:
//   - githubTokenRepository: repository.DropboxRepository
//   - serviceRepository: repository.ServiceRepository
//   - areaRepository: repository.AreaRepository
//   - tokenRepository: repository.TokenRepository
//
// Returns:
//   - DropboxService: a new instance of DropboxService
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

// GetServiceInfo returns the service information for the Dropbox service.
// It retrieves the service information from the dropboxService struct and
// returns it as a schemas.Service type.
func (service *dropboxService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

// GetServiceActionInfo retrieves the service action information for Dropbox.
// It constructs a default DropboxActionUpdateInFolder with a predefined path,
// marshals it into JSON, and fetches the service information from the repository.
// If any errors occur during marshalling or fetching the service information,
// they are printed to the console. The function returns a slice of Action
// containing the details of the Dropbox action.
//
// Returns:
//
//	[]schemas.Action: A slice containing the Dropbox action details.
func (service *dropboxService) GetServiceActionInfo() []schemas.Action {
	defaultValue := schemas.DropboxActionUpdateInFolder{
		Path: "folder/subfolder",
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
			Description:        "This action triggers when there is an update in a folder",
			Service:            service.serviceInfo,
			Option:             actionUpdateInFolder,
			MinimumRefreshRate: 10,
		},
	}
}

// GetServiceReactionInfo retrieves the reaction information for the Dropbox service.
// It initializes a default value for the DropboxSaveUrlReactionOption, marshals it into JSON,
// and updates the service information by finding the service by name.
// If any errors occur during marshaling or finding the service, they are printed to the console.
// The function returns a slice of Reaction containing the name, description, service information,
// and the marshaled option for saving a URL to a file in Dropbox.
func (service *dropboxService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.DropboxSaveUrlReactionOption{
		Path: "folder/subfolder/file.txt",
		URL:  "site.com/robot.txt",
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

// FindActionByName returns a function that performs a specific action based on the provided name.
// The returned function takes a channel, a JSON raw message, and an area schema as parameters.
//
// Parameters:
//   - name: The name of the action to find.
//
// Returns:
//   - A function that performs the specified action, or nil if the action name is not recognized.
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

// FindReactionByName returns a function that corresponds to the given reaction name.
// The returned function takes a JSON raw message and an Area schema as parameters and returns a string.
// If the reaction name matches a known reaction, the corresponding function is returned.
// If the reaction name does not match any known reactions, nil is returned.
//
// Parameters:
//   - name: The name of the reaction to find.
//
// Returns:
//   - A function that takes a JSON raw message and an Area schema as parameters and returns a string,
//     or nil if the reaction name does not match any known reactions.
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

// AuthGetServiceAccessToken exchanges an authorization code for a Dropbox access token.
// It retrieves the client ID and client secret from environment variables, constructs
// the request to the Dropbox API, and decodes the response into a Token schema.
//
// Parameters:
//   - code: The authorization code received from Dropbox after user authorization.
//
// Returns:
//   - token: A Token schema containing the access token, refresh token, and expiration time.
//   - err: An error if the process fails at any step, including missing environment variables,
//     request creation, or response decoding.
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

	redirectURI, err := getRedirectURI(service.serviceInfo.Name)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to get redirect URI because %w", err)
	}

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

// GetUserInfo retrieves the current user's information from Dropbox using the provided access token.
// It sends a POST request to the Dropbox API endpoint for getting the current account information.
//
// Parameters:
//   - accessToken: A string containing the OAuth 2.0 access token for authenticating the request.
//
// Returns:
//   - user: A schemas.User struct containing the user's email and display name.
//   - err: An error if the request fails or the response cannot be decoded.
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

// GetUserAllFolderAndFileList retrieves a list of all folders and files for a user from Dropbox.
// It takes the user's Dropbox token as an input and returns a slice of DropboxEntry schemas and an error, if any.
//
// Parameters:
//   - userDropboxToken: A string representing the user's Dropbox access token.
//
// Returns:
//   - folderAndFileList: A slice of DropboxEntry schemas representing the user's folders and files.
//   - err: An error object if an error occurs during the retrieval process.
func (service *dropboxService) GetUserAllFolderAndFileList(
	userDropboxToken string,
) (folderAndFileList []schemas.DropboxEntry, err error) {
	return service.GetUserFolderAndFileList(userDropboxToken, "")
}

// GetUserFolderAndFileList retrieves the list of folders and files from the user's Dropbox account
// for the specified path.
//
// Parameters:
//   - userDropboxToken: The OAuth token for the user's Dropbox account.
//   - path: The path in the Dropbox account to list the folders and files from.
//
// Returns:
//   - folderAndFileList: A slice of DropboxEntry structs representing the folders and files.
//   - err: An error if the request fails or the response cannot be decoded.
//
// The function makes an HTTP POST request to the Dropbox API to list the folders and files
// at the specified path. It sets the necessary headers, including the Authorization header
// with the user's Dropbox token. The response is decoded into a DropboxListFolderResult struct,
// and the entries are returned as a slice of DropboxEntry structs.
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

// GetUserFileList filters the provided list of Dropbox entries and returns only the entries that are files.
//
// Parameters:
//
//	folderAndFileList []schemas.DropboxEntry - A list of Dropbox entries which may include both files and folders.
//
// Returns:
//
//	[]schemas.DropboxEntry - A list of Dropbox entries that are files.
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

// GetUserFolderList filters the provided list of Dropbox entries and returns only the entries that are folders.
//
// Parameters:
//
//	folderAndFileList - A slice of DropboxEntry objects representing files and folders.
//
// Returns:
//
//	A slice of DropboxEntry objects that are folders.
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

// CountDropboxEntry counts the number of entries in the provided list of Dropbox entries.
// It takes a slice of DropboxEntry structs as input and returns the count as a uint64.
//
// Parameters:
//
//	folderAndFileList []schemas.DropboxEntry - A slice of DropboxEntry structs representing the folder and file entries.
//
// Returns:
//
//	uint64 - The count of Dropbox entries in the provided list.
func (service *dropboxService) CountDropboxEntry(
	folderAndFileList []schemas.DropboxEntry,
) uint64 {
	return uint64(len(folderAndFileList))
}

// GetPathDisplayDropboxEntry extracts the PathDisplay field from each DropboxEntry
// in the provided folderAndFileList and returns a slice of these path displays.
//
// Parameters:
//
//	folderAndFileList - A slice of DropboxEntry structs from which to extract the PathDisplay field.
//
// Returns:
//
//	A slice of strings containing the PathDisplay values from the provided DropboxEntry structs.
func (service *dropboxService) GetPathDisplayDropboxEntry(
	folderAndFileList []schemas.DropboxEntry,
) (pathDisplay []string) {
	for _, entry := range folderAndFileList {
		pathDisplay = append(pathDisplay, entry.PathDisplay)
	}
	return pathDisplay
}

// SaveUrl saves a file from a given URL to a specified path in the user's Dropbox.
//
// Parameters:
//   - userDropboxToken: The OAuth token for the user's Dropbox account.
//   - path: The path in the Dropbox where the file should be saved. Must not be empty and should start with '/'.
//   - url: The URL of the file to be saved. Must not be empty and should start with "http://" or "https://".
//
// Returns:
//   - fileJobStatus: A struct containing the result of the save URL operation.
//   - err: An error if the operation fails, otherwise nil.
//
// Errors:
//   - Returns an error if the path or URL is empty.
//   - Returns an error if the HTTP request cannot be created or executed.
//   - Returns an error if the response status code is not 200 OK.
//   - Returns an error if the response cannot be decoded into the result struct.
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

// SaveUrlCheckJobStatus checks the status of a save URL job in Dropbox.
//
// Parameters:
//   - userDropboxToken: The OAuth token for the user's Dropbox account.
//   - saveUrlResult: The result of the save URL operation containing the async job ID.
//
// Returns:
//   - saveUrlFile: The Dropbox entry representing the saved file.
//   - err: An error if the operation fails.
//
// This function sends a request to the Dropbox API to check the status of a save URL job.
// It prepares the request body with the async job ID, sets the necessary headers, and
// makes the HTTP request. If the request is successful, it decodes the JSON response
// into the saveUrlFile struct. If there is an error at any step, it returns the error.
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

// IsEntryUpdate checks if any entry in the provided list of Dropbox entries
// has been modified after the specified date.
//
// Parameters:
//
//	folderAndFileList - a slice of DropboxEntry objects to check for updates
//	date - the time to compare each entry's modification date against
//
// Returns:
//
//	true if any entry in the list has a ClientModified date after the specified date,
//	false otherwise.
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

// DropboxActionUpdateInFolder performs an update action in a specified Dropbox folder.
// It checks for updates in the folder and sends a notification if any updates are found.
//
// Parameters:
// - channel: A channel to send update notifications.
// - option: A JSON raw message containing the options for the action.
// - area: The area schema containing user and action information.
//
// The function performs the following steps:
// 1. Retrieves the user's token for the Dropbox service.
// 2. Initializes or retrieves the storage variable for the action.
// 3. Unmarshals the options for the action.
// 4. Retrieves the list of files and folders in the specified Dropbox path.
// 5. Checks if there are any updates in the folder since the last check.
// 6. Sends a notification if updates are found and updates the storage variable.
// 7. Sleeps for the specified refresh rate before the next check.
//
// If any errors occur during these steps, appropriate error messages are printed and the function returns.
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

// DropboxReactionSaveUrl saves a file from a given URL to Dropbox.
//
// Parameters:
//   - option: A JSON raw message containing the options for the Dropbox save URL reaction.
//   - area: An Area schema containing user and reaction information.
//
// Returns:
//
//	A string message indicating the result of the operation.
//
// The function performs the following steps:
//  1. Unmarshals the JSON options into a DropboxSaveUrlReactionOption struct.
//  2. Finds the user's token for the Dropbox service.
//  3. Uses the token to save the file from the provided URL to the specified path in Dropbox.
//  4. Checks the job status of the save operation.
//  5. Returns a message indicating the path of the saved file and the source URL.
//
// If any error occurs during these steps, an appropriate error message is returned.
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
