package schemas

import (
	"time"
)

// AreaResult represents the result of an area with associated metadata.
// It includes the unique identifier, foreign key to the Area, the result string,
// and timestamps for creation and last update.
//
// Fields:
// - Id: Unique identifier for the area result.
// - AreaId: Foreign key for the associated Area.
// - Area: The Area that the result belongs to, with a required binding and cascade delete constraint.
// - Result: The result of the area, with a required binding.
// - CreatedAt: Timestamp for when the area result was created, with a default value of the current timestamp.
// - UpdateAt: Timestamp for when the area result was last updated, with a default value of the current timestamp.
type AreaResult struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"                                     json:"id,omitempty"`                      // Unique identifier for the area result
	AreaId    uint64    `                                                                    json:"-"`                                 // Foreign key for Area
	Area      Area      `gorm:"foreignKey:AreaId;references:Id;constraint:OnDelete:CASCADE;" json:"area,omitempty" binding:"required"` // Area that the result belongs to
	Result    string    `                                                                    json:"result"         binding:"required"` // Result of the area
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"                                    json:"created_at"`                        // Time when the area result was created
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"                                    json:"update_at"`                         // Time when the area result was last updated
}
