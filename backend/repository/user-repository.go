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
	FindByUserName(username string) []schemas.User
	FindById(id uint64) schemas.User
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
	err := repo.db.Connection.Create(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *userRepository) Update(user schemas.User) {
	err := repo.db.Connection.Save(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *userRepository) Delete(user schemas.User) {
	err := repo.db.Connection.Delete(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *userRepository) FindAll() []schemas.User {
	var users []schemas.User
	err := repo.db.Connection.Find(&users)
	if err.Error != nil {
		panic(err.Error)
	}
	return users
}

func (repo *userRepository) FindByEmail(email string) []schemas.User {
	var users []schemas.User
	err := repo.db.Connection.Where(&schemas.User{Email: email}).Find(&users)
	if err.Error != nil {
		panic(err.Error)
	}
	return users
}

func (repo *userRepository) FindByUserName(username string) []schemas.User {
	var users []schemas.User
	err := repo.db.Connection.Where(&schemas.User{Username: username}).Find(&users)
	if err.Error != nil {
		panic(err.Error)
	}
	return users
}

func (repo *userRepository) FindById(id uint64) schemas.User {
	var users schemas.User
	err := repo.db.Connection.Where(&schemas.User{Id: id}).First(&users)
	if err.Error != nil {
		panic(err.Error)
	}
	return users
}
