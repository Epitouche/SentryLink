package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Tom-Mendy/SentryLink/api"
	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/docs"
	"github.com/Tom-Mendy/SentryLink/middlewares"
	"github.com/Tom-Mendy/SentryLink/repository"
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
	if clientID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GITHUB_CLIENT_ID is not set"})
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
	redirectURI := "http://localhost:" + os.Getenv("APP_PORT") + "/integrations/github/oauth2/callback"
	authURL := "https://github.com/login/oauth/authorize" +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state

	// Redirect to GitHub's OAuth page
	c.Redirect(http.StatusFound, authURL)
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
		linkRepository repository.LinkRepository = repository.NewLinkRepository()
		linkService    service.LinkService       = service.New(linkRepository)
		loginService   service.LoginService      = service.NewLoginService()
		jwtService     service.JWTService        = service.NewJWTService()

		linkController  controller.LinkController  = controller.New(linkService)
		loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
	)

	linkApi := api.NewLinkAPI(loginController, linkController)

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", linkApi.Authenticate)
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

	router.GET("/auth/github", redirectToGithub)

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
