package schemas

import (
	"time"
)

// GithubToken represents the GithubToken entity in the database.
type AreaResult struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"        json:"id,omitempty"`
	AreaId    uint64    `                                       json:"-"` // Foreign key for Area
	Area      Area      `gorm:"foreignKey:AreaId;references:Id" json:"area,omitempty" binding:"required"`
	Result    string    `                                       json:"result"         binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"       json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"       json:"update_at"`
}
