package schemas

import (
	"time"
)

// GithubToken represents the GithubToken entity in the database.
type Reaction struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Name        string    `binding:"required"               json:"name"`
	Description string    `binding:"required"               json:"description"`
	ServiceId   uint64    `json:"-"` // Foreign key for Service
	Service     Service   `binding:"required"               gorm:"foreignKey:ServiceId;references:Id" json:"service_id,omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}
