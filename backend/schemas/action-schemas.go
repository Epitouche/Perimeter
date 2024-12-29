package schemas

import (
	"time"
)

type ActionJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GithubToken represents the GithubToken entity in the database.
type Action struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`                         // Unique identifier for the action
	Name        string    `                                          json:"name"              binding:"required"` // Name of the action
	Description string    `                                          json:"description"       binding:"required"` // Description of the action
	ServiceId   uint64    `                                          json:"-"`                                    // Foreign key for Service
	Service     Service   `gorm:"foreignKey:ServiceId;references:Id" json:"service,omitempty" binding:"required"` // Service that the action belongs to
	Option      string    `                                          json:"option"            binding:"required"` // Option for the action
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"created_at"`                           // Time when the action was created
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`                            // Time when the action was last updated
}
