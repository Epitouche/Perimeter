package schemas

import (
	"time"
)

// GithubToken represents the GithubToken entity in the database
type Action struct {
	Id          uint64    `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	ServiceId   uint64    `json:"-"` // Foreign key for Service
	Service     Service   `json:"service_id,omitempty" binding:"required" gorm:"foreignKey:ServiceId;references:Id"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt    time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP"`
}
