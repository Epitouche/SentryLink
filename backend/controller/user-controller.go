package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/service"
)

type UserController interface {
	Login(ctx *gin.Context) string
	Registration(ctx *gin.Context) string
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

func (controller *userController) Login(ctx *gin.Context) string {
	var credentials schemas.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.userService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}

func (controller *userController) Registration(ctx *gin.Context) string {
	var credentials schemas.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.userService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
