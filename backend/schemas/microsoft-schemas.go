package schemas

import (
	"errors"
	"time"
)

type MicrosoftAction string // MicrosoftAction is a string type to represent the action to be performed.

const (
	EventStarting        MicrosoftAction = "EventStarting"        // EventStarting is the action to start an event.
	ReceiveMicrosoftMail MicrosoftAction = "ReceiveMicrosoftMail" // ReceiveMicrosoftMail is the action to receive emails from Microsoft.
)

type MicrosoftReaction string // MicrosoftReaction is a string type to represent the reaction to be performed.

const (
	SendMicrosoftMail MicrosoftReaction = "SendMicrosoftMail" // SendMicrosoftMail is the reaction to send an email using Microsoft.
	CreateEvent       MicrosoftReaction = "createEvent"       // CreateEvent is the reaction to create an event using Microsoft.
)

type MicrosoftEventIncomingOptions struct {
	Name string `json:"name"` // The name of the event
}

type MicrosoftReactionSendMailOptions struct {
	Subject   string `json:"subject"`   // The subject of the email
	Body      string `json:"body"`      // The body of the email
	Recipient string `json:"recipient"` // The recipient of the email
}

type MicrosoftCreateEventOptions struct {
	Subject  string `json:"subject"`  // The subject of the event
	Body     string `json:"body"`     // The body of the event
	Location string `json:"location"` // The location of the event
	Start    string `json:"start"`    // The start time of the event
	End      string `json:"end"`      // The end time of the event
}

type MicrosoftTokenResponse struct {
	AccessToken  string `json:"access_token"`  // The access token
	ExpiresIn    uint64 `json:"expires_in"`    // The expiration time of the token in seconds
	Scope        string `json:"scope"`         // The scope of the token
	TokenType    string `json:"token_type"`    // The type of the token
	RefreshToken string `json:"refresh_token"` // The refresh token
}

type MicrosoftVariableTime struct {
	Time time.Time `json:"time"`
}

// error messages
var (
	ErrMicrosoftClientIdNotSet = errors.New("MICROSOFT_CLIENT_ID is not set")
	ErrMicrosoftSecretNotSet   = errors.New("MICROSOFT_SECRET is not set")
)

type MicrosoftUserInfo struct {
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
	DisplayName       string `json:"displayName"`
}

type MicrosoftEmailResponse struct {
	Value []struct {
		ID      string `json:"id"`
		Subject string `json:"subject"`
		From    struct {
			EmailAddress struct {
				Address string `json:"address"`
			} `json:"emailAddress"`
		} `json:"from"`
		ReceivedDateTime string `json:"receivedDateTime"`
	} `json:"value"`
}

type MicrosoftEventListResponse struct {
	Value []struct {
		Subject string `json:"subject"`
		Start   struct {
			DateTime string `json:"dateTime"`
		} `json:"start"`
		End struct {
			DateTime string `json:"dateTime"`
		} `json:"end"`
	} `json:"value"`
}
