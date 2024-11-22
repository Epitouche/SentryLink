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
	GetUserInfo(c *gin.Context) (userInfo schemas.GithubUserInfo, err error)
}

type githubTokenController struct {
	service     service.GithubTokenService
	serviceUser service.UserService
}

var validateGithubToken *validator.Validate

func NewGithubTokenController(service service.GithubTokenService, serviceUser service.UserService) GithubTokenController {
	validateGithubToken = validator.New()
	return &githubTokenController{
		service:     service,
		serviceUser: serviceUser,
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

	githubTokenResponse, err := controller.service.AuthGetGithubAccessToken(code, path)
	if err != nil {
		return "", errors.New("unable to get access token because " + err.Error())
	}

	newGithubToken := schemas.GithubToken{
		AccessToken: githubTokenResponse.AccessToken,
		Scope:       githubTokenResponse.Scope,
		TokenType:   githubTokenResponse.TokenType,
	}

	// Save the access token in the database
	tokenId, err := controller.service.SaveToken(newGithubToken)
	userAlreadExists := false
	if err != nil {
		if err.Error() == "token already exists" {
			userAlreadExists = true
		} else {
			return "", errors.New("unable to save token because " + err.Error())
		}
	}
	userInfo, err := controller.service.GetUserInfo(newGithubToken.AccessToken)
	if err != nil {
		return "", errors.New("unable to get user info because " + err.Error())
	}
	newUser := schemas.User{
		Username: userInfo.Login,
		Email:    userInfo.Email,
		GithubId: tokenId,
	}

	if userAlreadExists {
		token, err := controller.serviceUser.Login(newUser)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		token, err := controller.serviceUser.Register(newUser)
		if err != nil {
			return "", err
		}
		return token, nil
	}
}

func (controller *githubTokenController) GetUserInfo(ctx *gin.Context) (userInfo schemas.GithubUserInfo, err error) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	token, err := controller.service.GetTokenById(user.GithubId)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	githubUserInfo, err := controller.service.GetUserInfo(token.AccessToken)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}

	return githubUserInfo, nil
}
