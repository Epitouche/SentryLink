package service

import (
	"github.com/Tom-Mendy/SentryLink/database"
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type UserService interface {
	Login(username string, password string) bool
	Registration(username string, email string, password string) bool
}

type userService struct {
	authorizedUsername string
	authorizedPassword string
	repository         repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository:         userRepository,
	}
}

func (service *userService) Login(username string, password string) bool {
	userWiththisUserName := service.repository.FindByUserName(username)
	if len(userWiththisUserName) == 0 {
		return false
	}
	for _, user := range userWiththisUserName {
		if database.DoPasswordsMatch(user.Password, password) {
			return true
		}
	}
	return false
}

func (service *userService) Registration(username string, email string, password string) bool {
	userWiththisEmail := service.repository.FindByEmail(email)
	if len(userWiththisEmail) != 0 {
		return false
	}
	hashedPassword, err := database.HashPassword(password)
	if err != nil {
		return false
	}
	newUser := schemas.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	service.repository.Save(newUser)
	return true
}
