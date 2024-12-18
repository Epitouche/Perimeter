package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"area/repository"
	"area/schemas"
)

type SpotifyService interface {
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
	SpotifyReactionPlayMusic(option string, idArea uint64)
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	GetActionsName() []string
	GetReactionsName() []string
}

type spotifyService struct {
	repository        repository.SpotifyRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	actionsName       []string
	reactionsName     []string
	serviceInfo       schemas.Service
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
		serviceInfo: schemas.Service{
			Name:        schemas.Spotify,
			Description: "This service is a music service",
		},
	}
}

func (service *spotifyService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *spotifyService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrSpotifyClientIdNotSet
	}

	clientSecret := os.Getenv("SPOTIFY_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrSpotifySecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/spotify"

	apiURL := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to create request because %w",
			err,
		)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to make request because %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		println("Status code", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("body: %+v\n", body)
		return schemas.Token{}, fmt.Errorf(
			"unable to get token because %v",
			resp.Status,
		)
	}

	var result schemas.SpotifyTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	if result.AccessToken == "" {
		fmt.Printf("Token exchange failed. Response body: %v\n", resp.Body)
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

func (service *spotifyService) GetUserInfo(accessToken string) (user schemas.User, err error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	println("accessToken", accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to make request because %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errorResponse := schemas.SpotifyErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return schemas.User{}, fmt.Errorf(
				"unable to decode error response because %w",
				err,
			)
		}
		resp.Body.Close()
		return schemas.User{}, fmt.Errorf(
			"unable to get user info because %v %v",
			errorResponse.Error.Status,
			errorResponse.Error.Message,
		)
	}

	result := schemas.SpotifyUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	user = schemas.User{
		Username: result.DisplayName,
		Email:    result.Email,
	}

	return user, nil
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
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return
	}

	token := service.tokenRepository.FindByUserIdAndServiceId(area.UserId, area.Reaction.ServiceId)
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return
	}

	apiURL := "https://api.spotify.com/v1/me/player/play"

	body := `{
		"context_uri": "spotify:album:5ht7ItJgpBH7W6vJ5BqpPr",
		"offset": {
			"position": 5
		},
		"position_ms": 0
	}`

	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func (service *spotifyService) GetServiceActionInfo() []schemas.Action {
	// service.actionsName = append(service.actionsName, )
	return []schemas.Action{}
}

func (service *spotifyService) GetServiceReactionInfo() []schemas.Reaction {
	service.reactionsName = append(service.reactionsName, string(schemas.PlayMusic))
	return []schemas.Reaction{
		{
			Name:        string(schemas.PlayMusic),
			Description: "This reaction will play music",
			Service:     service.serviceRepository.FindByName(schemas.Spotify),
			Option:      "{}",
		},
	}
}

func (service *spotifyService) GetActionsName() []string {
	return service.actionsName
}

func (service *spotifyService) GetReactionsName() []string {
	return service.reactionsName
}
