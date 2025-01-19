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
	"strings"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type MicrosoftService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	MicrosoftActionReceiveMail(
		channel chan string,
		option json.RawMessage,
		area schemas.Area,
	)
	MicrosoftActionEventStarting(
		channel chan string,
		option json.RawMessage,
		area schemas.Area,
	)
	// Reactions functions
	MicrosoftReactionSendMail(
		option json.RawMessage,
		area schemas.Area,
	) string
	MicrosoftReactionCreateEvent(
		option json.RawMessage,
		area schemas.Area,
	) string
}

// microsoftService is a struct that encapsulates various repositories and service information
// required for Microsoft-related operations.
//
// Fields:
// - repository: MicrosoftRepository interface for interacting with Microsoft data.
// - serviceRepository: ServiceRepository interface for interacting with service data.
// - areaRepository: AreaRepository interface for interacting with area data.
// - tokenRepository: TokenRepository interface for managing tokens.
// - serviceInfo: Service schema containing information about the service.
type microsoftService struct {
	repository        repository.MicrosoftRepository // Microsoft repository
	serviceRepository repository.ServiceRepository   // Service repository
	areaRepository    repository.AreaRepository      // Area repository
	tokenRepository   repository.TokenRepository     // Token repository
	serviceInfo       schemas.Service                // Service information
}

// NewMicrosoftService creates a new instance of MicrosoftService with the provided repositories.
// It initializes the service with predefined information such as name, description, OAuth support, color, and icon.
//
// Parameters:
//   - githubTokenRepository: repository.MicrosoftRepository - Repository for handling Microsoft tokens.
//   - serviceRepository: repository.ServiceRepository - Repository for handling service-related operations.
//   - areaRepository: repository.AreaRepository - Repository for handling area-related operations.
//   - tokenRepository: repository.TokenRepository - Repository for handling token-related operations.
//
// Returns:
//   - MicrosoftService: A new instance of MicrosoftService.
func NewMicrosoftService(
	githubTokenRepository repository.MicrosoftRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) MicrosoftService {
	return &microsoftService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Microsoft,
			Description: "This service is used to interact with Microsoft services",
			Oauth:       true,
			Color:       "#001DDA",
			Icon:        "https://api.iconify.design/mdi:microsoft.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

// GetServiceInfo retrieves the service information for the Microsoft service.
// It returns a schemas.Service object containing the service details.
func (service *microsoftService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

// GetServiceActionInfo retrieves information about available actions for the Microsoft service.
// It returns a slice of schemas.Action, each representing a specific action that can be performed
// using Microsoft services. The function initializes default options and event incoming options,
// marshals them into JSON format, and assigns them to the respective actions. If any errors occur
// during the marshalling or service lookup process, they are logged to the console.
func (service *microsoftService) GetServiceActionInfo() []schemas.Action {
	defaultValue := struct{}{}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		fmt.Println("Error marshalling default options:", err)
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Microsoft,
	)
	if err != nil {
		println("error find service by name: " + err.Error())
	}

	defaultValueEventIncoming := schemas.MicrosoftEventIncomingOptions{
		Name: "Meeting with the boss",
	}
	optionEventIncoming, err := json.Marshal(defaultValueEventIncoming)
	if err != nil {
		fmt.Println("Error marshalling default options:", err)
	}
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:               string(schemas.ReceiveMicrosoftMail),
			Description:        "Receive a mail using Microsoft services",
			Service:            service.serviceInfo,
			Option:             option,
			MinimumRefreshRate: 10,
		},
		{
			Name:        string(schemas.EventStarting),
			Description: "Event starting using Microsoft services",
			Service:     service.serviceInfo,
			Option:      optionEventIncoming,
		},
	}
}

