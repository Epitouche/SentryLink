package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Tom-Mendy/SentryLink/schemas"
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

func NewLinkRepository() LinkRepository {
	dsn := "host=postgres user=admin password=password dbname=mydatabase port=5432 sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// err = conn.AutoMigrate(&schemas.Link{}, &schemas.LinkInfo{})
	err = conn.AutoMigrate(&schemas.LinkUrl{}, &schemas.Link{})
	if err != nil {
		panic("failed to migrate database")
	}
	println("Connection to database established")
	return &database{
		connection: conn,
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
	db.connection = db.connection.Debug() // Enable debugging
	db.connection.Preload("UrlId").Find(&links)
	return links
}
