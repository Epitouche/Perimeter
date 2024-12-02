package schemas

import "time"

type ServiceName string

const (
	Spotify        ServiceName = "spotify"
	OpenWeatherMap ServiceName = "openWeatherMap"
	Timer          ServiceName = "timer"
)

type ServiceJson struct {
	Name     string         `json:"name"`
	Action   []ActionJson   `json:"actions"`
	Reaction []ReactionJson `json:"reactions"`
}

// GithubToken represents the GithubToken entity in the database.
type Service struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Name        string    `binding:"required"               json:"name"`
	Description string    `binding:"required"               json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}
