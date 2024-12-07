package service

import (
	"bytes"
	"encoding/base64"
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

type GmailService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GmailTokenResponse, error)
	GetUserInfo(accessToken string) (result schemas.GmailUserInfo, err error)
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
	GetActionsName() []string
	GetReactionsName() []string
	GmailReactionSendMail(option string, idArea uint64)
	// Token operations
}

type gmailService struct {
	repository        repository.GmailRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	actionName        []string
	reactionName      []string
}

func NewGmailService(
	githubTokenRepository repository.GmailRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) GmailService {
	return &gmailService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
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

	redirectURI := "http://localhost:8081/services/gmail"

	apiURL := "https://oauth2.googleapis.com/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

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

	if (result.AccessToken == "") || (result.TokenType == "") {
		return schemas.GmailTokenResponse{}, fmt.Errorf("access token not found in response")
	}

	resp.Body.Close()
	return result, nil
}

func GetUserGmailProfile(accessToken string) (result schemas.GmailProfile, err error) {
	// Create a new HTTP request
	req, err := http.NewRequest(
		"GET",
		"https://gmail.googleapis.com/gmail/v1/users/me/profile",
		nil,
	)
	if err != nil {
		return schemas.GmailProfile{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GmailProfile{}, fmt.Errorf("unable to make request because %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GmailProfile{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()
	return result, nil
}

func GetUserGoogleProfile(accessToken string) (result schemas.GoogleProfile, err error) {
	// Create a new HTTP request
	req, err := http.NewRequest(
		"GET",
		"https://people.googleapis.com/v1/people/me?personFields=names",
		nil,
	)
	if err != nil {
		return schemas.GoogleProfile{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GoogleProfile{}, fmt.Errorf("unable to make request because %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GoogleProfile{}, fmt.Errorf("unable to decode response because %w", err)
	}
	resp.Body.Close()
	return result, nil
}

func (service *gmailService) GetUserInfo(
	accessToken string,
) (result schemas.GmailUserInfo, err error) {
	gmailProfile, err := GetUserGmailProfile(accessToken)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to get gmail profile because %w", err)
	}
	googleProfile, err := GetUserGoogleProfile(accessToken)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to get google profile because %w", err)
	}
	result.Email = gmailProfile.EmailAddress
	result.Login = googleProfile.Names[0].DisplayName

	return result, nil
}

func (service *gmailService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{}
}

func (service *gmailService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{
		{
			Name:        string(schemas.SendMail),
			Description: "Send an email",
			Service:     service.serviceRepository.FindByName(schemas.Gmail),
			Option:      "{}",
		},
	}
}

func (service *gmailService) FindActionbyName(name string) func(c chan string, option string, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *gmailService) FindReactionbyName(name string) func(option string, idArea uint64) {
	switch name {
	case string(schemas.SendMail):
		println("SendMail")
		return service.GmailReactionSendMail
	default:
		return nil
	}
}

func (service *gmailService) GetActionsName() []string {
	return service.actionName
}

func (service *gmailService) GetReactionsName() []string {
	return service.reactionName
}

func (service *gmailService) GmailReactionSendMail(option string, idArea uint64) {
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

	apiURL := "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"

	email := []byte("From: me@example.com\r\n" +
		"To: recipient@example.com\r\n" +
		"Subject: Hello World\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n\r\n" +
		"<h1>Hello, World!</h1>\r\n")

	raw := base64.URLEncoding.EncodeToString(email)

	body := fmt.Sprintf(`{"raw": "%s"}`, raw)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(body)))
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("Failed to send email. Status: %s, Response: %s\n", resp.Status, string(respBody))
		return
	}

	fmt.Println("Email sent successfully!")
}
