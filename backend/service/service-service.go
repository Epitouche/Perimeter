package service

import (
	"os"

	"area/repository"
	"area/schemas"
)

type ServiceService interface {
	FindAll() (allServices []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	GetAllServices() (allServicesJSON []schemas.ServiceJSON, err error)
	GetServices() []interface{}
	GetServicesInfo() (allService []schemas.Service, err error)
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
	FindServiceByName(name string) schemas.Service
	RedirectToServiceOauthPage(
		serviceName schemas.ServiceName,
		oauthUrl string,
		scope string,
	) (authURL string, err error)
}

type ServiceInterface interface {
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
}

type serviceService struct {
	repository        repository.ServiceRepository
	spotifyService    SpotifyService
	timerService      TimerService
	gmailService      GmailService
	allService        []interface{}
	allServiceSchemas []schemas.Service
}

func NewServiceService(
	repository repository.ServiceRepository,
	timerService TimerService,
	spotifyService SpotifyService,
	gmailService GmailService,
) ServiceService {
	newService := serviceService{
		repository:     repository,
		spotifyService: spotifyService,
		timerService:   timerService,
		gmailService:   gmailService,
		allServiceSchemas: []schemas.Service{
			{
				Name:        schemas.Spotify,
				Description: "This service is a music service",
			},
			// {
			// 	Name:        schemas.OpenWeatherMap,
			// 	Description: "This service is a weather service",
			// },
			{
				Name:        schemas.Timer,
				Description: "This service is a time service",
			},
			{
				Name:        schemas.Gmail,
				Description: "This service is a mail service",
			},
		},
		allService: []interface{}{spotifyService, timerService, gmailService},
	}
	newService.InitialSaveService()
	return &newService
}

func (service *serviceService) InitialSaveService() {
	for _, oneService := range service.allServiceSchemas {
		serviceByName := service.repository.FindAllByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
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
	}
	if clientID == "" {
		return "", schemas.ErrNotOauthService
	}

	frontendPort := os.Getenv("FRONTEND_PORT")
	if frontendPort == "" {
		return "", schemas.ErrFrontendPortNotSet
	}

	// Generate the CSRF token
	// state, err := tools.GenerateCSRFToken()
	// if err != nil {
	// 	return "", fmt.Errorf("unable to generate CSRF token because %w", err)
	// }

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	// ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the GitHub authorization URL
	redirectURI := "http://localhost:" + frontendPort + "/services/" + string(serviceName)
	authURL = oauthUrl +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=" + scope +
		"&redirect_uri=" + redirectURI
	return authURL, nil
}

func (service *serviceService) FindAll() (allServices []schemas.Service) {
	return service.repository.FindAll()
}

func (service *serviceService) GetAllServices() (allServicesJSON []schemas.ServiceJSON, err error) {
	allServices := service.repository.FindAll()
	for _, oneService := range allServices {
		println(oneService.Name)
		allServicesJSON = append(allServicesJSON, schemas.ServiceJSON{
			Name: schemas.ServiceName(oneService.Name),
		})
	}
	return allServicesJSON, nil
}

func (service *serviceService) FindByName(serviceName schemas.ServiceName) schemas.Service {
	return service.repository.FindByName(serviceName)
}

func (service *serviceService) GetServices() []interface{} {
	return service.allService
}

func (service *serviceService) FindActionbyName(
	name string,
) func(c chan string, option string, idArea uint64) {
	for _, service := range service.allService {
		if service.(ServiceInterface).FindActionbyName(name) != nil {
			return service.(ServiceInterface).FindActionbyName(name)
		}
	}
	return nil
}

func (service *serviceService) FindReactionbyName(name string) func(option string, idArea uint64) {
	for _, service := range service.allService {
		if service.(ServiceInterface).FindReactionbyName(name) != nil {
			return service.(ServiceInterface).FindReactionbyName(name)
		}
	}
	return nil
}

func (service *serviceService) GetServicesInfo() (allService []schemas.Service, err error) {
	return service.repository.FindAll(), nil
}

func (service *serviceService) FindServiceByName(name string) schemas.Service {
	return service.repository.FindByName(schemas.ServiceName(name))
}
