package schemas

import (
	"errors"
	"time"
)

type ServiceName string

const (
	Spotify        ServiceName = "spotify"
	Openweathermap ServiceName = "openweathermap"
	Timer          ServiceName = "timer"
	Gmail          ServiceName = "gmail"
	Github         ServiceName = "github"
	Dropbox        ServiceName = "dropbox"
)

type ServiceJSON struct {
	Name     ServiceName    `json:"name"`      // Name of the service
	Action   []ActionJSON   `json:"actions"`   // List of actions for the service
	Reaction []ReactionJSON `json:"reactions"` // List of reactions for the service
}

// GithubToken represents the GithubToken entity in the database.
type Service struct {
	Id          uint64      `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Name        ServiceName `                                 json:"name"         binding:"required"`
	Description string      `                                 json:"description"  binding:"required"`
	Oauth       bool        `                                 json:"oauth"        binding:"required"`
	Icon        string      `                                 json:"icon"         binding:"required"`
	Color       string      `                                 json:"color"        binding:"required"`
	CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

// Errors Messages.
var (
	ErrNotOauthService = errors.New("service is not an oauth service")
)

type MobileTokenRequest struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresIn    time.Time `json:"accessTokenExpirationDate"`
}
