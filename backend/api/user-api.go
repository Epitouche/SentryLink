package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type UserApi struct {
	userController controller.UserController
}

func NewUserAPI(userController controller.UserController) *UserApi {
	return &UserApi{
		userController: userController,
	}
}

func (api *UserApi) Login(ctx *gin.Context) {
	token, err := api.userController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &schemas.Response{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &schemas.JWT{
		Token: token,
	})
}

func (api *UserApi) Register(ctx *gin.Context) {
	_, err := api.userController.Register(ctx)
	if err != nil {
		ctx.JSON(http.StatusConflict, &schemas.Response{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &schemas.Response{
		Message: "User registered successfully",
	})
}
