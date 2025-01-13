package schemas

import (
	"errors"
	"time"
)

// Link represents the Link entity and is associated with LinkURL.
type Token struct {
	Id           uint64    `gorm:"primaryKey;autoIncrement"                                     json:"id,omitempty"`  // Unique identifier for the token
	UserId       uint64    `                                                                    json:"-"`             // Foreign key for LinkURL
	User         User      `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"user"`          // User that the token belongs to
	ServiceId    uint64    `                                                                    json:"-"`             // Foreign key for LinkURL
	Service      Service   `gorm:"foreignKey:ServiceId;references:Id"                           json:"service"`       // Service that the token belongs to
	Token        string    `                                                                    json:"token"`         // Token
	RefreshToken string    `                                                                    json:"refresh_token"` // Refresh token
	ExpireAt     time.Time `                                                                    json:"expire_at"`     // Time when the token expires
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"                                    json:"created_at"`    // Time when the token was created
	UpdateAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"                                    json:"update_at"`     // Time when the token was last updated
}

// Errors Messages.
var (
	ErrTokenAlreadyExists            = errors.New("token already exists")
	ErrAccessTokenNotFoundInResponse = errors.New("access token not found in response")
	ErrUnableToSaveToken             = errors.New("unable to save token")
	ErrTokenBelongToUser             = errors.New("token belongs to user")
)
