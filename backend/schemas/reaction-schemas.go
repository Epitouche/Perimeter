package schemas

import (
	"time"
)

// GithubToken represents the GithubToken entity in the database.
type Reaction struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`
	Name        string    `                                          json:"name"                 binding:"required"`
	Description string    `                                          json:"description"          binding:"required"`
	ServiceId   uint64    `                                          json:"-"` // Foreign key for Service
	Service     Service   `gorm:"foreignKey:ServiceId;references:Id" json:"service_id,omitempty" binding:"required"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"created_at"`
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`
}
