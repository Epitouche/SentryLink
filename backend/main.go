package main

import (
	"fmt"
	"net/http"
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
		scrap := apiRoutes.Group("/scrap")
		{
			scrap.GET("", scrapApi.GetScrappedUrl)
		}

		// Github
		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", func(c *gin.Context) {
				githubApi.RedirectToGithub(c, github.BasePath()+"/auth/callback")
			})

			github.GET("/auth/callback", func(c *gin.Context) {
				githubApi.HandleGithubTokenCallback(c, github.BasePath()+"/auth/callback")
			})

			githubInfo := github.Group("/info", middlewares.AuthorizeJWT())
			{
				githubInfo.GET("/user", githubApi.GetUserInfo)
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

// type Dependencies struct {
// 	UserAPI   *api.UserApi
// 	LinkAPI   *api.LinkApi
// 	ScrapAPI  *api.ScrapApi
// 	GithubAPI *api.GithubApi
// }

// // initDependencies initializes all required dependencies
// func initDependencies() Dependencies {
// 	// Database connection
// 	databaseConnection := database.Connection()

// 	// Repositories
// 	linkRepository := repository.NewLinkRepository(databaseConnection)
// 	githubTokenRepository := repository.NewGithubTokenRepository(databaseConnection)
// 	userRepository := repository.NewUserRepository(databaseConnection)
// 	scrapRepository := repository.NewScrapRepository(databaseConnection)

// 	// Services
// 	linkService := service.NewLinkService(linkRepository)
// 	githubTokenService := service.NewGithubTokenService(githubTokenRepository)
// 	jwtService := service.NewJWTService()
// 	userService := service.NewUserService(userRepository, jwtService)
// 	scrapService := service.NewScrapService(scrapRepository)

// 	// Controllers
// 	linkController := controller.NewLinkController(linkService)
// 	githubTokenController := controller.NewGithubTokenController(githubTokenService, userService)
// 	userController := controller.NewUserController(userService, jwtService)
// 	scrapController := controller.NewScrapController(scrapService)

// 	// APIs
// 	return Dependencies{
// 		UserAPI:   api.NewUserAPI(userController),
// 		LinkAPI:   api.NewLinkAPI(linkController),
// 		ScrapAPI:  api.NewScrapApi(scrapController),
// 		GithubAPI: api.NewGithubAPI(githubTokenController),
// 	}
// }

// // initRoutes initializes custom routes (e.g., Swagger routes)
// func initRoutes(deps Dependencies) {

// }

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

var (
	// Database connection
	databaseConnection *gorm.DB = database.Connection()
	// Repositories
	linkRepository        repository.LinkRepository        = repository.NewLinkRepository(databaseConnection)
	githubTokenRepository repository.GithubTokenRepository = repository.NewGithubTokenRepository(databaseConnection)
	userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)
	scrapRepository       repository.ScrapRepository       = repository.NewScrapRepository(databaseConnection)
	// Services
	linkService        service.LinkService        = service.NewLinkService(linkRepository)
	githubTokenService service.GithubTokenService = service.NewGithubTokenService(githubTokenRepository)
	jwtService         service.JWTService         = service.NewJWTService()
	userService        service.UserService        = service.NewUserService(userRepository, jwtService)
	scrapService       service.ScrapService       = service.NewScrapService(scrapRepository)
	// Controllers
	linkController        controller.LinkController        = controller.NewLinkController(linkService)
	githubTokenController controller.GithubTokenController = controller.NewGithubTokenController(githubTokenService, userService)
	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	scrapController       controller.ScrapController       = controller.NewScrapController(scrapService)
)

var (
	linkApi   *api.LinkApi   = api.NewLinkAPI(linkController)
	githubApi *api.GithubApi = api.NewGithubAPI(githubTokenController)
	userApi   *api.UserApi   = api.NewUserAPI(userController)
	scrapApi  *api.ScrapApi  = api.NewScrapApi(scrapController)
)

func ensureSwaggerDocsUpdated() {
	var routes = []schemas.Route{
		{
			Path:           "/auth/register",
			Method:         "POST",
			Handler:        userApi.Register,
			Description:    "Register a new user",
			Product:        []string{"application/json"},
			Tags:           []string{"auth"},
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
			Path:           "/auth/login",
			Method:         "POST",
			Handler:        userApi.Login,
			Description:    "Authenticate a user and provide a JWT to authorize API calls",
			Product:        []string{"application/json"},
			Tags:           []string{"auth"},
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
		{
			Path:           "/scrap",
			Method:         "GET",
			Handler:        scrapApi.GetScrappedUrl,
			Description:    "Scrap an url and return all the links",
			Product:        []string{"application/json"},
			Tags:           []string{"scrap"},
			ParamQueryType: "query",
			Params: map[string]string{
				"linkToScrap": "string",
			},
			Responses: map[int][]string{
				http.StatusOK: {
					"Scrapped links",
					"schemas.Response",
				},
				http.StatusUnauthorized: {
					"Unauthorized",
					"schemas.Response",
				},
			},
		},
		// {
		// 	Path:           "/toto",
		// 	Method:         "GET",
		// 	Handler:        userApi.Register,
		// 	Description:    "toot",
		// 	Product:        []string{"application/json"},
		// 	Tags:           []string{"apart"},
		// 	ParamQueryType: "query",
		// 	Params: map[string]string{
		// 		"linkToScrap": "string",
		// 	},
		// 	Responses: map[int][]string{
		// 		http.StatusOK: {
		// 			"Scrapped links",
		// 			"schemas.Response",
		// 		},
		// 		http.StatusUnauthorized: {
		// 			"Unauthorized",
		// 			"schemas.Response",
		// 		},
		// 	},
		// },
	}
	filePath := "docs/swagger.json"

	pathsOfRoutesWanted := []string{}
	existingRoutes := swaggerui.ExtractExistingRoutes(filePath)
	for _, route := range routes {
		pathsOfRoutesWanted = append(pathsOfRoutesWanted, route.Path)
	}

	fmt.Printf("pathsOfRoutesWanted: %++v\n", pathsOfRoutesWanted)
	routesToRemove := swaggerui.FindRoutesToRemove(existingRoutes, pathsOfRoutesWanted)
	fmt.Printf("routesToRemove: %++v\n", routesToRemove)

	if len(routesToRemove) > 0 || (len(pathsOfRoutesWanted) != len(existingRoutes)) {
		swaggerui.ImpactSwaggerFiles(routes, routesToRemove)
	}

	println("Updated docs/docs.go")
}

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization.
func main() {
	ensureSwaggerDocsUpdated()

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

	router := setupRouter()

	// Listen and Server in 0.0.0.0:8000
	err := router.Run(":7979")
	if err != nil {
		panic("Error when running the server")
	}
}
