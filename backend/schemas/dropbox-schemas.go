package schemas

import "errors"

type DropboxAction string

type DropboxReaction string

const ()

// DropboxTokenResponse represents the response from Dropbox when a token is requested.
type DropboxTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type DropboxUserInfo struct {
	AccountId   string `json:"account_id"`
	AccountType struct {
		Tag string `json:".tag"`
	} `json:"account_type"`
	Country       string `json:"country"`
	Disabled      bool   `json:"disabled"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	IsPaired      bool   `json:"is_paired"`
	Locale        string `json:"locale"`
	Name          struct {
		AbbreviatedName string `json:"abbreviated_name"`
		DisplayName     string `json:"display_name"`
		FamiliarName    string `json:"familiar_name"`
		GivenName       string `json:"given_name"`
		Surname         string `json:"surname"`
	} `json:"name"`
	ReferralLink string `json:"referral_link"`
	RootInfo     struct {
		Tag             string `json:".tag"`
		HomeNamespaceId string `json:"home_namespace_id"`
		RootNamespaceId string `json:"root_namespace_id"`
	} `json:"root_info"`
}

type DropboxReactionSendMailOption struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// Errors Messages.
var (
	ErrDropboxSecretNotSet   = errors.New("DROPBOX_SECRET is not set")
	ErrDropboxClientIdNotSet = errors.New("DROPBOX_CLIENT_ID is not set")
)

type DropboxMobileTokenRequest struct {
	Token string `json:"token"`
}
