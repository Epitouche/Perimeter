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

type googleService struct {
	repository        repository.GoogleRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	serviceInfo       schemas.Service
}

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

func (service *googleService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

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

func (service *googleService) GetServiceReactionInfo() []schemas.Reaction {
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
			Description: "Send an email with google service",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

// Service specific functions

func (service *googleService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("GMAIL_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrGoogleClientIdNotSet
	}

	clientSecret := os.Getenv("GMAIL_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrGoogleSecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/google"

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
			time.Sleep(time.Minute)
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

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions

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
