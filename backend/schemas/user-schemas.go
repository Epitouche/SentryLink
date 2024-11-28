package schemas

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Username  string    `binding:"required"                gorm:"type:varchar(100);unique" json:"username"`
	Email     string    `binding:"requiredcredentials"     gorm:"type:varchar(100);unique" json:"email"`
	Password  string    `gorm:"type:varchar(100)"          json:"password"` // can be null for Oauth2.0 users
	GithubId  uint64    `gorm:"default:null"               json:"github_id"`     // Foreign key for Github
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"updated_at"`
}
