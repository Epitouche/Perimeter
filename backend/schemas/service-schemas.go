package schemas

import "time"


// GithubToken represents the GithubToken entity in the database
type Service struct {
	Id          uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name          string    `json:"name" gorm:"primary_key;auto_increment"`
	Description          string    `json:"name" gorm:"primary_key;auto_increment"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt   time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP"`
}