// GetServiceReactionInfo retrieves the reaction information for Microsoft services.
// It returns a slice of schemas.Reaction which includes details for sending an email
// and creating an event using Microsoft services.
//
// The function initializes default options for sending an email and creating an event,
// marshals these options into JSON format, and retrieves the service information from
// the service repository. If any errors occur during marshalling or retrieving the service
// information, they are printed to the console.
//
// Returns:
//
//	[]schemas.Reaction: A slice containing reaction information for sending an email
//	and creating an event using Microsoft services.
func (service *microsoftService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.MicrosoftReactionSendMailOptions{
		Subject:   "newsletter",
		Body:      "a beautiful email",
		Recipient: "test@example.com",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		fmt.Println("Error marshalling default options:", err)
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Microsoft,
	)
	if err != nil {
		println("error find service by name: " + err.Error())
	}

	defaultValueCreateEvent := schemas.MicrosoftCreateEventOptions{
		Subject:  "Meeting",
		Body:     "Weekly meeting",
		Location: "Bordeaux",
		Start:    "2025-01-12T16:06:00",
		End:      "YYYY-MM-DDTHH:MM:SS",
	}
	optionCreateEvent, err := json.Marshal(defaultValueCreateEvent)
	if err != nil {
		fmt.Println("Error marshalling default options:", err)
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.SendMicrosoftMail),
			Description: "Send a mail using Microsoft services",
			Service:     service.serviceInfo,
			Option:      option,
		},
		{
			Name:        string(schemas.CreateEvent),
			Description: "Create an event using Microsoft services",
			Service:     service.serviceInfo,
			Option:      optionCreateEvent,
		},
	}
}

// FindActionByName returns a function that matches the provided action name.
// The returned function takes a channel, a JSON raw message, and an area schema as parameters.
//
// Parameters:
//   - name: The name of the action to find.
//
// Returns:
//   - A function that matches the provided action name, or nil if no match is found.
func (service *microsoftService) FindActionByName(
	name string,
) func(c chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.ReceiveMicrosoftMail):
		return service.MicrosoftActionReceiveMail
	case string(schemas.EventStarting):
		return service.MicrosoftActionEventStarting
	default:
		return nil
	}
}

// FindReactionByName returns a function that performs a specific Microsoft service reaction
// based on the provided name. The returned function takes a JSON raw message and an area schema
// as parameters and returns a string.
//
// Parameters:
//   - name: The name of the reaction to find.
//
// Returns:
//   - A function that takes a JSON raw message and an area schema, and returns a string.
//     If the name does not match any known reactions, it returns nil.
func (service *microsoftService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {
	case string(schemas.SendMicrosoftMail):
		return service.MicrosoftReactionSendMail
	case string(schemas.CreateEvent):
		return service.MicrosoftReactionCreateEvent
	default:
		return nil
	}
}

// AuthGetServiceAccessToken exchanges an authorization code for an access token from Microsoft's OAuth2 service.
// It takes an authorization code as input and returns a Token schema or an error if the process fails.
//
// The function performs the following steps:
// 1. Retrieves the Microsoft client ID from environment variables.
// 2. Gets the redirect URI based on the service name.
// 3. Constructs the request to the Microsoft OAuth2 token endpoint.
// 4. Sends the request and reads the response.
// 5. Parses the response and extracts the access token, refresh token, and expiration time.
//
// Parameters:
// - code: The authorization code received from Microsoft's OAuth2 authorization endpoint.
//
// Returns:
// - token: A Token schema containing the access token, refresh token, and expiration time.
// - err: An error if the process fails at any step.
func (service *microsoftService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("MICROSOFT_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrMicrosoftClientIdNotSet
	}

	redirectURI, err := getRedirectURI(service.serviceInfo.Name)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to get redirect URI because %w", err)
	}

	apiURL := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to create request because %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to make request because %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("response body: ", string(bodyBytes))

	var result schemas.MicrosoftTokenResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to decode response because %w", err)
	}

	if result.AccessToken == "" || result.TokenType == "" {
		return schemas.Token{}, schemas.ErrAccessTokenNotFoundInResponse
	}

	token = schemas.Token{
		Token:        result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpireAt:     time.Now().Add(time.Duration(result.ExpiresIn) * time.Second),
	}
	return token, nil
}

