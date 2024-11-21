package controller

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/service"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
}

type userController struct {
	userService service.UserService
	jWtService  service.JWTService
}

func NewUserController(userService service.UserService,
	jWtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jWtService:  jWtService,
	}
}

func (controller *userController) Login(ctx *gin.Context) (string, error) {
	var credentials schemas.LoginCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", err
	}
	isAuthenticated := controller.userService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true), nil
	}
	return "", errors.New("bad credentials")
}

func (controller *userController) Register(ctx *gin.Context) (string, error) {
	var credentials schemas.RegisterCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", err
	}
	if len(credentials.Username) < 4 {
		return "", errors.New("username must be at least 4 characters long")
	}
	if len(credentials.Password) < 8 {
		return "", errors.New("password must be at least 8 characters long")
	}
	if len(credentials.Email) < 4 {
		return "", errors.New("email must be at least 4 characters long")
	}
	isCreated := controller.userService.Registration(credentials.Username, credentials.Email, credentials.Password)
	if isCreated {
		return controller.jWtService.GenerateToken(credentials.Username, true), nil
	}
	return "", errors.New("email already exists")
}
