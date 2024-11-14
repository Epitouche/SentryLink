package schemas

import "time"

// LinkUrl represents the URL entity in the database
type LinkUrl struct {
	Id  uint64 `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Url string `json:"url" binding:"required" gorm:"type:varchar(256);unique"`
}

// Link represents the Link entity and is associated with LinkUrl
type Link struct {
	Id         uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	LinkId     uint64    `json:"-"` // Foreign key for LinkUrl
	UrlId      LinkUrl   `json:"url_id,omitempty" gorm:"foreignKey:LinkId;references:Id"`
	StatusCode uint64    `json:"status_code" binding:"required"`
	Response   string    `json:"response" binding:"required" gorm:"type:varchar(100)"`
	Ping       uint64    `json:"ping" binding:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
