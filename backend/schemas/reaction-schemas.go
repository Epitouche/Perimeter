package schemas

import (
	"encoding/json"
	"time"
)

type ReactionJSON struct {
	Name        string `json:"name"`        // Name of the reaction
	Description string `json:"description"` // Description of the reaction
}

type Reaction struct {
	Id          uint64          `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`
	Name        string          `                                          json:"name"              binding:"required"`
	Description string          `                                          json:"description"       binding:"required"`
	ServiceId   uint64          `                                          json:"-"` // Foreign key for Service
	Service     Service         `gorm:"foreignKey:ServiceId;references:Id" json:"service,omitempty" binding:"required"`
	Option      json.RawMessage `gorm:"type:jsonb"                         json:"option"            binding:"required"`
	CreatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"created_at"`
	UpdateAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`
}
