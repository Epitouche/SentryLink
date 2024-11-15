package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Tom-Mendy/SentryLink/api"
	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/database"
	"github.com/Tom-Mendy/SentryLink/docs"
	"github.com/Tom-Mendy/SentryLink/middlewares"
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/service"
)

// Generate a random CSRF token
func generateCSRFToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func redirectToGithub(c *gin.Context) {

	clientID := os.Getenv("GITHUB_CLIENT_ID")
	appPort := os.Getenv("APP_PORT")
	if clientID == "" || appPort == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GITHUB_CLIENT_ID or APP_PORT is not set"})
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
	redirectURI := "http://localhost:" + appPort + "/github/auth/callback"
	authURL := "https://github.com/user/oauth/authorize" +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state

	// Redirect to GitHub's OAuth page
	c.Redirect(http.StatusFound, authURL)
}

func getGithubAccessToken(code string) (schemas.GitHubTokenResponse, error) {
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
	redirectURI := "http://localhost:" + appPort + "/github/auth/callback"

	apiURL := "https://github.com/user/oauth/access_token"

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

func setupRouter() *gin.Engine {

	docs.SwaggerInfo.Title = "SentryLink API"
	docs.SwaggerInfo.Description = "SentryLink - Crawler API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("APP_PORT")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	var (
		// Database connection
		databaseConnection *gorm.DB = database.Connection()

		// Repositories
		linkRepository        repository.LinkRepository        = repository.NewLinkRepository(databaseConnection)
		githubTokenRepository repository.GithubTokenRepository = repository.NewGithubTokenRepository(databaseConnection)
		userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)

		// Services
		linkService        service.LinkService        = service.NewLinkService(linkRepository)
		githubTokenService service.GithubTokenService = service.NewGithubTokenService(githubTokenRepository)
		userService        service.UserService        = service.NewUserService(userRepository)
		jwtService         service.JWTService         = service.NewJWTService()

		// Controllers
		linkController        controller.LinkController        = controller.NewLinkController(linkService)
		githubTokenController controller.GithubTokenController = controller.NewGithubTokenController(githubTokenService)
		userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	)

	linkApi := api.NewLinkAPI(userController, linkController, githubTokenController)

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", linkApi.Login)
			auth.POST("/register", linkApi.Register)
		}

		links := apiRoutes.Group("/links", middlewares.AuthorizeJWT())
		{
			links.GET("", linkApi.GetLink)
			links.POST("", linkApi.CreateLink)
			links.PUT(":id", linkApi.UpdateLink)
			links.DELETE(":id", linkApi.DeleteLink)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/github/auth", redirectToGithub)

	router.GET("/github/auth/callback", func(c *gin.Context) {
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

		githubTokenResponse, err := getGithubAccessToken(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get access token because " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": githubTokenResponse.AccessToken, "state": state})
	})

	// view request received but not found
	router.NoRoute(func(c *gin.Context) {
		// get the path
		path := c.Request.URL.Path
		// get the method
		method := c.Request.Method
		print("\n\n" + method + " " + path + "\n\n\n")
		c.JSON(http.StatusNotFound, gin.H{"error": "not found", "path": path, "method": method})
	})

	return router
}

func init() {
	// err := .Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }
}

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	router := setupRouter()

	// Listen and Server in 0.0.0.0:8000
	err := router.Run(":8000")
	if err != nil {
		panic("Error when running the server")
	}
}
