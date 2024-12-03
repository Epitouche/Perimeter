package schemas

import "time"

// GmailTokenResponse represents the response from Gmail when a token is requested.
type GmailTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GmailToken represents the GmailToken entity in the database.
type GmailToken struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	AccessToken string    `                                 json:"access_token"`
	Scope       string    `                                 json:"scope"`
	TokenType   string    `                                 json:"token_type"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}

type GmailUserInfo struct {
	Login     string `json:"login"`
	Id        uint64 `json:"id"         gorm:"primaryKey"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
	HtmlUrl   string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
