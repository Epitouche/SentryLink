package main

import (
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
)

type ActionService struct {
	Service string
	Action  string
}

func timerAction(hour int, minute int, c chan ActionService, response ActionService) {
	var dt time.Time
	for {
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

	var (
		// Database connection
		databaseConnection *gorm.DB = database.Connection()

		// Repositories
		linkRepository        repository.LinkRepository        = repository.NewLinkRepository(databaseConnection)
		githubTokenRepository repository.GithubTokenRepository = repository.NewGithubTokenRepository(databaseConnection)
		userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)

		// Services
		jwtService         service.JWTService         = service.NewJWTService()
		linkService        service.LinkService        = service.NewLinkService(linkRepository)
		githubTokenService service.GithubTokenService = service.NewGithubTokenService(githubTokenRepository)
		userService        service.UserService        = service.NewUserService(userRepository, jwtService)

		// Controllers
		linkController        controller.LinkController        = controller.NewLinkController(linkService)
		githubTokenController controller.GithubTokenController = controller.NewGithubTokenController(githubTokenService, userService)
		userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	)

	linkApi := api.NewLinkAPI(linkController)

	userApi := api.NewUserAPI(userController)

	githubApi := api.NewGithubAPI(githubTokenController)

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
	// view request received but not found
	router.NoRoute(func(c *gin.Context) {
		// get the path
		path := c.Request.URL.Path
		// get the method
		method := c.Request.Method
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

func handleAction(mychannel chan ActionService) {
	for {
		x := <-mychannel
		if x.Service == "Timer" {
			println(x.Action)
		} else {
			println("Unknown service")
		}
	}
}

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {

	// Create a channel list
	var mychannel = make([]chan ActionService, 2)
	mychannel[0] = make(chan ActionService)
	mychannel[1] = make(chan ActionService)
	mychannel2 := make(chan ActionService)
	mychannel = append(mychannel, mychannel2)

	dt := time.Now().Local()
	hour := dt.Hour()
	minute := dt.Minute() + 1
	if minute > 59 {
		hour = hour + 1
		minute = 0
	}

	go timerAction(hour, minute, mychannel[0], ActionService{
		Service: "Timer",
		Action:  "say Hello",
	})
	go timerAction(hour, minute, mychannel[2], ActionService{
		Service: "Timer",
		Action:  "say Bolo",
	})
	go timerAction(hour, minute, mychannel[0], ActionService{
		Service: "Timer",
		Action:  "say Toto",
	})

	go handleAction(mychannel[0])
	go handleAction(mychannel[2])

	router := setupRouter()

	// Listen and Server in 0.0.0.0:8000
	err := router.Run(":8080")
	if err != nil {
		panic("Error when running the server")
	}
}
