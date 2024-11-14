package service

import (
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type GithubTokenService interface {
	Save(schemas.GithubToken) error
	Update(schemas.GithubToken) error
	Delete(schemas.GithubToken) error
	FindAll() []schemas.GithubToken
}

type githubTokenService struct {
	repository repository.GithubTokenRepository
}

func NewGithubTokenService(videoRepository repository.GithubTokenRepository) GithubTokenService {
	return &githubTokenService{
		repository: videoRepository,
	}
}

func (service *githubTokenService) Save(token schemas.GithubToken) error {
	service.repository.Save(token)
	return nil
}

func (service *githubTokenService) Update(token schemas.GithubToken) error {
	service.repository.Update(token)
	return nil
}

func (service *githubTokenService) Delete(token schemas.GithubToken) error {
	service.repository.Delete(token)
	return nil
}

func (service *githubTokenService) FindAll() []schemas.GithubToken {
	return service.repository.FindAll()
}
