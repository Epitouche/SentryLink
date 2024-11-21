package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/tools"
)

type GithubApi struct {
	githubTokenController controller.GithubTokenController
}

func NewGithubAPI(githubTokenController controller.GithubTokenController) *GithubApi {
	return &GithubApi{
		githubTokenController: githubTokenController,
	}
}

func (api *GithubApi) RedirectToGithub(c *gin.Context, path string) {

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
	state, err := tools.GenerateCSRFToken()
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
	c.JSON(http.StatusOK, gin.H{"github_authentication_url": authURL})
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

func (api *GithubApi) HandleGithubTokenCallback(c *gin.Context, path string) {
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

	// Save the access token in the database

	c.JSON(http.StatusOK, gin.H{"access_token": githubTokenResponse.AccessToken, "state": state})
}
