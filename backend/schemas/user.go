package schemas

import "time"

type User struct {
	Id        uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Username  string    `json:"username" binding:"required" gorm:"type:varchar(100);unique"`
	Password  string    `json:"password" binding:"required" gorm:"type:varchar(100)"`
	Email     string    `json:"email" binding:"required" gorm:"type:varchar(100);unique"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
