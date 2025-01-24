package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"area/repository"
	"area/schemas"
	"area/tools"
)

// ServiceService defines the interface for service-related operations.
type ServiceService interface {
	// FindAll retrieves all services.
	FindAll() (allServices []schemas.Service)

	// FindByName retrieves a service by its name.
	FindByName(serviceName schemas.ServiceName) schemas.Service

	// GetAllServices retrieves all services in JSON format.
	GetAllServices() (allServicesJSON []schemas.ServiceJSON, err error)

	// GetServices retrieves all services as a slice of interfaces.
	GetServices() []interface{}

	// GetServicesInfo retrieves detailed information about all services.
	GetServicesInfo() (allService []schemas.Service, err error)

	// FindActionByName retrieves an action function by its name.
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)

	// FindReactionByName retrieves a reaction function by its name.
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string

	// FindServiceByName retrieves a service by its name.
	FindServiceByName(name string) schemas.Service

	// RedirectToServiceOauthPage generates the OAuth URL for a service.
	RedirectToServiceOauthPage(
		serviceName schemas.ServiceName,
		oauthUrl string,
		scope string,
	) (authURL string, err error)

	// HandleServiceCallback handles the OAuth callback for a service.
	HandleServiceCallback(
		code string,
		authorization string,
		serviceName schemas.ServiceName,
		authGetServiceAccessToken func(code string) (schemas.Token, error),
		serviceUser UserService,
		getUserInfo func(token string) (userInfo schemas.User, err error),
		tokenService TokenService,
	) (string, error)

	// HandleServiceCallbackMobile handles the OAuth callback for a service on mobile.
	HandleServiceCallbackMobile(
		authorization string,
		serviceName schemas.ServiceName,
		credentials schemas.MobileTokenRequest,
		serviceUser UserService,
		getUserInfo func(token string) (userInfo schemas.User, err error),
		tokenService TokenService,
	) (string, error)

	// GetServiceById retrieves a service by its ID.
	GetServiceById(serverId uint64) schemas.Service
}

// ServiceInterface defines the methods that a service must implement to interact with actions and reactions.
// It includes methods to find actions and reactions by name, and to retrieve service information.
type ServiceInterface interface {
	// FindActionByName returns a function that processes an action based on its name.
	// The returned function takes a channel to send strings, a JSON raw message for options, and an area schema.
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)

	// FindReactionByName returns a function that processes a reaction based on its name.
	// The returned function takes a JSON raw message for options and an area schema, and returns a string.
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string

	// GetServiceInfo returns the service information as a Service schema.
	GetServiceInfo() schemas.Service
}

// serviceService is a struct that provides services related to the ServiceRepository.
// It contains a repository field which is an instance of ServiceRepository and an allService
// field which is a slice of interfaces representing all services.
type serviceService struct {
	repository repository.ServiceRepository
	allService []interface{}
}

// NewServiceService creates a new instance of ServiceService with the provided dependencies.
// It initializes the service with the given repository and various service implementations,
// and performs an initial save operation.
//
// Parameters:
//   - repository: an instance of ServiceRepository for data access.
//   - timerService: an instance of TimerService for timer-related operations.
//   - spotifyService: an instance of SpotifyService for Spotify-related operations.
//   - googleService: an instance of GoogleService for Google-related operations.
//   - githubService: an instance of GithubService for GitHub-related operations.
//   - dropboxService: an instance of DropboxService for Dropbox-related operations.
//   - microsoftService: an instance of MicrosoftService for Microsoft-related operations.
//   - openWeatherMapService: an instance of OpenWeatherMapService for weather-related operations.
//
// Returns:
//   - ServiceService: a new instance of ServiceService initialized with the provided dependencies.
func NewServiceService(
	repository repository.ServiceRepository,
	timerService TimerService,
	spotifyService SpotifyService,
	googleService GoogleService,
	githubService GithubService,
	dropboxService DropboxService,
	microsoftService MicrosoftService,
	openWeatherMapService OpenWeatherMapService,
) ServiceService {
	newService := serviceService{
		repository: repository,
		allService: []interface{}{
			spotifyService,
			timerService,
			googleService,
			githubService,
			dropboxService,
			microsoftService,
			openWeatherMapService,
		},
	}
	newService.InitialSaveService()
	return &newService
}

