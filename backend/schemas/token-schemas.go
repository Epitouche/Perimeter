package schemas

import (
	"errors"
	"time"
)

// Link represents the Link entity and is associated with LinkURL.
type Token struct {
	Id           uint64    `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`
	UserId       uint64    `                                          json:"-"` // Foreign key for LinkURL
	User         User      `gorm:"foreignKey:UserId;references:Id"    json:"user"`
	ServiceId    uint64    `                                          json:"-"` // Foreign key for LinkURL
	Service      Service   `gorm:"foreignKey:ServiceId;references:Id" json:"service"`
	Token        string    `                                          json:"token"`
	RefreshToken string    `                                          json:"refresh_token"`
	ExpireAt     time.Time `                                          json:"expire_at"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"created_at"`
	UpdateAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"update_at"`
}

// Errors Messages.
var (
	ErrTokenAlreadyExists            = errors.New("token already exists")
	ErrAccessTokenNotFoundInResponse = errors.New("access token not found in response")
	ErrUnableToSaveToken             = errors.New("unable to save token")
)
