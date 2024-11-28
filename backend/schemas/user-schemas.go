package schemas

import (
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Username  string    `json:"username" binding:"required" gorm:"type:varchar(100);unique"`
	Email     string    `json:"email" binding:"requiredcredentials" gorm:"type:varchar(100);unique"`
	Password  string    `json:"password" gorm:"type:varchar(100)"` // can be null for Oauth2.0 users
	GithubId  uint64    `json:"github_id" gorm:"default:null"`     // Foreign key for Github
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
