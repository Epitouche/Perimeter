package schemas

import "time"

type ServiceName string

const (
	Spotify        ServiceName = "Spotify"
	OpenWeatherMap ServiceName = "OpenWeatherMap"
	Timer          ServiceName = "Timer"
)

// GithubToken represents the GithubToken entity in the database
type Service struct {
	Id          uint64    `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt    time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP"`
}
