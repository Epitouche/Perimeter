package schemas

import (
	"encoding/json"
	"time"
)

// ReactionJSON represents the structure of a reaction with its name and description.
// Name is the name of the reaction.
// Description is the description of the reaction.
type ReactionJSON struct {
	Name        string `json:"name"`        // Name of the reaction
	Description string `json:"description"` // Description of the reaction
}

// Reaction represents a reaction entity with details such as name, description, associated service, and options.
// It includes metadata like creation and update timestamps.
//
// Fields:
// - Id: Unique identifier for the reaction.
// - Name: Name of the reaction (required).
// - Description: Description of the reaction (required).
// - ServiceId: Foreign key for the associated Service.
// - Service: The Service that the reaction belongs to (required).
// - Option: JSON options for the reaction (required).
// - CreatedAt: Timestamp when the reaction was created.
// - UpdateAt: Timestamp when the reaction was last updated.
type Reaction struct {
	Id          uint64          `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`                         // Unique identifier for the reaction
	Name        string          `                                          json:"name"              binding:"required"` // Name of the reaction
	Description string          `                                          json:"description"       binding:"required"` // Description of the reaction
	ServiceId   uint64          `                                          json:"-"`                                    // Foreign key for Service
	Service     Service         `gorm:"foreignKey:ServiceId;references:Id" json:"service,omitempty" binding:"required"` // Service that the reaction belongs to
	Option      json.RawMessage `gorm:"type:jsonb"                         json:"option"            binding:"required"` // Option of the reaction
	CreatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"created_at"`                           // Time when the reaction was created
	UpdateAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`                            // Time when the reaction was last updated
}
