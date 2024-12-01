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
	Name        string    `                                 json:"name"         binding:"required"`
	Description string    `                                 json:"description"  binding:"required"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}
