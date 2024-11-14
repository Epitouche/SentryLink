package service

import (
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type LinkService interface {
	Save(schemas.Link) error
	Update(schemas.Link) error
	Delete(schemas.Link) error
	FindAll() []schemas.Link
}

type linkService struct {
	repository repository.LinkRepository
}

func NewLinkService(videoRepository repository.LinkRepository) LinkService {
	return &linkService{
		repository: videoRepository,
	}
}

func (service *linkService) Save(link schemas.Link) error {
	service.repository.Save(link)
	return nil
}

func (service *linkService) Update(link schemas.Link) error {
	service.repository.Update(link)
	return nil
}

func (service *linkService) Delete(link schemas.Link) error {
	service.repository.Delete(link)
	return nil
}

func (service *linkService) FindAll() []schemas.Link {
	return service.repository.FindAll()
}