// GetUserInfo retrieves the user information from Microsoft Graph API using the provided access token.
// It sends a GET request to the "https://graph.microsoft.com/v1.0/me" endpoint and decodes the response
// into a schemas.User object.
//
// Parameters:
//   - accessToken: A string representing the OAuth 2.0 access token for authenticating the request.
//
// Returns:
//   - user: A schemas.User object containing the user's email and username.
//   - err: An error object if there was an issue creating the request, making the request, or decoding the response.
//
// The function returns an error if:
//   - The request could not be created.
//   - The request could not be made.
//   - The response status code is not 200 OK.
//   - The response body could not be decoded.
func (service *microsoftService) GetUserInfo(
	accessToken string,
) (user schemas.User, err error) {
	ctx := context.Background()

	url := "https://graph.microsoft.com/v1.0/me"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to create request because %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to make request because %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return schemas.User{}, fmt.Errorf("failed to fetch user info: %s", resp.Status)
	}

	var result schemas.MicrosoftUserInfo
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	user = schemas.User{
		Email:    result.Mail,
		Username: result.DisplayName,
	}
	if user.Email == "" {
		user.Email = result.UserPrincipalName
	}

	return user, nil
}

// Actions functions

// MicrosoftActionEventStarting handles the event of a Microsoft action starting.
// It retrieves the user's Microsoft events, checks if any event matches the specified options,
// and updates the area storage variable if a matching event is found.
//
// Parameters:
//   - channel: A channel to send messages about the event status.
//   - option: A JSON raw message containing the options for the Microsoft event.
//   - area: The area schema containing user and action information.
//
// The function performs the following steps:
//  1. Unmarshals the options from the JSON raw message.
//  2. Initializes the Microsoft storage variable.
//  3. Retrieves the user's token for the Microsoft service.
//  4. Makes an HTTP GET request to the Microsoft Graph API to fetch the user's events.
//  5. Checks if any event matches the specified options.
//  6. Updates the area storage variable and sends a message to the channel if a matching event is found.
//  7. Sleeps for 10 seconds before returning.
//
// If any error occurs during these steps, it prints an error message and returns.
func (service *microsoftService) MicrosoftActionEventStarting(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	options := schemas.MicrosoftEventIncomingOptions{}
	err := json.Unmarshal(option, &options)
	if err != nil {
		println("error unmarshalling options: " + err.Error())
		return
	}

	variable, err := initializedMicrosoftStorageVariable(area, *service)
	if err != nil {
		println("error initializing storage variable: " + err.Error())
		return
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Action.ServiceId,
	)
	if err != nil {
		println("error retrieving token: " + err.Error())
		return
	}

	apiURL := "https://graph.microsoft.com/v1.0/me/events?$select=subject,start,end"

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		println("error creating request: " + err.Error())
		return
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("error making request: " + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("error status code: " + fmt.Sprint(resp.StatusCode))
		return
	}

	var response schemas.MicrosoftEventListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		println("error decoding response: " + err.Error())
		return
	}

	for _, event := range response.Value {
		if event.Subject == options.Name {
			eventStartTime, err := time.Parse("2006-01-02T15:04:05.0000000", event.Start.DateTime)
			if err != nil {
				println("error parsing event start time: " + err.Error())
				continue
			}
			eventEndTime, err := time.Parse("2006-01-02T15:04:05.0000000", event.End.DateTime)
			if err != nil {
				println("error parsing event end time: " + err.Error())
				continue
			}

			if variable.Time.After(eventStartTime) && variable.Time.Before(eventEndTime) {
				variable.Time = eventEndTime.Add(time.Second)
				area.StorageVariable, err = json.Marshal(variable)
				if err != nil {
					println("error marshalling storage variable: " + err.Error())
					return
				}
				err = service.areaRepository.Update(area)
				if err != nil {
					println("error updating area: " + err.Error())
					return
				}
				channel <- fmt.Sprintf("Event '%s' is starting at %s", event.Subject, event.Start.DateTime)
				time.Sleep(time.Second * 10)
				return
			}
		}
	}

	println("no matching events found")
	time.Sleep(time.Second * 10)
}

