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
	repo.db.Connection.Create(&token)
}

func (repo *githubTokenRepository) Update(token schemas.GithubToken) {
	repo.db.Connection.Save(&token)
}

func (repo *githubTokenRepository) Delete(token schemas.GithubToken) {
	repo.db.Connection.Delete(&token)
}

func (repo *githubTokenRepository) FindAll() []schemas.GithubToken {
	var tokens []schemas.GithubToken
	repo.db.Connection = repo.db.Connection.Debug() // Enable debugging
	repo.db.Connection.Find(&tokens)
	return tokens
}
