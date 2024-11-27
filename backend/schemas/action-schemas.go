package schemas

import "time"

// GithubToken represents the GithubToken entity in the database
type Action struct {
	Id          uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" gorm:"primary_key;auto_increment"`
	Description string    `json:"description" gorm:"primary_key;auto_increment"`
	ServiceId   uint64    `json:"-"` // Foreign key for ServiceId
	ServiceRef  Service   `json:"service_id,omitempty" gorm:"foreignKey:ServiceId;references:Id"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt    time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP"`
}
