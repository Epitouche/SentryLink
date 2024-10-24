package repository

import (
	"github.com/Tom-Mendy/SentryLink/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LinkRepository interface {
	Save(link schemas.Link)
	Update(link schemas.Link)
	Delete(link schemas.Link)
	FindAll() []schemas.Link
}

type database struct {
	connection *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schemas.Link{}, &schemas.LinkInfo{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(video schemas.Link) {
	db.connection.Create(&video)
}

func (db *database) Update(video schemas.Link) {
	db.connection.Save(&video)
}

func (db *database) Delete(video schemas.Link) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []schemas.Link {
	var links []schemas.Link
	db.connection.Set("gorm:auto_preload", true).Find(&links)
	return links
}
