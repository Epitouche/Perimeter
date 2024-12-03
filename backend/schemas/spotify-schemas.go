package schemas

import "time"

// SpotifyTokenResponse represents the response from Github when a token is requested.
type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// SpotifyToken represents the SpotifyToken entity in the database.
type SpotifyToken struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	AccessToken string    `                                 json:"access_token"`
	Scope       string    `                                 json:"scope"`
	TokenType   string    `                                 json:"token_type"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}

type SpotifyUserInfo struct {
	Login     string `json:"login"`
	Id        uint64 `json:"id"         gorm:"primaryKey"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
	HtmlUrl   string `json:"html_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
