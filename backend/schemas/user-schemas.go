package schemas

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(100);unique"  json:"username"     binding:"required"`
	Email     string    `gorm:"type:varchar(100);unique"  json:"email"        binding:"required"`
	Password  string    `gorm:"type:varchar(100)"         json:"password"` // can be null for Oauth2.0 users
	TokenId   uint64    `                                 json:"token_id"` // Foreign key for LinkURL
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserAllInfo struct {
	User   User    `json:"user"`
	Tokens []Token `json:"tokens"`
}
