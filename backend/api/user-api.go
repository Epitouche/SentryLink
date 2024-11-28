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

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} schemas.JWT
// @Failure 401 {object} schemas.Response
// @Router /auth/token [post]
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
	token, err := api.userController.Register(ctx)
	if err != nil {
		ctx.JSON(http.StatusConflict, &schemas.Response{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, &schemas.JWT{
		Token: token,
	})
}
