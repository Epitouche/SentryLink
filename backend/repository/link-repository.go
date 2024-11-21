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
	err := repo.db.Connection.Create(&video)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *linkRepository) Update(video schemas.Link) {
	err := repo.db.Connection.Save(&video)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *linkRepository) Delete(video schemas.Link) {
	err := repo.db.Connection.Delete(&video)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *linkRepository) FindAll() []schemas.Link {
	var links []schemas.Link
	err := repo.db.Connection.Preload("UrlId").Find(&links)
	if err.Error != nil {
		panic(err.Error)
	}
	return links
}
