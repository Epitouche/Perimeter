package schemas

import (
	"encoding/json"
	"time"
)

// ActionJSON represents the JSON schema for an action.
// It includes the name and description of the action.
type ActionJSON struct {
	Name        string `json:"name"`        // The name of the action
	Description string `json:"description"` // The description of the action
}

// GithubToken represents the GithubToken entity in the database.
type Action struct {
	Id                 uint64          `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`                            // The unique identifier for the action
	Name               string          `                                          json:"name"                 binding:"required"` // The name of the action
	Description        string          `                                          json:"description"          binding:"required"` // The description of the action
	ServiceId          uint64          `                                          json:"-"`                                       // Foreign key for Service
	Service            Service         `gorm:"foreignKey:ServiceId;references:Id" json:"service,omitempty"    binding:"required"` // The service that the action belongs to
	Option             json.RawMessage `gorm:"type:jsonb"                         json:"option"               binding:"required"` // The option of the action
	CreatedAt          time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"createdAt"`
	UpdateAt           time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`
	MinimumRefreshRate uint64          `                                          json:"minimum_refresh_rate" binding:"required"`
}
