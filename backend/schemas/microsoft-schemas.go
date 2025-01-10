package schemas

import (
	"errors"
	"time"
)

type MicrosoftAction string

const (
	ReceiveMicrosoftMail MicrosoftAction = "ReceiveMicrosoftMail"
)

type MicrosoftReaction string

const (
	SendMicrosoftMail MicrosoftReaction = "SendMicrosoftMail"
)

type MicrosoftReactionSendMailOptions struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

type MicrosoftTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type MicrosoftVariableReceiveMail struct {
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
