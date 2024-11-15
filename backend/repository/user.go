package repository

import (
	"gorm.io/gorm"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type UserRepository interface {
	Save(user schemas.User)
	Update(user schemas.User)
	Delete(user schemas.User)
	FindAll() []schemas.User
	FindByEmail(email string) []schemas.User
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

func (repo *userRepository) Save(user schemas.User) {
	repo.db.Connection.Create(&user)
}

func (repo *userRepository) Update(user schemas.User) {
	repo.db.Connection.Save(&user)
}

func (repo *userRepository) Delete(user schemas.User) {
	repo.db.Connection.Delete(&user)
}

func (repo *userRepository) FindAll() []schemas.User {
	var users []schemas.User
	repo.db.Connection = repo.db.Connection.Debug() // Enable debugging
	repo.db.Connection.Find(&users)
	return users
}

func (repo *userRepository) FindByEmail(email string) []schemas.User {
	var users []schemas.User
	repo.db.Connection = repo.db.Connection.Debug() // Enable debugging
	repo.db.Connection.Find(&users).Where("email = ?", email)
	return users
}