// InitialSaveService iterates over all services in the serviceService instance,
// checks if each service already exists in the repository by its name, and if not,
// saves the service to the repository. If an error occurs during the process of
// finding or saving a service, it prints an error message.
func (service *serviceService) InitialSaveService() {
	for _, oneService := range service.allService {
		serviceByName, err := service.repository.FindAllByName(
			oneService.(ServiceInterface).GetServiceInfo().Name,
		)
		if err != nil {
			println(fmt.Errorf("unable to find service by name because %w", err))
		}
		if len(serviceByName) == 0 {
			err = service.repository.Save(oneService.(ServiceInterface).GetServiceInfo())
			if err != nil {
				println(fmt.Errorf("unable to save service because %w", err))
			}
		}
	}
}

// getRedirectURI constructs the redirect URI for a given service name based on the environment variables.
// It returns the constructed redirect URI or an error if any required environment variable is not set.
//
// Environment Variables:
// - FRONTEND_PORT: The port number of the frontend service.
// - FRONTEND_EXTERNAL_HOST: The external host address of the frontend service.
// - IS_PRODUCTION: A flag indicating whether the environment is production.
//
// Parameters:
// - serviceName: The name of the service for which the redirect URI is being constructed.
//
// Returns:
// - redirectURI: The constructed redirect URI as a string.
// - err: An error if any required environment variable is not set.
func getRedirectURI(
	serviceName schemas.ServiceName,
) (redirectURI string, err error) {
	frontendPort := os.Getenv("FRONTEND_PORT")
	if frontendPort == "" {
		return "", schemas.ErrFrontendPortNotSet
	}
	frontendExternalHost := os.Getenv("FRONTEND_EXTERNAL_HOST")
	if frontendExternalHost == "" {
		return "", schemas.ErrFrontendExternalHostNotSet
	}

	isProd := os.Getenv("IS_PRODUCTION")
	if isProd == "" {
		return "", schemas.ErrIsProductionNotSet
	}

	host := ""
	if isProd == "true" {
		host = "https://" + frontendExternalHost
	} else {
		host = "http://" + frontendExternalHost + ":" + frontendPort
	}

	return host + "/services/" + strings.ToLower(string(serviceName)), nil
}

