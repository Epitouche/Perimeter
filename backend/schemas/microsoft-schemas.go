package schemas

import (
	"errors"
	"time"
)

type MicrosoftAction string

const (
	EventStarting        MicrosoftAction = "EventStarting"
	ReceiveMicrosoftMail MicrosoftAction = "ReceiveMicrosoftMail"
)

type MicrosoftReaction string

const (
	SendMicrosoftMail MicrosoftReaction = "SendMicrosoftMail"
	CreateEvent       MicrosoftReaction = "createEvent"
)

type MicrosoftEventIncomingOptions struct {
	Name string `json:"name"`
}

type MicrosoftReactionSendMailOptions struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

type MicrosoftCreateEventOptions struct {
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Location    string `json:"location"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

type MicrosoftTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
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
