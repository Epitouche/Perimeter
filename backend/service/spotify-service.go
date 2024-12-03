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
	// Token operations
	SaveToken(token schemas.SpotifyToken) (tokenID uint64, err error)
	Update(token schemas.SpotifyToken) error
	Delete(token schemas.SpotifyToken) error
	FindAll() []schemas.SpotifyToken
	GetTokenById(id uint64) (schemas.SpotifyToken, error)
}

type spotifyService struct {
	repository repository.SpotifyRepository
}

func NewSpotifyService(
	githubTokenRepository repository.SpotifyRepository,
) SpotifyService {
	return &spotifyService{
		repository: githubTokenRepository,
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

	redirectURI := "http://localhost:" + appPort + path

	apiURL := "https://accounts.spotify.com/api/token"

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

func (service *spotifyService) SaveToken(
	token schemas.SpotifyToken,
) (tokenID uint64, err error) {
	tokens := service.repository.FindByAccessToken(token.AccessToken)
	for _, t := range tokens {
		if t.AccessToken == token.AccessToken {
			return t.Id, fmt.Errorf("token already exists")
		}
	}

	service.repository.Save(token)
	tokens = service.repository.FindByAccessToken(token.AccessToken)

	for _, t := range tokens {
		if t.AccessToken == token.AccessToken {
			return t.Id, nil
		}
	}
	return 0, fmt.Errorf("unable to save token")
}

func (service *spotifyService) GetUserInfo(accessToken string) (schemas.SpotifyUserInfo, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
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

func (service *spotifyService) GetTokenById(id uint64) (schemas.SpotifyToken, error) {
	return service.repository.FindById(id), nil
}

func (service *spotifyService) Update(token schemas.SpotifyToken) error {
	service.repository.Update(token)
	return nil
}

func (service *spotifyService) Delete(token schemas.SpotifyToken) error {
	service.repository.Delete(token)
	return nil
}

func (service *spotifyService) FindAll() []schemas.SpotifyToken {
	return service.repository.FindAll()
}
