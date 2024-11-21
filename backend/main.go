package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
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

// ExtractLinks extracts all the links from an HTML page
func ExtractLinks(pageURL string) ([]string, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch the page: %s", resp.Status)
	}

	links := []string{}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						// Resolve relative URLs
						parsedURL, err := url.Parse(attr.Val)
						if err != nil {
							continue
						}
						base, err := url.Parse(pageURL)
						if err != nil {
							continue
						}
						links = append(links, base.ResolveReference(parsedURL).String())
					}
				}
			}
		}
	}

	return links, nil
}

func setupRouter() *gin.Engine {

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		panic("APP_PORT is not set")
	}

	docs.SwaggerInfo.Title = "SentryLink API"
	docs.SwaggerInfo.Description = "SentryLink - Crawler API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + appPort
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.Default()

	// Ping test
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "pong",
		})
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

	linkApi := api.NewLinkAPI(linkController)

	userApi := api.NewUserAPI(userController)

	githubApi := api.NewGithubAPI(githubTokenController)

	// scrapApi := api

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		// User Auth
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", userApi.Login)
			auth.POST("/register", userApi.Register)
		}

		// Links
		links := apiRoutes.Group("/links", middlewares.AuthorizeJWT())
		{
			links.GET("", linkApi.GetLink)
			links.POST("", linkApi.CreateLink)
			links.PUT(":id", linkApi.UpdateLink)
			links.DELETE(":id", linkApi.DeleteLink)
		}

		// Scrap
		// scrap := apiRoutes.Group("/scrap", middlewares.AuthorizeJWT())
		// {
		// 	scrap.GET("", )
		// }

		// Github
		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", func(c *gin.Context) {
				githubApi.RedirectToGithub(c, github.BasePath()+"/auth/callback")
			})

			github.GET("/auth/callback", func(c *gin.Context) {
				githubApi.HandleGithubTokenCallback(c, github.BasePath()+"/auth/callback")
			})
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// view request received but not found
	router.NoRoute(func(c *gin.Context) {
		// get the path
		path := c.Request.URL.Path
		// get the method
		method := c.Request.Method
		print("\n\n" + method + " " + path + "\n\n\n")
		c.JSON(http.StatusNotFound, gin.H{"error": "not found", "path": path, "method": method})
	})

	router.GET("/scrap", func(c *gin.Context) {
		pageURL := c.Query("url")
		if pageURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url parameter is required"})
			return
		}

		// Ensure the URL starts with http:// or https://
		if !strings.HasPrefix(pageURL, "http://") && !strings.HasPrefix(pageURL, "https://") {
			pageURL = "http://" + pageURL
		}

		links, err := ExtractLinks(pageURL)
		if err != nil {
			log.Println("Error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"links": links})
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
