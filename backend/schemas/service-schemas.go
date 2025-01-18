package schemas

import (
	"errors"
	"time"
)

type ServiceName string

const (
	Spotify        ServiceName = "Spotify"        // Spotify is the music streaming service.
	OpenWeatherMap ServiceName = "OpenWeatherMap" // OpenWeatherMap is a weather service.
	Timer          ServiceName = "Timer"          // Timer is a service for setting timers.
	Google         ServiceName = "Google"         // Google is a service for Google Calendar.
	Github         ServiceName = "Github"         // Github is a service for Github.
	Dropbox        ServiceName = "Dropbox"        // Dropbox is a service for Dropbox.
	Microsoft      ServiceName = "Microsoft"      // Microsoft is a service for Microsoft.
)

type ServiceJSON struct {
	Name     ServiceName    `json:"name"`      // Name of the service
	Action   []ActionJSON   `json:"actions"`   // List of actions for the service
	Reaction []ReactionJSON `json:"reactions"` // List of reactions for the service
}

// GithubToken represents the GithubToken entity in the database.
type Service struct {
	Id          uint64      `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`                    // Unique identifier for the service
	Name        ServiceName `                                 json:"name"         binding:"required"` // Name of the service
	Description string      `                                 json:"description"  binding:"required"` // Description of the service
	Oauth       bool        `                                 json:"oauth"        binding:"required"` // Whether the service uses OAuth
	Icon        string      `                                 json:"icon"         binding:"required"` // Icon for the service
	Color       string      `                                 json:"color"        binding:"required"` // Color for the service
	CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`                      // Time when the service was created
	UpdateAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`                       // Time when the service was last updated
}

// Errors Messages.
var (
	ErrNotOauthService = errors.New(
		"service is not an oauth service",
	) // Error message for non-OAuth services
)

type MobileTokenRequest struct {
	AccessToken  string    `json:"accessToken"`               // Access token for the user
	RefreshToken string    `json:"refreshToken"`              // Refresh token for the user
	ExpiresIn    time.Time `json:"accessTokenExpirationDate"` // Expiration date for the access token
}
