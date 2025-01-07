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

type DiscordService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	GetActionsName() []string
	GetReactionsName() []string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	// Reactions functions
	DiscordReactionSendMessage(
		option json.RawMessage,
		idArea uint64,
	) string
}

type discordService struct {
	repository        repository.DiscordRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	actionName        []string
	reactionName      []string
	serviceInfo       schemas.Service
}

func NewDiscordService(
	githubTokenRepository repository.DiscordRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) DiscordService {
	return &discordService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Discord,
			Description: "This service is a messaging platform.",
			Oauth:       true,
			Color:       "#001DDA",
			Icon:        "https://api.iconify.design/mdi:discord.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *discordService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *discordService) GetServiceActionInfo() []schemas.Action {
	return []schemas.Action{}
}

func (service *discordService) GetServiceReactionInfo() []schemas.Reaction {
	service.reactionName = append(service.reactionName, string(schemas.SendMessage))
	defaultValue := schemas.DiscordReactionSendMessageOptions{
		User:   "",
		Message: "",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		fmt.Println("Error marshalling default options:", err)
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.SendMessage),
			Description: "Send a message to a user in Discord.",
			Service:    service.serviceInfo,
			Option:    option,
		},
	}
}

func (service *discordService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	default:
		return nil
	}
}

func (service *discordService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	case string(schemas.SendMessage):
		return service.DiscordReactionSendMessage
	default:
		return nil
	}
}

func (service *discordService) GetActionsName() []string {
	return service.actionName
}

func (service *discordService) GetReactionsName() []string {
	return service.reactionName
}

// Service specific functions

func (service *discordService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("DISCORD_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrDropboxClientIdNotSet
	}

	clientSecret := os.Getenv("DISCORD_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrDropboxSecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/discord"

	apiURL := "https://discord.com/api/v8/oauth2/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
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

	var result schemas.DropboxTokenResponse
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

func (service *discordService) GetUserInfo(
	accessToken string,
) (user schemas.User, err error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://discord.com/api/v8/users/@me",
		nil,
	)
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

	result := schemas.DropboxUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	user = schemas.User{
		Email:    result.Email,
		Username: result.Name.DisplayName,
	}

	return user, nil
}

func (service *discordService) DiscordReactionSendMessage(
	option json.RawMessage,
	idArea uint64,
) string {
	// Parse options
	optionJson := schemas.DiscordReactionSendMessageOptions{}
	err := json.Unmarshal(option, &optionJson)
	if err != nil {
		fmt.Println("Error unmarshalling options:", err)
		time.Sleep(time.Second)
		return "Error unmarshalling options: " + err.Error()
	}

	// Find area
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return "Error finding area:" + err.Error()
	}

	// Find token
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

	// Create the API URL
	apiUrl := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", optionJson.User)

	// Prepare the request body
	body := map[string]string{"content": optionJson.Message}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return "Error marshalling request body:" + err.Error()
	}

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(bodyJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request:" + err.Error()
	}

	// Set headers
	req.Header.Set("Authorization", "Bot "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "Error making request:" + err.Error()
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("Discord API returned an error: %s\n", string(bodyBytes))
		return fmt.Sprintf("Discord API error: %s", resp.Status)
	}

	fmt.Println("Message sent successfully!")
	return "Message sent successfully!"
}
