package service

import (
	"bytes"
	"context"
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

// Constructor

type GmailService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	// Token operations
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	// Reactions functions
	GmailReactionSendMail(option json.RawMessage, idArea uint64) string
}

type gmailService struct {
	repository        repository.GmailRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	serviceInfo       schemas.Service
}

func NewGmailService(
	repository repository.GmailRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) GmailService {
	return &gmailService{
		repository:        repository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Google,
			Description: "This service is a mail service",
			Oauth:       true,
			Color:       "#E60000",
			Icon:        "https://api.iconify.design/mdi:google.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *gmailService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *gmailService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *gmailService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	case string(schemas.SendMail):
		println("SendMail")
		return service.GmailReactionSendMail
	default:
		return nil
	}
}

func (service *gmailService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{}
}

func (service *gmailService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.GmailReactionSendMailOption{
		To:      "",
		Subject: "",
		Body:    "",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Google,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.SendMail),
			Description: "Send an email",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

// Service specific functions

func (service *gmailService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("GMAIL_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrGmailClientIdNotSet
	}

	clientSecret := os.Getenv("GMAIL_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrGmailSecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/gmail"

	apiURL := "https://oauth2.googleapis.com/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, nil)
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

	var result schemas.GmailTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	if (result.AccessToken == "") || (result.TokenType == "") {
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

func GetUserGmailProfile(accessToken string) (result schemas.GmailProfile, err error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
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
	ctx := context.Background()
	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
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
) (user schemas.User, err error) {
	gmailProfile, err := GetUserGmailProfile(accessToken)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get gmail profile because %w", err)
	}

	googleProfile, err := GetUserGoogleProfile(accessToken)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get google profile because %w", err)
	}

	user = schemas.User{
		Email:    gmailProfile.EmailAddress,
		Username: googleProfile.Names[0].DisplayName,
	}

	return user, nil
}

// Actions functions

// Reactions functions

func (service *gmailService) GmailReactionSendMail(
	option json.RawMessage,
	idArea uint64,
) string {
	optionJSON := schemas.GmailReactionSendMailOption{}

	println("gmail option: " + string(option))

	err := json.Unmarshal(option, &optionJSON)
	if err != nil {
		println("error unmarshal gmail option: " + err.Error())
		time.Sleep(time.Second)
		return "Error unmarshal gmail option" + err.Error()
	}

	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return "Error finding area" + err.Error()
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return "Error finding token" + err.Error()
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return "Error: Token not found"
	}

	// TODO check if the email is valid or not
	// TODO check if the subject is valid or not
	// TODO check if the body is valid or not

	apiURL := "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"

	email := []byte("From: me@example.com\r\n" +
		"To: " + optionJSON.To + "\r\n" +
		"Subject: " + optionJSON.Subject + "\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n\r\n" +
		optionJSON.Body +
		"\r\n")

	raw := base64.URLEncoding.EncodeToString(email)

	body := fmt.Sprintf(`{"raw": "%s"}`, raw)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		apiURL,
		bytes.NewBuffer([]byte(body)),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request" + err.Error()
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf(
			"Failed to send email. Status: %s, Response: %s\n",
			resp.Status,
			string(respBody),
		)
		return "Failed to send email"
	}

	fmt.Println("Email sent successfully!")
	return "Email sent successfully!"
}
