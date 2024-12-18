package schemas

import (
	"errors"
	"time"
)

type ServiceName string

const (
	Spotify        ServiceName = "spotify"
	OpenWeatherMap ServiceName = "openWeatherMap"
	Timer          ServiceName = "timer"
	Gmail          ServiceName = "gmail"
	Github         ServiceName = "github"
	Dropbox        ServiceName = "dropbox"
)

type ServiceJSON struct {
	Name     ServiceName    `json:"name"`
	Action   []ActionJSON   `json:"actions"`
	Reaction []ReactionJSON `json:"reactions"`
}

// GithubToken represents the GithubToken entity in the database.
type Service struct {
	Id          uint64      `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Name        ServiceName `                                 json:"name"         binding:"required"`
	Description string      `                                 json:"description"  binding:"required"`
	CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdateAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

// Errors Messages.
var (
	ErrNotOauthService = errors.New("service is not an oauth service")
)
