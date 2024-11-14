package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type LinkApi struct {
	loginController       controller.LoginController
	linkController        controller.LinkController
	githubTokenController controller.GithubTokenController
}

func NewLinkAPI(loginController controller.LoginController,
	linkController controller.LinkController, githubTokenController controller.GithubTokenController) *LinkApi {
	return &LinkApi{
		loginController:       loginController,
		linkController:        linkController,
		githubTokenController: githubTokenController,
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
func (api *LinkApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &schemas.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &schemas.Response{
			Message: "Not Authorized",
		})
	}
}

// GetLink godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} api.LinkApi
// @Failure 401 {object} schemas.Response
// @Router /videos [get]
func (api *LinkApi) GetLink(ctx *gin.Context) {
	ctx.JSON(200, api.linkController.FindAll())
}

// CreateLink godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body api.LinkApi true "Create video"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 401 {object} schemas.Response
// @Router /videos [post]
func (api *LinkApi) CreateLink(ctx *gin.Context) {
	err := api.linkController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "Success!",
		})
	}
}

// UpdateLink godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Param video body api.LinkApi true "Update video"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 401 {object} schemas.Response
// @Router /videos/{id} [put]
func (api *LinkApi) UpdateLink(ctx *gin.Context) {
	err := api.linkController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "Success!",
		})
	}
}

// DeleteLink godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 401 {object} schemas.Response
// @Router /videos/{id} [delete]
func (api *LinkApi) DeleteLink(ctx *gin.Context) {
	err := api.linkController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "Success!",
		})
	}
}
