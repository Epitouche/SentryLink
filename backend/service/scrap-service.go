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
	// Scrap(string) (string, error)
}

type scrapService struct {
	repository repository.ScrapRepository
}

func NewScrapService(scrapRepository repository.ScrapRepository) ScrapService {
	return &scrapService{
		repository: scrapRepository,
	}
}

func (service *scrapService) Save(link schemas.LinkToLinkUrl) error {
	service.repository.Save(link)
	return nil
}

func (service *scrapService) Update(link schemas.LinkToLinkUrl) error {
	service.repository.Update(link)
	return nil
}

func (service *scrapService) Delete(link schemas.LinkToLinkUrl) error {
	service.repository.Delete(link)
	return nil
}

func (service *scrapService) FindAll() []schemas.LinkToLinkUrl {
	return service.repository.FindAll()
}

// func (service *scrapService) Scrap(pageURL string) (string, error) {
// 	return ExtractLinks(pageURL)
// }

