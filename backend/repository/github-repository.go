package repository

import (
	"gorm.io/gorm"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type GithubTokenRepository interface {
	Save(token schemas.GithubToken)
	Update(token schemas.GithubToken)
	Delete(token schemas.GithubToken)
	FindAll() []schemas.GithubToken
	FindByAccessToken(accessToken string) []schemas.GithubToken
	FindById(id uint64) schemas.GithubToken
}

// Define a struct that embeds `*schemas.Database` and implements `GithubTokenRepository`
type githubTokenRepository struct {
	db *schemas.Database
}

func NewGithubTokenRepository(conn *gorm.DB) GithubTokenRepository {
	err := conn.AutoMigrate(&schemas.GithubToken{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &githubTokenRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *githubTokenRepository) Save(token schemas.GithubToken) {
	err := repo.db.Connection.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}

}

func (repo *githubTokenRepository) Update(token schemas.GithubToken) {
	err := repo.db.Connection.Save(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *githubTokenRepository) Delete(token schemas.GithubToken) {
	err := repo.db.Connection.Delete(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *githubTokenRepository) FindAll() []schemas.GithubToken {
	var tokens []schemas.GithubToken
	err := repo.db.Connection.Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *githubTokenRepository) FindByAccessToken(accessToken string) []schemas.GithubToken {
	var tokens []schemas.GithubToken
	err := repo.db.Connection.Where(&schemas.GithubToken{AccessToken: accessToken}).Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *githubTokenRepository) FindById(id uint64) schemas.GithubToken {
	var tokens schemas.GithubToken
	err := repo.db.Connection.Where(&schemas.GithubToken{Id: id}).First(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}
