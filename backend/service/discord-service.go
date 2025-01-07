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
		User:    "",
		Message: "",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		fmt.Println("Error marshalling default options:", err)
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Discord,
	)
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.SendMessage),
			Description: "Send a message to a user in Discord.",
			Service:     service.serviceInfo,
			Option:      option,
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
	// Parse the options
	optionJson := schemas.DiscordReactionSendMessageOptions{}
	err := json.Unmarshal(option, &optionJson)
	if err != nil {
		fmt.Println("Error unmarshalling options:", err)
		return "Error unmarshalling options: " + err.Error()
	}

	// Retrieve the area
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return "Error finding area: " + err.Error()
	}

	// Retrieve the token
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return "Error finding token: " + err.Error()
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return "Error: Token not found"
	}

	// Step 1: Create a DM channel
	apiUrl := "https://discord.com/api/v10/users/@me/channels"
	body := map[string]string{"recipient_id": optionJson.User}
	bodyJson, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(bodyJson))
	if err != nil {
		fmt.Println("Error creating DM channel request:", err)
		return "Error creating DM channel request: " + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making DM channel request:", err)
		return "Error making DM channel request: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Error creating DM channel:", string(bodyBytes))
		return "Error creating DM channel: " + string(bodyBytes)
	}

	var dmResponse struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dmResponse); err != nil {
		fmt.Println("Error decoding DM channel response:", err)
		return "Error decoding DM channel response: " + err.Error()
	}

	// Step 2: Send the message to the DM channel
	messageUrl := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", dmResponse.ID)
	messageBody := map[string]string{"content": optionJson.Message}
	messageJson, _ := json.Marshal(messageBody)

	msgReq, err := http.NewRequest(http.MethodPost, messageUrl, bytes.NewBuffer(messageJson))
	if err != nil {
		fmt.Println("Error creating message request:", err)
		return "Error creating message request: " + err.Error()
	}

	msgReq.Header.Set("Authorization", "Bearer "+token.Token)
	msgReq.Header.Set("Content-Type", "application/json")

	msgResp, err := client.Do(msgReq)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return "Error sending message: " + err.Error()
	}
	defer msgResp.Body.Close()

	if msgResp.StatusCode != http.StatusOK && msgResp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(msgResp.Body)
		fmt.Println("Error sending message:", string(bodyBytes))
		return "Error sending message: " + string(bodyBytes)
	}

	return "Message sent successfully!"
}
