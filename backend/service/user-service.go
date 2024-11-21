package service

import (
	"errors"

	"github.com/Tom-Mendy/SentryLink/database"
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type UserService interface {
	Login(username string, password string) (string, error)
	Register(newUser schemas.User) (string, error)
}

type userService struct {
	authorizedUsername string
	authorizedPassword string
	repository         repository.UserRepository
	serviceJWT         JWTService
}

func NewUserService(userRepository repository.UserRepository, serviceJWT JWTService) UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository:         userRepository,
		serviceJWT:         serviceJWT,
	}
}

func (service *userService) Login(username string, password string) (string, error) {
	userWiththisUserName := service.repository.FindByUserName(username)
	if len(userWiththisUserName) == 0 {
		return "", errors.New("invalid credentials")
	}
	for _, user := range userWiththisUserName {
		if database.DoPasswordsMatch(user.Password, password) {
			return service.serviceJWT.GenerateToken(username, true), nil
		}
	}
	return "", errors.New("invalid credentials")
}

func (service *userService) Register(newUser schemas.User) (string, error) {
	userWiththisEmail := service.repository.FindByEmail(newUser.Email)
	if len(userWiththisEmail) != 0 {
		return "", errors.New("email already in use")
	}

	if newUser.Password != "" {
		hashedPassword, err := database.HashPassword(newUser.Password)
		if err != nil {
			return "", errors.New("error while hashing the password")
		}
		newUser.Password = hashedPassword
	}

	service.repository.Save(newUser)
	return service.serviceJWT.GenerateToken(newUser.Username, true), nil
}
