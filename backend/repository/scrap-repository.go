package repository

import (
	"gorm.io/gorm"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type ScrapRepository interface {
	Save(scrap schemas.LinkToLinkUrl)
	Update(scrap schemas.LinkToLinkUrl)
	Delete(scrap schemas.LinkToLinkUrl)
	FindAll() []schemas.LinkToLinkUrl
	FindByUrl(url string) []schemas.LinkToLinkUrl
}

type scrapRepository struct {
	db *schemas.Database
}

func NewScrapRepository(conn *gorm.DB) ScrapRepository {
	err := conn.AutoMigrate(&schemas.LinkToLinkUrl{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &scrapRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *scrapRepository) Save(scrap schemas.LinkToLinkUrl) {
	err := repo.db.Connection.Create(&scrap)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *scrapRepository) Update(scrap schemas.LinkToLinkUrl) {
	err := repo.db.Connection.Save(&scrap)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *scrapRepository) Delete(scrap schemas.LinkToLinkUrl) {
	err := repo.db.Connection.Delete(&scrap)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *scrapRepository) FindAll() []schemas.LinkToLinkUrl {
	var scraps []schemas.LinkToLinkUrl
	err := repo.db.Connection.Find(&scraps)
	if err.Error != nil {
		panic(err.Error)
	}
	return scraps
}

func (repo *scrapRepository) FindByUrl(url string) []schemas.LinkToLinkUrl {
	var scraps []schemas.LinkToLinkUrl
	err := repo.db.Connection.Where(&schemas.LinkToLinkUrl{ActualLink: url}).Find(&scraps)
	if err.Error != nil {
		panic(err.Error)
	}
	return scraps
}
