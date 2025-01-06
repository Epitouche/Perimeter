package service

import (
	"bytes"
	"context"
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

// Constructor

type SpotifyService interface {
	// Service interface functions
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	GetActionsName() []string
	GetReactionsName() []string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	// Reactions functions
	SpotifyReactionPlayMusic(option json.RawMessage, idArea uint64) string
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
			Oauth:       true,
			Color:       "#1DC000",
			Icon:        "https://api.iconify.design/mdi:spotify.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *spotifyService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *spotifyService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *spotifyService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	case string(schemas.PlayMusic):
		return service.SpotifyReactionPlayMusic
	default:
		return nil
	}
}

func (service *spotifyService) GetServiceActionInfo() []schemas.Action {
	// service.actionsName = append(service.actionsName, )
	return []schemas.Action{}
}

func (service *spotifyService) GetServiceReactionInfo() []schemas.Reaction {
	service.reactionsName = append(service.reactionsName, string(schemas.PlayMusic))
	defaultValue := struct{}{}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.PlayMusic),
			Description: "This reaction will play music",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

func (service *spotifyService) GetActionsName() []string {
	return service.actionsName
}

func (service *spotifyService) GetReactionsName() []string {
	return service.reactionsName
}

// Service specific functions

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

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, nil)
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
	ctx := context.Background()
	// Create a new HTTP request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.spotify.com/v1/me",
		nil,
	)
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

// Actions functions

// Reactions functions

func (service *spotifyService) SpotifyReactionPlayMusic(
	option json.RawMessage,
	idArea uint64,
) string {
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return "Error finding area:" + err.Error()
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return "Error finding token:" + err.Error()
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return "Error: Token not found"
	}

	apiURL := "https://api.spotify.com/v1/me/player/play"

	body := `{
		"context_uri": "spotify:album:5ht7ItJgpBH7W6vJ5BqpPr",
		"offset": {
			"position": 5
		},
		"position_ms": 0
	}`
	ctx := context.Background()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		apiURL,
		bytes.NewBuffer([]byte(body)),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request:" + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "Error making request:" + err.Error()
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	return "Response Status:" + resp.Status
}
