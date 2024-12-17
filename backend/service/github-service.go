package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"area/repository"
	"area/schemas"
)

type GithubService interface {
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
}

type githubService struct {
	repository repository.GithubRepository
}

func NewGithubService(
	githubTokenRepository repository.GithubRepository,
) GithubService {
	return &githubService{
		repository: githubTokenRepository,
	}
}

func (service *githubService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrGithubClientIdNotSet
	}

	clientSecret := os.Getenv("GITHUB_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrGithubSecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/github"

	apiURL := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", apiURL, nil)
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

	var result schemas.GitHubTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()

	token = schemas.Token{
		Token: result.AccessToken,
		// RefreshToken:  result.RefreshToken,
		// ExpireAt: result.ExpiresIn,
	}
	return token, nil
}

func (service *githubService) GetUserInfo(accessToken string) (user schemas.User, err error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
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

	result := schemas.GithubUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	user = schemas.User{
		Username: result.Login,
		Email:    result.Email,
	}
	return user, nil
}