// RedirectToServiceOauthPage generates an OAuth authorization URL for the specified service.
//
// Parameters:
//   - serviceName: The name of the service for which to generate the OAuth URL.
//   - oauthUrl: The base URL for the OAuth authorization endpoint.
//   - scope: The scope of the OAuth authorization request.
//
// Returns:
//   - authURL: The generated OAuth authorization URL.
//   - err: An error if the client ID for the specified service is not set, if the frontend port is not set,
//     if there is an error generating the CSRF token, or if there is an error getting the redirect URI.
//
// The function retrieves the client ID for the specified service from environment variables. It also retrieves
// the frontend port from environment variables. It generates a CSRF token and constructs the OAuth authorization
// URL using the provided parameters and the retrieved values.
func (service *serviceService) RedirectToServiceOauthPage(
	serviceName schemas.ServiceName,
	oauthUrl string,
	scope string,
) (authURL string, err error) {
	clientID := ""

	isProd := os.Getenv("IS_PRODUCTION")
	if isProd == "" {
		return "", schemas.ErrIsProductionNotSet
	}

	switch serviceName {
	case schemas.Spotify:
		clientID = os.Getenv("SPOTIFY_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrSpotifyClientIdNotSet
		}
	case schemas.Google:
		clientID = os.Getenv("GOOGLE_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrGoogleClientIdNotSet
		}
	case schemas.Github:
		if isProd == "true" {
			clientID = os.Getenv("GITHUB_PRODUCTION_CLIENT_ID")
			if clientID == "" {
				return "", schemas.ErrGithubProductionClientIdNotSet
			}
		} else {
			clientID = os.Getenv("GITHUB_CLIENT_ID")
			if clientID == "" {
				return "", schemas.ErrGithubClientIdNotSet
			}
		}
	case schemas.Dropbox:
		clientID = os.Getenv("DROPBOX_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrDropboxClientIdNotSet
		}
	case schemas.Microsoft:
		clientID = os.Getenv("MICROSOFT_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrMicrosoftClientIdNotSet
		}
	}

	if clientID == "" {
		return "", schemas.ErrNotOauthService
	}

	frontendPort := os.Getenv("FRONTEND_PORT")
	if frontendPort == "" {
		return "", schemas.ErrFrontendPortNotSet
	}

	// Generate the CSRF token
	state, err := tools.GenerateCSRFToken()
	if err != nil {
		return "", fmt.Errorf("unable to generate CSRF token because %w", err)
	}

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	// ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the GitHub authorization URL
	redirectURI, err := getRedirectURI(serviceName)
	if err != nil {
		return "", fmt.Errorf("unable to get redirect URI because %w", err)
	}

	authURL = oauthUrl +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=" + scope +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}

// HandleServiceCallback handles the callback from a service after user authorization.
// It retrieves the access token using the provided code, fetches user information, and
// either logs in or registers the user. It then saves the service token in the database.
//
// Parameters:
//   - code: The authorization code received from the service.
//   - authorization: The authorization header containing the bearer token.
//   - serviceName: The name of the service.
//   - authGetServiceAccessToken: A function to get the service access token using the code.
//   - serviceUser: The user service to interact with user information.
//   - getUserServiceInfo: A function to get user information using the service token.
//   - tokenService: The token service to save the service token.
//
// Returns:
//   - A string containing the bearer token.
//   - An error if any operation fails.
func (service *serviceService) HandleServiceCallback(
	code string,
	authorization string,
	serviceName schemas.ServiceName,
	authGetServiceAccessToken func(code string) (schemas.Token, error),
	serviceUser UserService,
	getUserServiceInfo func(token string) (userInfo schemas.User, err error),
	tokenService TokenService,
) (string, error) {
	authHeader := authorization
	newUser := schemas.User{}
	var bearerToken string

	serviceToken, err := authGetServiceAccessToken(code)
	if err != nil {
		return "", fmt.Errorf("unable to get access token because %w", err)
	}

	if len(authHeader) > len("Bearer ") {
		bearerToken = authHeader[len("Bearer "):]

		newUser, err = serviceUser.GetUserInfo(bearerToken)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
	} else {
		userInfo, err := getUserServiceInfo(serviceToken.Token)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
		newUser = schemas.User{
			Username: userInfo.Username,
			Email:    userInfo.Email,
		}

		bearerTokenLogin, _, err := serviceUser.Login(newUser)
		if err == nil {
			return bearerTokenLogin, nil
		}

		bearerTokenRegister, newUserId, err := serviceUser.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		bearerToken = bearerTokenRegister
		newUser, err = serviceUser.GetUserById(newUserId)
		if err != nil {
			return "", fmt.Errorf("unable to get user by id because %w", err)
		}
	}

	serviceService := service.FindByName(serviceName)

	newServiceToken := schemas.Token{
		Token:        serviceToken.Token,
		RefreshToken: serviceToken.RefreshToken,
		ExpireAt:     serviceToken.ExpireAt,
		Service:      serviceService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := tokenService.SaveToken(newServiceToken)
	if err != nil {
		if errors.Is(err, schemas.ErrTokenAlreadyExists) {
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	if len(authHeader) == 0 {
		newUser.TokenId = tokenId

		err = serviceUser.UpdateUserInfo(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to update user info because %w", err)
		}
	}
	return bearerToken, nil
}

// HandleServiceCallbackMobile handles the callback for mobile services.
// It processes the authorization header and credentials to authenticate or register the user,
// retrieves user information, and saves the service token in the database.
//
// Parameters:
//   - authorization: The authorization header containing the bearer token.
//   - serviceName: The name of the service.
//   - credentials: The mobile token request containing access and refresh tokens.
//   - serviceUser: The user service interface for user-related operations.
//   - getUserInfo: A function to get user information using a token.
//   - tokenService: The token service interface for token-related operations.
//
// Returns:
//   - A string containing the bearer token.
//   - An error if any operation fails.
func (service *serviceService) HandleServiceCallbackMobile(
	authorization string,
	serviceName schemas.ServiceName,
	credentials schemas.MobileTokenRequest,
	serviceUser UserService,
	getUserInfo func(token string) (userInfo schemas.User, err error),
	tokenService TokenService,
) (string, error) {
	authHeader := authorization
	newUser := schemas.User{}
	var bearerToken string
	var err error

	if len(authHeader) > len("Bearer ") {
		bearerToken = authHeader[len("Bearer "):]

		newUser, err = serviceUser.GetUserInfo(bearerToken)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
	} else {
		userInfo, err := getUserInfo(credentials.AccessToken)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
		newUser = schemas.User{
			Username: userInfo.Username,
			Email:    userInfo.Email,
		}

		bearerTokenLogin, _, err := serviceUser.Login(newUser)
		if err == nil {
			return bearerTokenLogin, nil
		}

		bearerTokenRegister, newUserId, err := serviceUser.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		bearerToken = bearerTokenRegister
		newUser, err = serviceUser.GetUserById(newUserId)
		if err != nil {
			return "", fmt.Errorf("unable to get user by id because %w", err)
		}
	}

	actualService := service.FindByName(serviceName)

	newServiceToken := schemas.Token{
		Token:        credentials.AccessToken,
		RefreshToken: credentials.RefreshToken,
		ExpireAt:     credentials.ExpiresIn,
		Service:      actualService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := tokenService.SaveToken(newServiceToken)
	if err != nil {
		if errors.Is(err, schemas.ErrTokenAlreadyExists) {
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	if len(authHeader) == 0 {
		newUser.TokenId = tokenId

		err = serviceUser.UpdateUserInfo(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to update user info because %w", err)
		}
	}
	return bearerToken, nil
}

// FindAll retrieves all services from the repository.
// It returns a slice of schemas.Service containing all the services.
// If an error occurs during the retrieval, it prints an error message.
func (service *serviceService) FindAll() (allServices []schemas.Service) {
	allServices, err := service.repository.FindAll()
	if err != nil {
		fmt.Println("Error when get all services")
	}
	return allServices
}

// GetAllServices retrieves all services from the repository, converts them to JSON format,
// and returns them. If an error occurs during the retrieval process, it prints an error message.
//
// Returns:
//   - allServicesJSON: A slice of ServiceJSON containing all services in JSON format.
//   - err: An error if there was an issue retrieving the services, otherwise nil.
func (service *serviceService) GetAllServices() (allServicesJSON []schemas.ServiceJSON, err error) {
	allServices, err := service.repository.FindAll()
	if err != nil {
		fmt.Println("Error when get all services")
	}
	for _, oneService := range allServices {
		println(oneService.Name)
		allServicesJSON = append(allServicesJSON, schemas.ServiceJSON{
			Name: schemas.ServiceName(oneService.Name),
		})
	}
	return allServicesJSON, nil
}

// FindByName retrieves a service by its name from the repository.
// It takes a serviceName of type schemas.ServiceName as an argument and returns a schemas.Service.
// If an error occurs during the retrieval, it prints an error message to the console.
func (service *serviceService) FindByName(serviceName schemas.ServiceName) schemas.Service {
	foundService, err := service.repository.FindByName(serviceName)
	if err != nil {
		fmt.Println("Error when get service by name")
	}
	return foundService
}

// GetServices retrieves all services.
// It returns a slice of interfaces representing the services.
func (service *serviceService) GetServices() []interface{} {
	return service.allService
}

// FindActionByName searches for an action by its name within the service's list of all services.
// It returns a function that takes a channel, an option in the form of json.RawMessage, and an area of type schemas.Area.
// If no action is found with the given name, it returns nil.
//
// Parameters:
//   - name: The name of the action to search for.
//
// Returns:
//   - A function that takes a channel (chan string), an option (json.RawMessage), and an area (schemas.Area),
//     or nil if no action with the specified name is found.
func (service *serviceService) FindActionByName(
	name string,
) func(c chan string, option json.RawMessage, area schemas.Area) {
	for _, service := range service.allService {
		if service.(ServiceInterface).FindActionByName(name) != nil {
			return service.(ServiceInterface).FindActionByName(name)
		}
	}
	return nil
}

// FindReactionByName searches for a reaction by its name within the service's list of all services.
// If a matching reaction is found, it returns a function that takes a JSON raw message and an area schema,
// and returns a string. If no matching reaction is found, it returns nil.
//
// Parameters:
//   - name: The name of the reaction to search for.
//
// Returns:
//   - A function that takes a JSON raw message and an area schema, and returns a string, if a matching reaction is found.
//   - nil if no matching reaction is found.
func (service *serviceService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	for _, service := range service.allService {
		if service.(ServiceInterface).FindReactionByName(name) != nil {
			return service.(ServiceInterface).FindReactionByName(name)
		}
	}
	return nil
}

// GetServicesInfo retrieves information about all services.
// It returns a slice of Service schemas and an error if any occurs during the retrieval process.
func (service *serviceService) GetServicesInfo() (allService []schemas.Service, err error) {
	return service.repository.FindAll()
}

func (service *serviceService) FindServiceByName(name string) schemas.Service {
	services, err := service.repository.FindByName(schemas.ServiceName(name))
	if err != nil {
		fmt.Println("Error when get service by name")
	}
	return services
}

func (service *serviceService) GetServiceById(id uint64) schemas.Service {
	foundService, err := service.repository.FindById(id)
	if err != nil {
		fmt.Println("Error when get service by id")
	}
	return foundService
}
