package schemas

import (
	"fmt"
	"time"
)

// User represents the User entity in the database.
type User struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"  json:"id,omitempty"`                    // Unique identifier for the user
	Username  string    `gorm:"type:varchar(100);unique"  json:"username"     binding:"required"` // Username of the user
	Email     string    `gorm:"type:varchar(100);unique"  json:"email"        binding:"required"` // Email of the user
	Password  string    `gorm:"type:varchar(100)"         json:"password"`                        // can be null for Oauth2.0 users
	TokenId   uint64    `                                 json:"token_id"`                        // Foreign key for LinkURL
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`                      // Time when the user was created
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`                      // Time when the user was last updated
}

type UserCredentials struct {
	Username string `json:"username"` // Username of the user
	Email    string `json:"email"`    // Email of the user
}

type UserAllInfo struct {
	User   User    `json:"user"`   // User
	Tokens []Token `json:"tokens"` // List of tokens
}

var (
	ErrUsernameTooShort = fmt.Errorf("username must be at least %d characters long", UsernameMinimumLength)
	ErrPasswordTooShort = fmt.Errorf("password must be at least %d characters long", PasswordMinimumLength)
	ErrEmailTooShort    = fmt.Errorf("email must be at least %d characters long", EmailMinimumLength)
)
