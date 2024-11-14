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

type linkRepository struct {
	db *schemas.Database
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
	return &linkRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *linkRepository) Save(video schemas.Link) {
	repo.db.Connection.Create(&video)
}

func (repo *linkRepository) Update(video schemas.Link) {
	repo.db.Connection.Save(&video)
}

func (repo *linkRepository) Delete(video schemas.Link) {
	repo.db.Connection.Delete(&video)
}

func (repo *linkRepository) FindAll() []schemas.Link {
	var links []schemas.Link
	repo.db.Connection = repo.db.Connection.Debug() // Enable debugging
	repo.db.Connection.Preload("UrlId").Find(&links)
	return links
}
