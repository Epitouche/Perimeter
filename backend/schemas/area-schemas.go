package schemas

import (
	"encoding/json"
	"time"
)

// AreaMessage represents the schema for an area message in the system.
// It contains information about the action and reaction options, their respective IDs,
// and additional metadata such as title, description, and action refresh rate.
type AreaMessage struct {
	ActionOption      json.RawMessage `gorm:"type:jsonb" json:"action_option"       binding:"required"` // The option of the action
	ActionId          uint64          `                  json:"action_id"`                              // Foreign key for Action
	ReactionOption    json.RawMessage `gorm:"type:jsonb" json:"reaction_option"     binding:"required"` // The option of the reaction
	ReactionId        uint64          `                  json:"reaction_id"`                            // Foreign key for Reaction
	Title             string          `                  json:"title"               binding:"required"` // The title of the area
	Description       string          `                  json:"description"         binding:"required"` // The description of the area
	ActionRefreshRate int             `                  json:"action_refresh_rate" binding:"required"` // The refresh rate for the action
}

// Area represents a specific area in the system with associated actions and reactions.
// It includes metadata such as title, description, and timestamps for creation and updates.
//
// Fields:
//   - Id: Unique identifier for the area.
//   - UserId: Foreign key for User.
//   - User: User that the area belongs to.
//   - ActionOption: The option of the action.
//   - ActionId: Foreign key for Action.
//   - Action: Action that the area belongs to.
//   - ReactionOption: The option of the reaction.
//   - ReactionId: Foreign key for Reaction.
//   - Reaction: Reaction that the area belongs to.
//   - Enable: Enable or disable the area.
//   - Title: The title of the area.
//   - Description: The description of the area.
//   - StorageVariable: The storage variable of the area.
//   - CreatedAt: Time when the area was created.
//   - UpdateAt: Time when the area was last updated.
//   - ActionRefreshRate: The refresh rate for the action.
type Area struct {
	Id                uint64          `gorm:"primaryKey;autoIncrement"                                     json:"id,omitempty"`                           // Unique identifier for the area
	UserId            uint64          `                                                                    json:"-"`                                      // Foreign key for User
	User              User            `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"user,omitempty"      binding:"required"` // User that the area belongs to
	ActionOption      json.RawMessage `gorm:"type:jsonb"                                                   json:"action_option"       binding:"required"` // The option of the action
	ActionId          uint64          `                                                                    json:"-"`                                      // Foreign key for Action
	Action            Action          `gorm:"foreignKey:ActionId;references:Id"                            json:"action,omitempty"    binding:"required"` // Action that the area belongs to
	ReactionOption    json.RawMessage `gorm:"type:jsonb"                                                   json:"reaction_option"     binding:"required"` // The option of the reaction
	ReactionId        uint64          `                                                                    json:"-"`                                      // Foreign key for Reaction
	Reaction          Reaction        `gorm:"foreignKey:ReactionId;references:Id"                          json:"reaction,omitempty"  binding:"required"` // Reaction that the area belongs to
	Enable            bool            `gorm:"default:true"                                                 json:"enable"`                                 // Enable or disable the area
	Title             string          `                                                                    json:"title"               binding:"required"` // The title of the area
	Description       string          `                                                                    json:"description"         binding:"required"` // The description of the area
	StorageVariable   json.RawMessage `gorm:"type:jsonb"                                                   json:"storage_variable"`                       // The storage variable of the area
	CreatedAt         time.Time       `gorm:"default:CURRENT_TIMESTAMP"                                    json:"created_at"`                             // Time when the area was created
	UpdateAt          time.Time       `gorm:"default:CURRENT_TIMESTAMP"                                    json:"update_at"`                              // Time when the area was last updated
	ActionRefreshRate uint64          `                                                                    json:"action_refresh_rate" binding:"required"` // The refresh rate for the action
}
