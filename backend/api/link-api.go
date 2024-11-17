package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type LinkApi struct {
	userController        controller.UserController
	linkController        controller.LinkController
	githubTokenController controller.GithubTokenController
}

func NewLinkAPI(userController controller.UserController,
	linkController controller.LinkController, githubTokenController controller.GithubTokenController) *LinkApi {
	return &LinkApi{
		userController:        userController,
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
func (api *LinkApi) Login(ctx *gin.Context) {
	token := api.userController.Login(ctx)
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

func (api *LinkApi) Register(ctx *gin.Context) {
	token := api.userController.Register(ctx)
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

// Generate a random CSRF token
func generateCSRFToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (api *LinkApi) RedirectToGithub(c *gin.Context, path string) {

	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GITHUB_CLIENT_ID is not set"})
		return
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "APP_PORT is not set"})
		return
	}
	// Generate the CSRF token
	state, err := generateCSRFToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to generate CSRF token"})
		return
	}

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	c.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the GitHub authorization URL
	redirectURI := "http://localhost:" + appPort + path
	authURL := "https://github.com/login/oauth/authorize" +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state

	// Redirect to GitHub's OAuth page
	c.Redirect(http.StatusFound, authURL)
}

func GetGithubAccessToken(code string, path string) (schemas.GitHubTokenResponse, error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return schemas.GitHubTokenResponse{}, errors.New("GITHUB_CLIENT_ID is not set")
	}
	clientSecret := os.Getenv("GITHUB_SECRET")
	if clientSecret == "" {
		return schemas.GitHubTokenResponse{}, errors.New("GITHUB_SECRET is not set")
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		return schemas.GitHubTokenResponse{}, errors.New("APP_PORT is not set")
	}
	redirectURI := "http://localhost:" + appPort + path

	apiURL := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return schemas.GitHubTokenResponse{}, err
	}
	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30, // Adjust the timeout as needed
	}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.GitHubTokenResponse{}, err
	}
	defer resp.Body.Close()

	var result schemas.GitHubTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return schemas.GitHubTokenResponse{}, err
	}

	return result, nil
}

func (api *LinkApi) HandleGithubTokenCallback(c *gin.Context, path string) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code"})
		return
	}
	state := c.Query("state")

	latestCSRFToken, err := c.Cookie("latestCSRFToken")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing CSRF token"})
		return
	}

	if state != latestCSRFToken {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid CSRF token"})
		return
	}

	githubTokenResponse, err := GetGithubAccessToken(code, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get access token because " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": githubTokenResponse.AccessToken, "state": state})
}
