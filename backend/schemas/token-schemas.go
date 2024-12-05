package schemas

import "time"

// Link represents the Link entity and is associated with LinkUrl.
type Token struct {
	Id           uint64    `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`
	UserId       uint64    `                                          json:"-"` // Foreign key for LinkUrl
	User         User      `gorm:"foreignKey:UserId;references:Id"    json:"user_id"`
	ServiceId    uint64    `                                          json:"-"` // Foreign key for LinkUrl
	Service      Service   `gorm:"foreignKey:ServiceId;references:Id" json:"service_id"`
	Token        string    `                                          json:"token"`
	RefreshToken string    `                                          json:"refresh_token"`
	ExpireAt     time.Time `                                          json:"expireAt"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"createdAt"`
	UpdateAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"updateAt"`
}
