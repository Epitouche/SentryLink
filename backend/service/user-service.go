package service

import (
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
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}

func (service *userService) Registration(username string, email string, password string) bool {
	userWiththis := service.repository.FindByEmail(email)
	if userWiththis != nil {
		return false
	}
	newUser := schemas.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	service.repository.Save(newUser)
	return true
}
