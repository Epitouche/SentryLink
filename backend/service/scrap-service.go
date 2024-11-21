package service

import (
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type ScrapService interface {
	Save(schemas.LinkToLinkUrl) error
	Update(schemas.LinkToLinkUrl) error
	Delete(schemas.LinkToLinkUrl) error
	FindAll() []schemas.LinkToLinkUrl
	Scrap(string) (string, error)
}

type scrapService struct {
	repository repository.ScrapRepository
}

// func NewScrapService(scrapRepository repository.ScrapRepository) ScrapService {
// 	return &scrapService{
// 		repository: scrapRepository,
// 	}
// }