// initializedMicrosoftStorageVariable initializes the Microsoft storage variable for a given area.
// It attempts to unmarshal the storage variable from the area's StorageVariable field.
// If unmarshalling fails, it initializes the storage variable with the current UTC time and updates the area in the repository.
// If the storage variable's time is zero, it also initializes it with the current UTC time and updates the area in the repository.
//
// Parameters:
//   - area: The area containing the storage variable to be initialized.
//   - service: The microsoftService instance used to update the area in the repository.
//
// Returns:
//   - schemas.MicrosoftVariableTime: The initialized Microsoft storage variable.
//   - error: An error if any occurred during the process.
func initializedMicrosoftStorageVariable(
	area schemas.Area,
	service microsoftService,
) (schemas.MicrosoftVariableTime, error) {
	variable := schemas.MicrosoftVariableTime{}
	err := json.Unmarshal(area.StorageVariable, &variable)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return variable, err
		} else {
			println("initializing storage variable")
			variable = schemas.MicrosoftVariableTime{
				Time: time.Now().UTC(),
			}
			area.StorageVariable, err = json.Marshal(variable)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return variable, err
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return variable, err
			}
		}
	}

	if variable.Time.IsZero() {
		variable = schemas.MicrosoftVariableTime{
			Time: time.Now().UTC(),
		}
		area.StorageVariable, err = json.Marshal(variable)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return variable, err
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return variable, err
		}
	}
	return variable, nil
}

func getNewEmails(
	token schemas.Token,
	variable schemas.MicrosoftVariableTime,
) (schemas.MicrosoftEmailResponse, error) {
	var emailResponse schemas.MicrosoftEmailResponse
	apiURL := "https://graph.microsoft.com/v1.0/me/messages?$filter=receivedDateTime+gt+" + variable.Time.Format(
		"2006-01-02T15:04:05",
	) + "Z"
	println("apiURL: " + apiURL)
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		println("error creating request: " + err.Error())
		return emailResponse, err
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("error making request: " + err.Error())
		return emailResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("error status code: " + fmt.Sprint(resp.StatusCode))
		return emailResponse, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&emailResponse)
	if err != nil {
		println("error decoding response: " + err.Error())
		return emailResponse, err
	}

	return emailResponse, nil
}

// Actions functions

