package schemas

import (
	"time"
)

type ReactionJSON struct {
	Name        string `json:"name"`        // Name of the reaction
	Description string `json:"description"` // Description of the reaction
}

type Reaction struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`                         // Unique identifier for the reaction
	Name        string    `                                          json:"name"              binding:"required"` // Name of the reaction
	Description string    `                                          json:"description"       binding:"required"` // Description of the reaction
	ServiceId   uint64    `                                          json:"-"`                                    // Foreign key for Service
	Service     Service   `gorm:"foreignKey:ServiceId;references:Id" json:"service,omitempty" binding:"required"` // Service that the reaction belongs to
	Option      string    `                                          json:"option"            binding:"required"` // Option for the reaction
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"created_at"`                           // Time when the reaction was created
	UpdateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`                            // Time when the reaction was last updated
}
