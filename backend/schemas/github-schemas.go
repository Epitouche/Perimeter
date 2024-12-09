package schemas

import "errors"

// GitHubTokenResponse represents the response from Github when a token is requested.
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type GithubUserInfo struct {
	Login     string `json:"login"`
	Id        uint64 `json:"id"         gorm:"primaryKey"`
	AvatarURL string `json:"avatar_url"`
	Type      string `json:"type"`
	HtmlURL   string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

// Errors Messages.
var (
	ErrGithubSecretNotSet   = errors.New("GITHUB_SECRET is not set")
	ErrGithubClientIdNotSet = errors.New("GITHUB_CLIENT_ID is not set")
)
