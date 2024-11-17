package repository

import (
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

func NewLinkRepository(conn *gorm.DB) LinkRepository {
	err := conn.AutoMigrate(&schemas.LinkUrl{}, &schemas.Link{})
	if err != nil {
		panic("failed to migrate database")
	}
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
	repo.db.Connection.Preload("UrlId").Find(&links)
	return links
}
