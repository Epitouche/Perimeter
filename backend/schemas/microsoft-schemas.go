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

type MicrosoftTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
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
