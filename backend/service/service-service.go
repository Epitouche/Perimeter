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

type ServiceService interface {
	FindAll() (allServices []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	GetAllServices() (allServicesJSON []schemas.ServiceJSON, err error)
	GetServices() []interface{}
	GetServicesInfo() (allService []schemas.Service, err error)
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	FindServiceByName(name string) schemas.Service
	RedirectToServiceOauthPage(
		serviceName schemas.ServiceName,
		oauthUrl string,
		scope string,
	) (authURL string, err error)
	HandleServiceCallback(
		code string,
		authorization string,
		serviceName schemas.ServiceName,
		authGetServiceAccessToken func(code string) (schemas.Token, error),
		serviceUser UserService,
		getUserInfo func(token string) (userInfo schemas.User, err error),
		tokenService TokenService,
	) (string, error)
	HandleServiceCallbackMobile(
		authorization string,
		serviceName schemas.ServiceName,
		credentials schemas.MobileTokenRequest,
		serviceUser UserService,
		getUserInfo func(token string) (userInfo schemas.User, err error),
		tokenService TokenService,
	) (string, error)
	GetServiceById(serverId uint64) schemas.Service
}

type ServiceInterface interface {
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	GetServiceInfo() schemas.Service
}

type serviceService struct {
	repository repository.ServiceRepository
	allService []interface{}
}

func NewServiceService(
	repository repository.ServiceRepository,
	timerService TimerService,
	spotifyService SpotifyService,
	gmailService GmailService,
	githubService GithubService,
	dropboxService DropboxService,
	microsoftService MicrosoftService,
	openweathermapService OpenweathermapService,
) ServiceService {
	newService := serviceService{
		repository: repository,
		allService: []interface{}{
			spotifyService,
			timerService,
			gmailService,
			githubService,
			dropboxService,
			microsoftService,
			openweathermapService,
		},
	}
	newService.InitialSaveService()
	return &newService
}

func (service *serviceService) InitialSaveService() {
	for _, oneService := range service.allService {
		serviceByName, err := service.repository.FindAllByName(
			oneService.(ServiceInterface).GetServiceInfo().Name,
		)
		if err != nil {
			println(fmt.Errorf("unable to find service by name because %w", err))
		}
		if len(serviceByName) == 0 {
			service.repository.Save(oneService.(ServiceInterface).GetServiceInfo())
		}
	}
}

func (service *serviceService) RedirectToServiceOauthPage(
	serviceName schemas.ServiceName,
	oauthUrl string,
	scope string,
) (authURL string, err error) {
	clientID := ""

	switch serviceName {
	case schemas.Spotify:
		clientID = os.Getenv("SPOTIFY_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrSpotifyClientIdNotSet
		}
	case schemas.Gmail:
		clientID = os.Getenv("GMAIL_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrGmailClientIdNotSet
		}
	case schemas.Github:
		clientID = os.Getenv("GITHUB_CLIENT_ID")
		if clientID == "" {
			return "", schemas.ErrGithubClientIdNotSet
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
	redirectURI := "http://localhost:" + frontendPort + "/services/" + strings.ToLower(
		string(serviceName),
	)
	authURL = oauthUrl +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=" + scope +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}

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

func (service *serviceService) FindAll() (allServices []schemas.Service) {
	allServices, err := service.repository.FindAll()
	if err != nil {
		fmt.Println("Error when get all services")
	}
	return allServices
}

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

func (service *serviceService) FindByName(serviceName schemas.ServiceName) schemas.Service {
	foundService, err := service.repository.FindByName(serviceName)
	if err != nil {
		fmt.Println("Error when get service by name")
	}
	return foundService
}

func (service *serviceService) GetServices() []interface{} {
	return service.allService
}

func (service *serviceService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	for _, service := range service.allService {
		if service.(ServiceInterface).FindActionbyName(name) != nil {
			return service.(ServiceInterface).FindActionbyName(name)
		}
	}
	return nil
}

func (service *serviceService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	for _, service := range service.allService {
		if service.(ServiceInterface).FindReactionbyName(name) != nil {
			return service.(ServiceInterface).FindReactionbyName(name)
		}
	}
	return nil
}

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
