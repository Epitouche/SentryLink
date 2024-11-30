package main

import (
	"net/http"
	"os"
	"time"

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
	swaggerui "github.com/Tom-Mendy/SentryLink/toolbox/swaggerUI"
)

type ActionService struct {
	Service string
	Action  string
}

func timerAction(c chan ActionService, active *bool, hour int, minute int, response ActionService) {
	var dt time.Time
	for *active {
		dt = time.Now().Local()
		if dt.Hour() == hour && dt.Minute() == minute {
			println("current time is ", dt.String())
			c <- response // send sum to c
		}
		time.Sleep(30 * time.Second)
	}
}

func setupRouter(deps Dependencies) *gin.Engine {
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

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		// User Auth
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", deps.UserAPI.Login)
			auth.POST("/register", deps.UserAPI.Register)
		}

		// Links
		links := apiRoutes.Group("/links", middlewares.AuthorizeJWT())
		{
			links.GET("", deps.LinkAPI.GetLink)
			links.POST("", deps.LinkAPI.CreateLink)
			links.PUT(":id", deps.LinkAPI.UpdateLink)
			links.DELETE(":id", deps.LinkAPI.DeleteLink)
		}

		// Scrap
		scrap := apiRoutes.Group("/scrap")
		{
			scrap.GET("", deps.ScrapAPI.GetScrappedUrl)
		}

		// Github
		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", func(c *gin.Context) {
				deps.GithubAPI.RedirectToGithub(c, github.BasePath()+"/auth/callback")
			})

			github.GET("/auth/callback", func(c *gin.Context) {
				deps.GithubAPI.HandleGithubTokenCallback(c, github.BasePath()+"/auth/callback")
			})

			githubInfo := github.Group("/info", middlewares.AuthorizeJWT())
			{
				githubInfo.GET("/user", deps.GithubAPI.GetUserInfo)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(http.StatusNotFound, gin.H{"error": "not found", "path": path, "method": method})
	})

	return router
}

type Dependencies struct {
	UserAPI   *api.UserApi
	LinkAPI   *api.LinkApi
	ScrapAPI  *api.ScrapApi
	GithubAPI *api.GithubApi
}

// initDependencies initializes all required dependencies
func initDependencies() Dependencies {
	// Database connection
	databaseConnection := database.Connection()

	// Repositories
	linkRepository := repository.NewLinkRepository(databaseConnection)
	githubTokenRepository := repository.NewGithubTokenRepository(databaseConnection)
	userRepository := repository.NewUserRepository(databaseConnection)
	scrapRepository := repository.NewScrapRepository(databaseConnection)

	// Services
	linkService := service.NewLinkService(linkRepository)
	githubTokenService := service.NewGithubTokenService(githubTokenRepository)
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository, jwtService)
	scrapService := service.NewScrapService(scrapRepository)

	// Controllers
	linkController := controller.NewLinkController(linkService)
	githubTokenController := controller.NewGithubTokenController(githubTokenService, userService)
	userController := controller.NewUserController(userService, jwtService)
	scrapController := controller.NewScrapController(scrapService)

	// APIs
	return Dependencies{
		UserAPI:   api.NewUserAPI(userController),
		LinkAPI:   api.NewLinkAPI(linkController),
		ScrapAPI:  api.NewScrapApi(scrapController),
		GithubAPI: api.NewGithubAPI(githubTokenController),
	}
}

// initRoutes initializes custom routes (e.g., Swagger routes)
func initRoutes(deps Dependencies) {
	var routes = []schemas.Route{
		{
			Path:        "/auth/register",
			Method:      "POST",
			Handler:     deps.UserAPI.Register,
			Description: "Register a new user",
			Product:     []string{"application/json"},
			Tags:        []string{"auth"},
			ParamQueryType: "formData",
			Params: map[string]string{
				"username": "string",
				"email":    "string",
				"password": "string",
			},
			Responses: map[int][]string{
				http.StatusOK: {
					"User registered successfully",
					"schemas.Response",
				},
				http.StatusConflict: {
					"User already exists",
					"schemas.Response",
				},
				http.StatusBadRequest: {
					"Invalid request",
					"schemas.Response",
				},
			},
		},
		{
			Path:        "/auth/login",
			Method:      "POST",
			Handler:     deps.UserAPI.Login,
			Description: "Authenticate a user and provide a JWT to authorize API calls",
			Product:     []string{"application/json"},
			Tags:        []string{"auth"},
			ParamQueryType: "formData",
			Params: map[string]string{
				"username": "string",
				"password": "string",
			},
			Responses: map[int][]string{
				http.StatusOK: {
					"JWT",
					"schemas.JWT",
				},
				http.StatusUnauthorized: {
					"Unauthorized",
					"schemas.Response",
				},
			},
		},
	}
	swaggerui.ImpactSwaggerFiles(routes)
}


func handleAction(mychannel chan ActionService, active *bool) {
	for {
		x := <-mychannel
		if x.Service == "Timer" {
			println(x.Action)
			*active = true
		} else {
			println("Unknown service")
		}
	}
}

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization.
func main() {

	deps := initDependencies()

	initRoutes(deps)

	// Create a channel list
	allChannel := make([]chan ActionService, 2)
	allChannel[0] = make(chan ActionService)
	allChannel[1] = make(chan ActionService)
	newChannel := make(chan ActionService)
	allChannel = append(allChannel, newChannel)

	dt := time.Now().Local()
	hour := dt.Hour()
	minute := dt.Minute() + 1
	if minute > 59 {
		hour = hour + 1
		minute = 0
	}

	active := true

	go timerAction(allChannel[0], &active, hour, minute, ActionService{
		Service: "Timer",
		Action:  "say Hello",
	})

	go timerAction(allChannel[0], &active, hour, minute, ActionService{
		Service: "Timer",
		Action:  "say Bolo",
	})

	go timerAction(allChannel[0], &active, hour, minute, ActionService{
		Service: "Timer",
		Action:  "say Toto",
	})

	go handleAction(allChannel[0], &active)

	router := setupRouter(deps)

	// Listen and Server in 0.0.0.0:8000
	err := router.Run(":8080")
	if err != nil {
		panic("Error when running the server")
	}
}
