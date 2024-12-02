package schemas

import "time"

// GitHubTokenResponse represents the response from Github when a token is requested.
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GithubToken represents the GithubToken entity in the database.
type GithubToken struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	AccessToken string    `                                 json:"access_token"`
	Scope       string    `                                 json:"scope"`
	TokenType   string    `                                 json:"token_type"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}

type GithubUserInfo struct {
	Login     string `json:"login"`
	Id        uint64 `json:"id"         gorm:"primaryKey"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
	HtmlUrl   string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
