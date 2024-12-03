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

type GmailService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GmailTokenResponse, error)
	GetUserInfo(accessToken string) (schemas.GmailUserInfo, error)
	// Token operations
}

type gmailService struct {
	repository repository.GmailRepository
}

func NewGmailService(
	githubTokenRepository repository.GmailRepository,
) GmailService {
	return &gmailService{
		repository: githubTokenRepository,
	}
}

func (service *gmailService) AuthGetServiceAccessToken(
	code string,
	path string,
) (schemas.GmailTokenResponse, error) {
	clientID := os.Getenv("GMAIL_CLIENT_ID")
	if clientID == "" {
		return schemas.GmailTokenResponse{}, fmt.Errorf("GMAIL_CLIENT_ID is not set")
	}

	clientSecret := os.Getenv("GMAIL_SECRET")
	if clientSecret == "" {
		return schemas.GmailTokenResponse{}, fmt.Errorf("GMAIL_SECRET is not set")
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.GmailTokenResponse{}, fmt.Errorf("BACKEND_PORT is not set")
	}

	redirectURI := "http://localhost:" + appPort + path

	// TODO
	apiURL := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return schemas.GmailTokenResponse{}, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30, // Adjust the timeout as needed
	}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GmailTokenResponse{}, fmt.Errorf("unable to make request because %w", err)
	}

	var result schemas.GmailTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GmailTokenResponse{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()
	return result, nil
}

func (service *gmailService) GetUserInfo(accessToken string) (schemas.GmailUserInfo, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.GmailUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()
	return result, nil
}
