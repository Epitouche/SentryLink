package controller

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/service"
	"github.com/Tom-Mendy/SentryLink/tools"
)

type GithubTokenController interface {
	RedirectToGithub(ctx *gin.Context, path string) (string, error)
	HandleGithubTokenCallback(c *gin.Context, path string) (string, error)
}

type githubTokenController struct {
	service service.GithubTokenService
}

var validateGithubToken *validator.Validate

func NewGithubTokenController(service service.GithubTokenService) GithubTokenController {
	validateGithubToken = validator.New()
	return &githubTokenController{
		service: service,
	}
}

func (controller *githubTokenController) RedirectToGithub(ctx *gin.Context, path string) (string, error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return "", errors.New("GITHUB_CLIENT_ID is not set")
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		return "", errors.New("APP_PORT is not set")
	}
	// Generate the CSRF token
	state, err := tools.GenerateCSRFToken()
	if err != nil {
		return "", errors.New("unable to generate CSRF token")
	}

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the GitHub authorization URL
	redirectURI := "http://localhost:" + appPort + path
	authURL := "https://github.com/login/oauth/authorize" +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}

func (controller *githubTokenController) HandleGithubTokenCallback(c *gin.Context, path string) (string, error) {
	code := c.Query("code")
	if code == "" {
		return "", errors.New("missing code")
	}
	state := c.Query("state")

	latestCSRFToken, err := c.Cookie("latestCSRFToken")
	if err != nil {
		return "", errors.New("missing CSRF token")
	}

	if state != latestCSRFToken {
		return "", errors.New("invalid CSRF token")
	}

	githubTokenResponse, err := controller.service.GetGithubAccessToken(code, path)
	if err != nil {
		return "", errors.New("unable to get access token because " + err.Error())
	}

	newGithubToken := schemas.GithubToken{
		AccessToken: githubTokenResponse.AccessToken,
		Scope:       githubTokenResponse.Scope,
		TokenType:   githubTokenResponse.TokenType,
	}

	// Save the access token in the database
	controller.service.SaveToken(newGithubToken)
	userInfo, err := controller.service.GetUserInfo(newGithubToken)
	if err != nil {
		return "", errors.New("unable to get user info because " + err.Error())
	}
	userInfo.AvatarUrl = ""

	return githubTokenResponse.AccessToken, nil
}
