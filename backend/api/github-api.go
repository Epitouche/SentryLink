package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
)

type GithubApi struct {
	githubTokenController controller.GithubTokenController
}

func NewGithubAPI(githubTokenController controller.GithubTokenController) *GithubApi {
	return &GithubApi{
		githubTokenController: githubTokenController,
	}
}

func (api *GithubApi) RedirectToGithub(ctx *gin.Context, path string) {
	authURL, err := api.githubTokenController.RedirectToGithub(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"github_authentication_url": authURL})
	}
}

func (api *GithubApi) HandleGithubTokenCallback(ctx *gin.Context, path string) {
	github_token, err := api.githubTokenController.HandleGithubTokenCallback(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": github_token})
	}
}

func (api *GithubApi) GetUserInfo(ctx *gin.Context) {
	usetInfo, err := api.githubTokenController.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"id": usetInfo.Id, "name": usetInfo.Name, "login": usetInfo.Login, "email": usetInfo.Email, "avatar_url": usetInfo.AvatarUrl, "html_url": usetInfo.HtmlUrl, "type": usetInfo.Type}})
	}
}