// MicrosoftActionReceiveMail handles the action of receiving emails from Microsoft service.
// It initializes the storage variable, retrieves the token, fetches new emails, and updates the area repository.
//
// Parameters:
//   - channel: A channel to send the response string.
//   - option: A JSON raw message containing options.
//   - area: The area schema containing user and action details.
//
// The function performs the following steps:
//  1. Initializes the storage variable using the provided area and service.
//  2. Retrieves the token associated with the user and service.
//  3. Fetches new emails using the token and storage variable.
//  4. If new emails are found, it processes the latest email, updates the storage variable, and sends a response through the channel.
//  5. Updates the area repository with the new storage variable.
//  6. Sleeps for a duration based on the action's refresh rate.
func (service *microsoftService) MicrosoftActionReceiveMail(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	variable, err := initializedMicrosoftStorageVariable(area, *service)
	if err != nil {
		println("error initializing storage variable: " + err.Error())
		return
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Action.ServiceId,
	)
	if err != nil || token.Token == "" {
		println("error retrieving token or token not found")
		return
	}

	emailResponse, err := getNewEmails(token, variable)
	if err != nil {
		println("error getting new emails: " + err.Error())
		return
	}

	if len(emailResponse.Value) > 0 {
		latestEmail := emailResponse.Value[0]
		response := fmt.Sprintf("New email received from %s: object: %s",
			latestEmail.From.EmailAddress.Address,
			latestEmail.Subject,
		)
		println(response)
		variable.Time, err = time.Parse(time.RFC3339, latestEmail.ReceivedDateTime)
		if err != nil {
			println("error parsing time: " + err.Error())
			return
		}
		variable.Time = variable.Time.Add(time.Second)
		area.StorageVariable, err = json.Marshal(variable)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
		channel <- response
	} else {
		println("No new emails")
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions

// MicrosoftReactionSendMail sends an email using the Microsoft Graph API.
//
// Parameters:
//   - option: A JSON raw message containing the email options.
//   - area: A schemas.Area object containing user and service information.
//
// Returns:
//
//	A string indicating the result of the email sending operation.
//
// The function performs the following steps:
//  1. Unmarshals the email options from the provided JSON raw message.
//  2. Retrieves the user's token for the Microsoft service from the token repository.
//  3. Constructs the email payload and marshals it into JSON.
//  4. Creates an HTTP POST request to the Microsoft Graph API to send the email.
//  5. Sets the necessary headers, including the authorization token.
//  6. Sends the HTTP request and checks the response status code.
//  7. Returns a success message if the email is sent successfully, or an error message if any step fails.
func (service *microsoftService) MicrosoftReactionSendMail(
	option json.RawMessage,
	area schemas.Area,
) string {
	options := schemas.MicrosoftReactionSendMailOptions{}
	err := json.Unmarshal(option, &options)
	if err != nil {
		fmt.Println("Error unmarshalling options:", err)
		return "Error unmarshalling options: " + err.Error()
	}

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

	apiURL := "https://graph.microsoft.com/v1.0/me/sendMail"

	payload := map[string]interface{}{
		"message": map[string]interface{}{
			"subject": options.Subject,
			"body": map[string]string{
				"contentType": "Text",
				"content":     options.Body,
			},
			"toRecipients": []map[string]map[string]string{
				{
					"emailAddress": {
						"address": options.Recipient,
					},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling email payload:", err)
		return "Error marshalling email payload: " + err.Error()
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "Error creating HTTP request: " + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending email request:", err)
		return "Error sending email request: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Error sending email:", string(bodyBytes))
		return "Error sending email: " + string(bodyBytes)
	}

	return "Email sent successfully!"
}

// MicrosoftReactionCreateEvent creates a new event in the Microsoft calendar for the specified user.
// It takes a JSON raw message containing event options and an Area schema as input parameters.
// The function returns a string indicating the result of the operation.
//
// Parameters:
//   - option: A JSON raw message containing the event options.
//   - area: An Area schema containing user and service information.
//
// Returns:
//   - A string indicating the result of the event creation operation.
//
// The function performs the following steps:
//  1. Unmarshals the JSON raw message into MicrosoftCreateEventOptions.
//  2. Retrieves the user's token for the Microsoft service from the token repository.
//  3. Validates the retrieved token.
//  4. Constructs the event payload with the provided options.
//  5. Sends an HTTP POST request to the Microsoft Graph API to create the event.
//  6. Handles the response and returns the appropriate result message.
func (service *microsoftService) MicrosoftReactionCreateEvent(
	option json.RawMessage,
	area schemas.Area,
) string {
	options := schemas.MicrosoftCreateEventOptions{}
	err := json.Unmarshal(option, &options)
	if err != nil {
		fmt.Println("Error unmarshalling options:", err)
		return "Error unmarshalling options: " + err.Error()
	}

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

	apiURL := "https://graph.microsoft.com/v1.0/me/events"

	startTime, err := time.Parse("2006-01-02T15:04:05", options.Start)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return "Error parsing start time: " + err.Error()
	}

	startTime = startTime.Add(-time.Hour)

	endTime, err := time.Parse("2006-01-02T15:04:05", options.End)
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return "Error parsing end time: " + err.Error()
	}

	endTime = endTime.Add(-time.Hour)

	payload := map[string]interface{}{
		"subject":  options.Subject,
		"body":     map[string]string{"contentType": "Text", "content": options.Body},
		"location": map[string]string{"displayName": options.Location},
		"start": map[string]string{
			"dateTime": startTime.Format(time.RFC3339),
			"timeZone": "UTC",
		},
		"end": map[string]string{"dateTime": endTime.Format(time.RFC3339), "timeZone": "UTC"},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling event payload:", err)
		return "Error marshalling event payload: " + err.Error()
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "Error creating HTTP request: " + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating event request:", err)
		return "Error creating event request: " + err.Error()
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Error creating event:", string(bodyBytes))
		return "Error creating event: " + string(bodyBytes)
	}

	return "Event created successfully!"
}
