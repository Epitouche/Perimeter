package schemas

import (
	"time"
)

type Area struct {
	Id             uint64    `gorm:"primaryKey;autoIncrement"            json:"id,omitempty"`
	UserId         uint64    `                                           json:"-"` // Foreign key for User
	User           User      `gorm:"foreignKey:UserId;references:Id"     json:"user_id,omitempty"     binding:"required"`
	OptionAction   string    `                                           json:"action_option" binding:"required"`
	ActionId       uint64    `                                           json:"-"` // Foreign key for Action
	Action         Action    `gorm:"foreignKey:ActionId;references:Id"   json:"action_id,omitempty"   binding:"required"`
	OptionReaction string    `                                           json:"reaction_option" binding:"required"`
	ReactionId     uint64    `                                           json:"-"` // Foreign key for Reaction
	Reaction       Reaction  `gorm:"foreignKey:ReactionId;references:Id" json:"reaction_id,omitempty" binding:"required"`
	Enable         bool      `gorm:"default:true"                        json:"enable"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"           json:"createdAt"`
	UpdateAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"           json:"update_at"`
}
