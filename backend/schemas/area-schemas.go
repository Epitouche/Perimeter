package schemas

import (
	"encoding/json"
	"time"
)

type AreaMessage struct {
	ActionOption   json.RawMessage `gorm:"type:jsonb" json:"action_option"   binding:"required"`
	ActionId       uint64          `                  json:"action_id"` // Foreign key for Action
	ReactionOption json.RawMessage `gorm:"type:jsonb" json:"reaction_option" binding:"required"`
	ReactionId     uint64          `                  json:"reaction_id"` // Foreign key for Reaction
	Title          string          `                  json:"title"           binding:"required"`
	Description    string          `                  json:"description"     binding:"required"`
}

type Area struct {
	Id             uint64          `gorm:"primaryKey;autoIncrement"            json:"id,omitempty"`
	UserId         uint64          `                                           json:"-"` // Foreign key for User
	User           User            `gorm:"foreignKey:UserId;references:Id"     json:"user,omitempty"     binding:"required"`
	ActionOption   json.RawMessage `gorm:"type:jsonb"                          json:"action_option"      binding:"required"`
	ActionId       uint64          `                                           json:"-"` // Foreign key for Action
	Action         Action          `gorm:"foreignKey:ActionId;references:Id"   json:"action,omitempty"   binding:"required"`
	ReactionOption json.RawMessage `gorm:"type:jsonb"                          json:"reaction_option"    binding:"required"`
	ReactionId     uint64          `                                           json:"-"` // Foreign key for Reaction
	Reaction       Reaction        `gorm:"foreignKey:ReactionId;references:Id" json:"reaction,omitempty" binding:"required"`
	Enable         bool            `gorm:"default:true"                        json:"enable"`
	Title          string          `                                           json:"title"              binding:"required"`
	Description    string          `                                           json:"description"        binding:"required"`
	CreatedAt      time.Time       `gorm:"default:CURRENT_TIMESTAMP"           json:"createdAt"`
	UpdateAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP"           json:"update_at"`
}
