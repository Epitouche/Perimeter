package schemas

import "errors"

type MicrosoftAction string

type MicrosoftReaction string

const (
	SendMicrosoftMail MicrosoftReaction = "SendMicrosoftMail"
)

type MicrosoftReactionSendMailOptions struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

// error messages
var (
	ErrMicrosoftClientIdNotSet = errors.New("DISCORD_CLIENT_ID is not set")
	ErrMicrosoftSecretNotSet   = errors.New("DISCORD_SECRET is not set")
)

type MicrosoftUserInfo struct {
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
	DisplayName       string `json:"displayName"`
}
