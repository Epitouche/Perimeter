package schemas

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(100);unique"  json:"username"     binding:"required"`
	Email     string    `gorm:"type:varchar(100);unique"  json:"email"        binding:"requiredcredentials"`
	Password  string    `gorm:"type:varchar(100)"         json:"password"` // can be null for Oauth2.0 users
	TokenId   uint64    `                                 json:"token_id"` // Foreign key for LinkUrl
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
