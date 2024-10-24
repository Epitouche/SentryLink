package schemas

import "time"

type Link struct {
	// Id  int64  `json:"id" binding:"required, uid" gorm:"primery_key;auto_increment"`
	Id  uint64 `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Url string `json:"url" binding:"required, url" gorm:"type_varchar(100); UNIQUE"`
}

type LinkInfo struct {
	// Id         int64     `json:"id" binding:"required, uid" gorm:"primery_key;auto_increment"`
	Id         uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UrlId      Link      `json:"url_id" binding:"required, url" gorm:"foreignKey:UrlId"`
	StatusCode int       `json:"status_code" binding:"required, number" gorm:"type_varchar(100)"`
	Response   string    `json:"response" binding:"required, html" gorm:"type_varchar(100)"`
	Ping       float64   `json:"ping" binding:"required, float, min=0" gorm:"type_varchar(100)"`
	CreatedAt  time.Time `json:"created_at" binding:"required, datetime" gorm:"default:CURRENT_TIMESTAMP"`
}
