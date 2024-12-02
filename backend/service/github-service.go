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

type GithubTokenService interface {
	AuthGetGithubAccessToken(code string, path string) (schemas.GitHubTokenResponse, error)
	GetUserInfo(accessToken string) (schemas.GithubUserInfo, error)
	// Token operations
	SaveToken(token schemas.GithubToken) (tokenID uint64, err error)
	Update(token schemas.GithubToken) error
	Delete(token schemas.GithubToken) error
	FindAll() []schemas.GithubToken
	GetTokenById(id uint64) (schemas.GithubToken, error)
}

type githubTokenService struct {
	repository repository.GithubTokenRepository
}

func NewGithubTokenService(
	githubTokenRepository repository.GithubTokenRepository,
) GithubTokenService {
	return &githubTokenService{
		repository: githubTokenRepository,
	}
}

func (service *githubTokenService) AuthGetGithubAccessToken(
	code string,
	path string,
) (schemas.GitHubTokenResponse, error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return schemas.GitHubTokenResponse{}, fmt.Errorf("GITHUB_CLIENT_ID is not set")
	}

	clientSecret := os.Getenv("GITHUB_SECRET")
	if clientSecret == "" {
		return schemas.GitHubTokenResponse{}, fmt.Errorf("GITHUB_SECRET is not set")
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.GitHubTokenResponse{}, fmt.Errorf("BACKEND_PORT is not set")
	}

	redirectURI := "http://localhost:" + appPort + path

	apiURL := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return schemas.GitHubTokenResponse{}, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30, // Adjust the timeout as needed
	}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GitHubTokenResponse{}, fmt.Errorf("unable to make request because %w", err)
	}

	var result schemas.GitHubTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GitHubTokenResponse{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()
	return result, nil
}

func (service *githubTokenService) SaveToken(
	token schemas.GithubToken,
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

func (service *githubTokenService) GetUserInfo(accessToken string) (schemas.GithubUserInfo, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return schemas.GithubUserInfo{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GithubUserInfo{}, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.GithubUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GithubUserInfo{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()
	return result, nil
}

func (service *githubTokenService) GetTokenById(id uint64) (schemas.GithubToken, error) {
	return service.repository.FindById(id), nil
}

func (service *githubTokenService) Update(token schemas.GithubToken) error {
	service.repository.Update(token)
	return nil
}

func (service *githubTokenService) Delete(token schemas.GithubToken) error {
	service.repository.Delete(token)
	return nil
}

func (service *githubTokenService) FindAll() []schemas.GithubToken {
	return service.repository.FindAll()
}
