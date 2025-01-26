package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type GoogleService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Token operations
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	GoogleActionReceiveMail(channel chan string, option json.RawMessage, area schemas.Area)
	// Reactions functions
	GoogleReactionSendMail(option json.RawMessage, area schemas.Area) string
}

// googleService is a struct that encapsulates various repositories and service information
// required for interacting with Google services.
//
// Fields:
// - repository: An instance of GoogleRepository for handling Google-specific data operations.
// - serviceRepository: An instance of ServiceRepository for managing service-related data.
// - areaRepository: An instance of AreaRepository for managing area-related data.
// - tokenRepository: An instance of TokenRepository for handling token-related data.
// - serviceInfo: A Service schema containing information about the service.
type googleService struct {
	repository        repository.GoogleRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	serviceInfo       schemas.Service
}

// NewGoogleService creates a new instance of GoogleService with the provided repositories.
// It initializes the googleService struct with the given repositories and sets the serviceInfo
// with predefined values for the Google service.
//
// Parameters:
//   - repository: an instance of GoogleRepository for accessing Google-related data.
//   - serviceRepository: an instance of ServiceRepository for accessing general service data.
//   - areaRepository: an instance of AreaRepository for accessing area-related data.
//   - tokenRepository: an instance of TokenRepository for accessing token-related data.
//
// Returns:
//
//	A new instance of GoogleService.
func NewGoogleService(
	repository repository.GoogleRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) GoogleService {
	return &googleService{
		repository:        repository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Google,
			Description: "This service is a google service",
			Oauth:       true,
			Color:       "#E60000",
			Icon:        "https://api.iconify.design/mdi:google.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

// GetServiceInfo retrieves the service information for the Google service.
// It returns a schemas.Service object containing the service details.
func (service *googleService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

// FindActionByName returns a function that matches the given action name.
// The returned function takes a channel, a JSON raw message, and an area schema as parameters.
// If the action name matches a predefined action, the corresponding function is returned.
// If the action name does not match any predefined actions, nil is returned.
//
// Parameters:
// - name: The name of the action to find.
//
// Returns:
//   - A function that takes a channel, a JSON raw message, and an area schema as parameters,
//     or nil if the action name does not match any predefined actions.
func (service *googleService) FindActionByName(
	name string,
) func(c chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.ReceiveGoogleMail):
		return service.GoogleActionReceiveMail
	default:
		return nil
	}
}

// FindReactionByName returns a function that corresponds to the given reaction name.
// The returned function takes a JSON raw message and an area schema as parameters and returns a string.
// If the reaction name matches a predefined case, the corresponding function is returned.
// If the reaction name does not match any predefined cases, nil is returned.
//
// Parameters:
//   - name: The name of the reaction to find.
//
// Returns:
//   - A function that takes a JSON raw message and an area schema, and returns a string.
//     If the reaction name does not match any predefined cases, nil is returned.
func (service *googleService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {
	case string(schemas.SendMail):
		return service.GoogleReactionSendMail
	default:
		return nil
	}
}

// GetServiceActionInfo retrieves information about the Google service actions.
// It marshals a default value to JSON and fetches the service information from the repository.
// If any errors occur during these operations, they are printed to the console.
// The function returns a slice of Action schemas containing details about the Google service action.
func (service *googleService) GetServiceActionInfo() []schemas.Action {
	defaultValue := struct{}{}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Google,
	)
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:               string(schemas.ReceiveGoogleMail),
			Description:        "Receive an email with google service",
			Service:            service.serviceInfo,
			Option:             option,
			MinimumRefreshRate: 10,
		},
	}
}

// GetServiceReactionInfo retrieves the reaction information for the Google service.
// It creates a default Gmail reaction option, marshals it to JSON, and fetches the
// service information from the repository. If any errors occur during marshaling or
// fetching the service information, they are printed to the console. The function
// returns a slice of Reaction containing the reaction details.
//
// Returns:
//
//	[]schemas.Reaction: A slice containing the reaction information for the Google service.
func (service *googleService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.GmailReactionSendMailOption{
		To:      "test@example.com",
		Subject: "Test",
		Body:    "a beautiful email",
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
			Description: "Send an email with google service",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

// Service specific functions

// AuthGetServiceAccessToken exchanges an authorization code for an access token
// from Google's OAuth 2.0 server.
//
// It requires the following environment variables to be set:
// - GOOGLE_CLIENT_ID: The client ID obtained from the Google Developer Console.
// - GOOGLE_SECRET: The client secret obtained from the Google Developer Console.
//
// Parameters:
// - code: The authorization code received from the Google authorization server.
//
// Returns:
// - token: A schemas.Token containing the access token, refresh token, and expiration time.
// - err: An error if the token exchange fails or required environment variables are not set.
func (service *googleService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrGoogleClientIdNotSet
	}

	clientSecret := os.Getenv("GOOGLE_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrGoogleSecretNotSet
	}

	redirectURI, err := getRedirectURI(service.serviceInfo.Name)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to get redirect URI because %w", err)
	}

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

	defer resp.Body.Close()

	var result schemas.GoogleTokenResponse
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

	var expiresIn int64
	if result.ExpiresIn > math.MaxInt64 {
		expiresIn = math.MaxInt64
	} else {
		expiresIn = int64(result.ExpiresIn)
	}

	token = schemas.Token{
		Token:        result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpireAt:     time.Now().Add(time.Duration(expiresIn) * time.Second),
	}
	return token, nil
}

