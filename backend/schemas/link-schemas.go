package schemas

import "time"

// LinkUrl represents the URL entity in the database.
type LinkUrl struct {
	Id  uint64 `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Url string `binding:"required"                gorm:"type:varchar(256);unique" json:"url"`
}

// Link represents the Link entity and is associated with LinkUrl.
type Link struct {
	Id         uint64    `gorm:"primary_key;auto_increment"      json:"id,omitempty"`
	LinkId     uint64    `json:"-"` // Foreign key for LinkUrl
	UrlId      LinkUrl   `gorm:"foreignKey:LinkId;references:Id" json:"url_id,omitempty"`
	StatusCode uint64    `binding:"required"                     json:"status_code"`
	Response   string    `binding:"required"                     gorm:"type:varchar(100)" json:"response"`
	Ping       uint64    `binding:"required"                     json:"ping"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"       json:"created_at"`
}
