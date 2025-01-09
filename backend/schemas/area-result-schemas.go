package schemas

import (
	"time"
)

// GithubToken represents the GithubToken entity in the database.
type AreaResult struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"                                     json:"id,omitempty"`                      // Unique identifier for the area result
	AreaId    uint64    `                                                                    json:"-"`                                 // Foreign key for Area
	Area      Area      `gorm:"foreignKey:AreaId;references:Id;constraint:OnDelete:CASCADE;" json:"area,omitempty" binding:"required"` // Area that the result belongs to
	Result    string    `                                                                    json:"result"         binding:"required"` // Result of the area
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"                                    json:"created_at"`                        // Time when the area result was created
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"                                    json:"update_at"`                         // Time when the area result was last updated
}
