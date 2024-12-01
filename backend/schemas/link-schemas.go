package schemas

import "time"

// LinkUrl represents the URL entity in the database.
type LinkUrl struct {
	Id  uint64 `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Url string `gorm:"type:varchar(256);unique" json:"url"          binding:"required"`
}

// Link represents the Link entity and is associated with LinkUrl.
type Link struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement"        json:"id,omitempty"`
	LinkId     uint64    `                                       json:"-"` // Foreign key for LinkUrl
	UrlId      LinkUrl   `gorm:"foreignKey:LinkId;references:Id" json:"url_id,omitempty"`
	StatusCode uint64    `                                       json:"status_code"      binding:"required"`
	Response   string    `gorm:"type:varchar(100)"               json:"response"         binding:"required"`
	Ping       uint64    `                                       json:"ping"             binding:"required"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"       json:"created_at"`
}
