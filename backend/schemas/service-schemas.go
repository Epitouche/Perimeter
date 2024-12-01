package schemas

import "time"

type ServiceName string

const (
	Spotify        ServiceName = "Spotify"
	OpenWeatherMap ServiceName = "OpenWeatherMap"
	Timer          ServiceName = "Timer"
)

// GithubToken represents the GithubToken entity in the database.
type Service struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Name        string    `binding:"required"               json:"name"`
	Description string    `binding:"required"               json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}
