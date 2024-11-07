package main

import (
	"net/http"

	"github.com/Tom-Mendy/SentryLink/api"
	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/docs"
	"github.com/Tom-Mendy/SentryLink/middlewares"
	"github.com/Tom-Mendy/SentryLink/repository"
	"github.com/Tom-Mendy/SentryLink/service"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
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

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8000
	err := router.Run(":8000")
	if err != nil {
		panic("Error when running the server")
	}
}
