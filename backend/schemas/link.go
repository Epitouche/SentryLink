package schemas

import "time"

type LinkUrl struct {
	// Id  int64  `json:"id" binding:"required, uid" gorm:"primery_key;auto_increment"`
	Id  uint64 `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Url string `json:"url" binding:"required" gorm:"type:varchar(256);UNIQUE"`
}

type Link struct {
	// Id         int64     `json:"id" binding:"required, uid" gorm:"primery_key;auto_increment"`
	Id         uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UrlId      LinkUrl   `json:"url_id,omitempty" gorm:"foreignKey:LinkId"`
	LinkId     uint64    `json:"-"`
	StatusCode uint64    `json:"status_code" binding:"required"`
	Response   string    `json:"response" binding:"required" gorm:"type_varchar(100)"`
	Ping       uint64    `json:"ping" binding:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
