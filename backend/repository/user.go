package repository

import (
	"gorm.io/gorm"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type UserRepository interface {
	Save(token schemas.User)
	Update(token schemas.User)
	Delete(token schemas.User)
	FindAll() []schemas.User
}

// Define a struct that embeds `*schemas.Database` and implements `UserRepository`
type userRepository struct {
	db *schemas.Database
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	err := conn.AutoMigrate(&schemas.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &userRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *userRepository) Save(token schemas.User) {
	repo.db.Connection.Create(&token)
}

func (repo *userRepository) Update(token schemas.User) {
	repo.db.Connection.Save(&token)
}

func (repo *userRepository) Delete(token schemas.User) {
	repo.db.Connection.Delete(&token)
}

func (repo *userRepository) FindAll() []schemas.User {
	var tokens []schemas.User
	repo.db.Connection = repo.db.Connection.Debug() // Enable debugging
	repo.db.Connection.Find(&tokens)
	return tokens
}
