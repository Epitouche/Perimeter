package schemas

import (
	"encoding/json"
	"time"
)

// ActionJSON represents the structure of an action with its name and description.
// Name is the name of the action.
// Description is the description of the action.
type ActionJSON struct {
	Name        string `json:"name"`        // The name of the action
	Description string `json:"description"` // The description of the action
}

// Action represents an action entity with various attributes such as ID, name, description, service, options, and timestamps.
// Fields:
// - Id: The unique identifier for the action.
// - Name: The name of the action. This field is required.
// - Description: The description of the action. This field is required.
// - ServiceId: The foreign key for the associated service.
// - Service: The service that the action belongs to. This field is required.
// - Option: The option of the action, stored as a JSONB type. This field is required.
// - CreatedAt: The timestamp when the action was created. Defaults to the current timestamp.
// - UpdateAt: The timestamp when the action was last updated. Defaults to the current timestamp.
// - MinimumRefreshRate: The minimum refresh rate for the action. This field is required.
type Action struct {
	Id                 uint64          `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`                            // The unique identifier for the action
	Name               string          `                                          json:"name"                 binding:"required"` // The name of the action
	Description        string          `                                          json:"description"          binding:"required"` // The description of the action
	ServiceId          uint64          `                                          json:"-"`                                       // Foreign key for Service
	Service            Service         `gorm:"foreignKey:ServiceId;references:Id" json:"service,omitempty"    binding:"required"` // The service that the action belongs to
	Option             json.RawMessage `gorm:"type:jsonb"                         json:"option"               binding:"required"` // The option of the action
	CreatedAt          time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"createdAt"`                               // The timestamp when the action was created
	UpdateAt           time.Time       `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`                               // The timestamp when the action was last updated
	MinimumRefreshRate uint64          `                                          json:"minimum_refresh_rate" binding:"required"` // The minimum refresh rate for the action
}
