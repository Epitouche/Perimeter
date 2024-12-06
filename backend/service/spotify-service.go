package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"area/repository"
	"area/schemas"
)

type SpotifyService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.SpotifyTokenResponse, error)
	GetUserInfo(accessToken string) (schemas.SpotifyUserInfo, error)
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
	SpotifyReactionPlayMusic(option string, idArea uint64)
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
}

type spotifyService struct {
	repository        repository.SpotifyRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
}

func NewSpotifyService(
	githubTokenRepository repository.SpotifyRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) SpotifyService {
	return &spotifyService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
	}
}

func (service *spotifyService) AuthGetServiceAccessToken(
	code string,
	path string,
) (schemas.SpotifyTokenResponse, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientID == "" {
		return schemas.SpotifyTokenResponse{}, fmt.Errorf("SPOTIFY_CLIENT_ID is not set")
	}

	clientSecret := os.Getenv("SPOTIFY_SECRET")
	if clientSecret == "" {
		return schemas.SpotifyTokenResponse{}, fmt.Errorf("SPOTIFY_SECRET is not set")
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.SpotifyTokenResponse{}, fmt.Errorf("BACKEND_PORT is not set")
	}

	redirectURI := "http://localhost:8081/services/spotify"

	apiURL := "https://accounts.spotify.com/api/token"

	// println("redirectURI", redirectURI)
	// println("apiURL", apiURL)
	// println("code", code)

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return schemas.SpotifyTokenResponse{}, fmt.Errorf(
			"unable to create request because %w",
			err,
		)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	client := &http.Client{
		Timeout: time.Second * 30, // Adjust the timeout as needed
	}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.SpotifyTokenResponse{}, fmt.Errorf("unable to make request because %w", err)
	}

	// println("resp")
	// fmt.Printf("%+v\n", resp)
	// println("resp.header")
	// fmt.Printf("%+v\n", resp.Header)
	// println("resp.Body")
	// fmt.Printf("%+v\n", resp.Body)

	var result schemas.SpotifyTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.SpotifyTokenResponse{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()
	return result, nil
}

func (service *spotifyService) GetUserInfo(accessToken string) (schemas.SpotifyUserInfo, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return schemas.SpotifyUserInfo{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.SpotifyUserInfo{}, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.SpotifyUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.SpotifyUserInfo{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()
	return result, nil
}

func (service *spotifyService) FindActionbyName(
	name string,
) func(c chan string, option string, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *spotifyService) FindReactionbyName(name string) func(option string, idArea uint64) {
	switch name {
	case string(schemas.PlayMusic):
		return service.SpotifyReactionPlayMusic
	default:
		return nil
	}
}

func (service *spotifyService) SpotifyReactionPlayMusic(option string, idArea uint64) {
	// Find the area
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("area: %+v\n", area)
	token := service.tokenRepository.FindByUserIdAndServiceId(area.UserId, area.Reaction.ServiceId)
	fmt.Printf("token: %+v\n", token)
}

func (service *spotifyService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{}
}

func (service *spotifyService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{
		{
			Name:        string(schemas.PlayMusic),
			Description: "This reaction will play music",
			Service:     service.serviceRepository.FindByName(schemas.Spotify),
			Option:      "{}",
		},
	}
}
