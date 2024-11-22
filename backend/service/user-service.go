package service

import (
	"errors"
	"fmt"

	"github.com/Tom-Mendy/SentryLink/database"
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type UserService interface {
	Login(user schemas.User) (JWTtoken string, err error)
	Register(newUser schemas.User) (JWTtoken string, err error)
	GetUserInfo(token string) (userInfo schemas.User, err error)
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

func (service *userService) Login(newUser schemas.User) (JWTtoken string, err error) {
	userWiththisUserName := service.repository.FindByUserName(newUser.Username)
	if len(userWiththisUserName) == 0 {
		return "", errors.New("invalid credentials")
	}
	// regular user
	for _, user := range userWiththisUserName {
		if database.DoPasswordsMatch(user.Password, newUser.Password) {
			return service.serviceJWT.GenerateToken(fmt.Sprint(user.Id), user.Username, false), nil
		}
	}

	// Oauth2.0 user
	for _, user := range userWiththisUserName {
		if user.Email == newUser.Email {
			if newUser.GithubId != 0 {
				return service.serviceJWT.GenerateToken(fmt.Sprint(user.Id), user.Username, false), nil
			}
		}
	}

	return "", errors.New("invalid credentials")
}

func (service *userService) Register(newUser schemas.User) (JWTtoken string, err error) {
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
	return service.serviceJWT.GenerateToken(fmt.Sprint(newUser.Id), newUser.Username, false), nil
}

func (service *userService) GetUserInfo(token string) (userInfo schemas.User, err error) {
	userId, err := service.serviceJWT.GetUserIdfromJWTToken(token)
	if err != nil {
		return schemas.User{}, err
	}
	userInfo = service.repository.FindById(userId)
	return userInfo, nil
}
