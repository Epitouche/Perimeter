package schemas

import (
	"time"
)

type AreaMessage struct {
	ActionOption   string `json:"action_option"   binding:"required"` // Action option
	ActionId       uint64 `json:"action_id"`                          // Unique identifier for the action
	ReactionOption string `json:"reaction_option" binding:"required"` // Reaction option
	ReactionId     uint64 `json:"reaction_id"`                        // Unique identifier for the reaction
}

type Area struct {
	Id             uint64    `gorm:"primaryKey;autoIncrement"            json:"id,omitempty"`                          // Unique identifier for the area
	UserId         uint64    `                                           json:"-"`                                     // Foreign key for User
	User           User      `gorm:"foreignKey:UserId;references:Id"     json:"user,omitempty"     binding:"required"` // User that the area belongs to
	ActionOption   string    `                                           json:"action_option"      binding:"required"` // Action option
	ActionId       uint64    `                                           json:"-"`                                     // Foreign key for Action
	Action         Action    `gorm:"foreignKey:ActionId;references:Id"   json:"action,omitempty"   binding:"required"` // Action that the area belongs to
	ReactionOption string    `                                           json:"reaction_option"    binding:"required"` // Reaction option
	ReactionId     uint64    `                                           json:"-"`                                     // Foreign key for Reaction
	Reaction       Reaction  `gorm:"foreignKey:ReactionId;references:Id" json:"reaction,omitempty" binding:"required"` // Reaction that the area belongs to
	Enable         bool      `gorm:"default:true"                        json:"enable"`                                // Enable or disable the area
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"           json:"created_at"`                            // Time when the area was created
	UpdateAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"           json:"update_at"`                             // Time when the area was last updated
}