// GetUserGmailProfile retrieves the Gmail profile of the authenticated user using the provided access token.
// It sends a GET request to the Gmail API and decodes the response into a GmailProfile schema.
//
// Parameters:
//   - accessToken: A string containing the OAuth 2.0 access token for the authenticated user.
//
// Returns:
//   - result: A GmailProfile struct containing the user's Gmail profile information.
//   - err: An error if the request or decoding fails, otherwise nil.
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

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GmailProfile{}, fmt.Errorf("unable to decode response because %w", err)
	}

	return result, nil
}

// GetUserGoogleProfile retrieves the Google profile of a user using the provided access token.
// It sends a GET request to the Google People API to fetch the user's profile information.
//
// Parameters:
//   - accessToken: A string containing the OAuth 2.0 access token for the user.
//
// Returns:
//   - result: A schemas.GoogleProfile struct containing the user's profile information.
//   - err: An error if the request or decoding fails, otherwise nil.
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

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.GoogleProfile{}, fmt.Errorf("unable to decode response because %w", err)
	}

	return result, nil
}

// GetUserInfo retrieves user information from Google services using the provided access token.
// It fetches the user's Gmail profile and Google profile, and combines the relevant information
// into a schemas.User object.
//
// Parameters:
//   - accessToken: A string representing the OAuth 2.0 access token for accessing Google services.
//
// Returns:
//   - user: A schemas.User object containing the user's email and username.
//   - err: An error object if there was an issue retrieving the user information.
func (service *googleService) GetUserInfo(
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

// initializedGoogleStorageVariable initializes the Google storage variable for a given area.
// It attempts to unmarshal the storage variable from the area's StorageVariable field.
// If unmarshalling fails, it initializes a new GoogleVariableReceiveMail with the current time
// and updates the area's StorageVariable field. The updated area is then saved using the area repository.
//
// Parameters:
//   - area: The area containing the storage variable to be initialized.
//   - service: The googleService instance used to update the area repository.
//
// Returns:
//   - schemas.GoogleVariableReceiveMail: The initialized GoogleVariableReceiveMail.
//   - error: An error if any operation fails during the initialization process.
func initializedGoogleStorageVariable(
	area schemas.Area,
	service googleService,
) (schemas.GoogleVariableReceiveMail, error) {
	variable := schemas.GoogleVariableReceiveMail{}
	err := json.Unmarshal(area.StorageVariable, &variable)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return variable, err
		} else {
			println("initializing storage variable")
			variable = schemas.GoogleVariableReceiveMail{
				Time: time.Now(),
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
		variable = schemas.GoogleVariableReceiveMail{
			Time: time.Now(),
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

// getLastEmailId retrieves the most recent email ID from the user's Gmail inbox after a specified time.
// It takes a token for authorization and a GoogleVariableReceiveMail struct containing the time query.
// It returns a GmailEmailResponse struct and an error if any occurred during the process.
//
// Parameters:
//   - token: schemas.Token containing the authorization token.
//   - variable: schemas.GoogleVariableReceiveMail containing the time query.
//
// Returns:
//   - schemas.GmailEmailResponse: Struct containing the email response.
//   - error: Error if any occurred during the process.
func getLastEmailId(
	token schemas.Token,
	variable schemas.GoogleVariableReceiveMail,
) (schemas.GmailEmailResponse, error) {
	emailResponse := schemas.GmailEmailResponse{}
	timeQuery := variable.Time.Format("2006/01/02")
	apiURL := fmt.Sprintf(
		"https://gmail.googleapis.com/gmail/v1/users/me/messages?maxResults=1&labelIds=INBOX&q=after:%s",
		timeQuery,
	)

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
		return emailResponse, err
	}

	err = json.NewDecoder(resp.Body).Decode(&emailResponse)
	if err != nil {
		println("error decoding response: " + err.Error())
		return emailResponse, err
	}
	return emailResponse, nil
}

// getLastEmailDetails retrieves the details of the last email for a given message ID.
// It makes a request to the Gmail API to fetch the email headers and extracts the Date, From, and Subject fields.
//
// Parameters:
//   - id: The ID of the email message to retrieve.
//   - token: The authentication token required to access the Gmail API.
//
// Returns:
//   - schemas.EmailDetails: A struct containing the Date, From, and Subject of the email.
//   - error: An error if the request fails or if the email details cannot be found.
func getLastEmailDetails(id string, token schemas.Token) (schemas.EmailDetails, error) {
	var emailDetails schemas.EmailDetails
	ctx := context.Background()
	client := &http.Client{}

	apiURL := fmt.Sprintf(
		"https://gmail.googleapis.com/gmail/v1/users/me/messages/%s?fields=payload(headers),id",
		id,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		println("error creating request: " + err.Error())
		return emailDetails, err
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)

	resp, err := client.Do(req)
	if err != nil {
		println("error making request: " + err.Error())
		return emailDetails, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("error status code: " + fmt.Sprint(resp.StatusCode))
		return emailDetails, err
	}
	var emailAllDetails schemas.GmailMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&emailAllDetails)
	if err != nil {
		println("error decoding response: " + err.Error())
		return emailDetails, err
	}

	for _, header := range emailAllDetails.Payload.Headers {
		switch header.Name {
		case "Date":
			emailDetails.Date = header.Value
		case "From":
			emailDetails.From = header.Value
		case "Subject":
			emailDetails.Subject = header.Value
		}
	}

	if emailDetails.Date == "" || emailDetails.From == "" || emailDetails.Subject == "" {
		println("error: email details not found")
		return emailDetails, err
	}

	return emailDetails, nil
}

// Actions functions

// GoogleActionReceiveMail handles the process of receiving emails from a Google account.
// It initializes the Google storage variable, retrieves the token, gets the last email ID,
// fetches the email details, and checks if there are any new emails. If a new email is found,
// it updates the storage variable and sends a response through the provided channel.
//
// Parameters:
//   - channel: A channel to send the response string when a new email is received.
//   - option: A JSON raw message containing additional options (currently unused).
//   - area: A schemas.Area object containing user and action details.
//
// The function performs the following steps:
//  1. Initializes the Google storage variable.
//  2. Retrieves the token for the user and service.
//  3. Gets the last email ID.
//  4. Fetches the email details.
//  5. Parses the email time and checks if it is a new email.
//  6. Updates the storage variable and sends a response if a new email is found.
//  7. Sleeps for the specified refresh rate before returning.
func (service *googleService) GoogleActionReceiveMail(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	variable, err := initializedGoogleStorageVariable(area, *service)
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

	emailResponse, err := getLastEmailId(token, variable)
	if err != nil {
		println("error getting last email id: " + err.Error())
		return
	}

	if len(emailResponse.Messages) > 0 {
		id := emailResponse.Messages[0].Id
		emailDetails, err := getLastEmailDetails(id, token)
		if err != nil {
			println("error getting last email details: " + err.Error())
			return
		}

		emailTime, err := time.Parse(time.RFC1123Z, emailDetails.Date)
		if err != nil {
			emailTime, err = time.Parse(time.RFC1123, emailDetails.Date)
			if err != nil {
				println("error parsing time: " + err.Error())
				return
			}
		}
		if variable.Time.After(emailTime) {
			println("no new emails")
			WaitAction(area)
			return
		}
		response := fmt.Sprintf("New email received from %s: object: %s",
			emailDetails.From,
			emailDetails.Subject,
		)
		variable.Time, err = time.Parse(time.RFC1123Z, emailDetails.Date)
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

	WaitAction(area)
}

// Reactions functions

// GoogleReactionSendMail sends an email using the Gmail API based on the provided options and area information.
// It unmarshals the provided JSON options into a GmailReactionSendMailOption struct, retrieves the user's token,
// constructs the email message, and sends it via the Gmail API.
//
// Parameters:
//   - option: A JSON raw message containing the email options (To, Subject, Body).
//   - area: An Area struct containing user and reaction information.
//
// Returns:
//
//	A string indicating the result of the operation, either an error message or a success message.
func (service *googleService) GoogleReactionSendMail(
	option json.RawMessage,
	area schemas.Area,
) string {
	optionJSON := schemas.GmailReactionSendMailOption{}

	err := json.Unmarshal(option, &optionJSON)
	if err != nil {
		println("error unmarshal gmail option: " + err.Error())
		time.Sleep(time.Second)
		return "Error unmarshal gmail option" + err.Error()
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
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return "Error reading response body:" + err.Error()
		}
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
