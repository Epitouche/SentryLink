package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
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

// var routes = []schemas.Route{
// 	{
// 		Path: "/auth/register",
// 		Method: "POST",
// 		Handler: userApi.Register,
// 		Description: "Register a new user",
//		product: []string{"application/json"},
// 		Tags: []string{"auth"},
//	 	ParamType: "formData",
// 		Params: map[string]string{
// 			"username": "string",
// 			"email": "string",
// 			"password": "string",
// 		},
// 		Responses: map[int]string{
// 			http.StatusOK: "User registered successfully",
// 			http.StatusConflict: "User already exists",
// 			http.StatusBadRequest: "Invalid request",
// 		},
// 	},
// }


func impactSwaggerFiles(routes []schemas.Route) {
	var filePathOfFiles = []string{
		"docs/docs.go",
		"docs/swagger.json",
		"docs/swagger.yaml",
	}
	for _, route := range routes {
		for _, file := range filePathOfFiles {
			processFile(file, route)
		}
	}
}

func isJSONFile(filePath string) bool {
	return len(filePath) > 5 && filePath[len(filePath)-5:] == ".json"
}

func isYAMLFile(filePath string) bool {
	return len(filePath) > 5 && filePath[len(filePath)-5:] == ".yaml"
}

func isGOFile(filePath string) bool {
	return len(filePath) > 3 && filePath[len(filePath)-3:] == ".go"
}

func buildParameters(params map[string]string, paramType string) []map[string]interface{} {
	var parameters []map[string]interface{}
	for name, typ := range params {
		parameters = append(parameters, map[string]interface{}{
			"type":        typ,
			"description": typ,
			"name":        name,
			"in":          paramType,
			"required":    true,
		})
	}
	return parameters
}

func buildResponses(responses map[int][]string) map[string]interface{} {
	responseMap := make(map[string]interface{})
	for code, descAndSchema := range responses {
		description := descAndSchema[0]
		schemaRef := descAndSchema[1]

		responseMap[fmt.Sprintf("%d", code)] = map[string]interface{}{
			"description": description,
			"schema": map[string]interface{}{
				"$ref": fmt.Sprintf("#/definitions/%s", schemaRef),
			},
		}
	}
	return responseMap
}

func buildRouteEntry(route schemas.Route) map[string]interface{} {
	return map[string]interface{}{
		route.Method: map[string]interface{}{
			"tags":        route.Tags,
			"description": route.Description,
			"produces":    route.Product,
			"parameters":  buildParameters(route.Params, route.ParamQueryType),
			"responses":   buildResponses(route.Responses),
		},
	}
}

func updateDocTemplate(content string, route schemas.Route) string {
	startIndex := strings.Index(content, `"paths": {`)
	if startIndex == -1 {
		fmt.Println("Could not find paths section in docTemplate")
		return content
	}

	endIndex := strings.Index(content[startIndex:], `}`)
	if endIndex == -1 {
		fmt.Println("Could not find end of paths section in docTemplate")
		return content
	}
	endIndex += startIndex

	pathsContent := content[startIndex:endIndex+1]
	var paths map[string]interface{}
	err := json.Unmarshal([]byte(pathsContent), &paths)
	if err != nil {
		fmt.Printf("Error unmarshalling docTemplate paths: %v\n", err)
		return content
	}

	paths["paths"].(map[string]interface{})[route.Path] = buildRouteEntry(route)
	updatedPaths, err := json.MarshalIndent(paths, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling updated paths: %v\n", err)
		return content
	}

	return content[:startIndex] + string(updatedPaths) + content[endIndex+1:]
}

func processFile(filePath string, route schemas.Route) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filePath, err)
		return
	}

	var paths map[string]interface{}

	// if isGOFile(filePath) {
	// 	updatedContent := updateDocTemplate(string(fileData), route)
	// 	err := os.WriteFile(filePath, []byte(updatedContent), 0644)
	// 	if err != nil {
	// 		fmt.Printf("Error writing to Go file %s: %v\n", filePath, err)
	// 	} else {
	// 		fmt.Printf("Route added successfully to %s\n", filePath)
	// 	}
	// 	return
	if isJSONFile(filePath) {
		err = json.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", filePath, err)
			return
		}
	} else if isYAMLFile(filePath) {
		err = yaml.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling YAML file %s: %s\n", filePath, err)
			return
		}
	} else {
		fmt.Printf("Unsupported file type %s\n", filePath)
		return
	}

	if paths == nil {
		paths = make(map[string]interface{})
	}
	if _, ok := paths["paths"]; !ok {
		paths["paths"] = make(map[string]interface{})
	}
	pathsMap := paths["paths"].(map[string]interface{})
	pathsMap[route.Path] = buildRouteEntry(route)

	if isJSONFile(filePath) {
		updatedJSON, err := json.MarshalIndent(paths, "", "  ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filePath, err)
			return
		}

		err = os.WriteFile(filePath, updatedJSON, 0644)
		if err != nil {
			fmt.Printf("Error writing JSON to file %s: %v\n", filePath, err)
			return
		}
	} else if isYAMLFile(filePath) {
		updatedYAML, err := yaml.Marshal(paths)
		if err != nil {
			fmt.Printf("Error serializing YAML for file %s: %v\n", filePath, err)
			return
		}

		err = os.WriteFile(filePath, updatedYAML, 0644)
		if err != nil {
			fmt.Printf("Error writing YAML to file %s: %v\n", filePath, err)
			return
		}
	}

	fmt.Printf("Route added successfully to %s\n", filePath)
}


//! Need to change the location of the function later
func GenerateSwaggerDocs(routes []schemas.Route) {
	for _, route := range routes {
		fmt.Println("// " + route.Description)
		fmt.Printf("// @Tags %s\n", strings.Join(route.Tags, ","))
		if route.Params != nil {
			for param, desc := range route.Params {
				fmt.Printf("// @Param %s query string true \"%s\"\n", param, desc)
			}
		}
		for status, desc := range route.Responses {
			fmt.Printf("// @Success %d {object} map[string]interface{} \"%s\"\n", status, desc)
		}
		fmt.Printf("// @Router %s [%s]\n\n", route.Path, strings.ToLower(route.Method))
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

	// var (
	// 	// Database connection
	// 	databaseConnection *gorm.DB = database.Connection()

	// 	// Repositories
	// 	linkRepository        repository.LinkRepository        = repository.NewLinkRepository(databaseConnection)
	// 	githubTokenRepository repository.GithubTokenRepository = repository.NewGithubTokenRepository(databaseConnection)
	// 	userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)
	// 	scrapRepository	   repository.ScrapRepository       = repository.NewScrapRepository(databaseConnection)

	// 	// Services
	// 	linkService        service.LinkService        = service.NewLinkService(linkRepository)
	// 	githubTokenService service.GithubTokenService = service.NewGithubTokenService(githubTokenRepository)
	// 	userService        service.UserService        = service.NewUserService(userRepository)
	// 	jwtService         service.JWTService         = service.NewJWTService()
	// 	scrapService	   service.ScrapService       = service.NewScrapService(scrapRepository)

	// 	// Controllers
	// 	linkController        controller.LinkController        = controller.NewLinkController(linkService)
	// 	githubTokenController controller.GithubTokenController = controller.NewGithubTokenController(githubTokenService)
	// 	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	// 	scrapController       controller.ScrapController       = controller.NewScrapController(scrapService)
	// )

	// linkApi := api.NewLinkAPI(linkController)

	// userApi := api.NewUserAPI(userController)

	// githubApi := api.NewGithubAPI(githubTokenController)

	// scrapApi := api.NewScrapApi(scrapController)

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

	// var routes = []schemas.Route{
	// 	{
	// 		Path: "/auth/register",
	// 		Method: "POST",
	// 		Handler: userApi.Register,
	// 		Description: "Register a new user",
	// 		Product: []string{"application/json"},
	// 		Tags: []string{"auth"},
	// 		ParamQueryType: "formData",
	// 		Params: map[string]string{
	// 			"username": "string",
	// 			"email": "string",
	// 			"password": "string",
	// 		},
	// 		Responses: map[int][]string{
	// 			http.StatusOK: {
	// 				"User registered successfully",
	// 				"schemas.Response",
	// 			},
	// 			http.StatusConflict: {
	// 				"User already exists",
	// 				"schemas.Response",
	// 			},
	// 			http.StatusBadRequest: {
	// 				"Invalid request",
	// 				"schemas.Response",
	// 			},
	// 		},
	// 	},
	// }
	// // GenerateSwaggerDocs(routes)
	// impactSwaggerFiles(routes)

	return router
}

var (
	// Database connection
	databaseConnection *gorm.DB = database.Connection()

	// Repositories
	linkRepository        repository.LinkRepository        = repository.NewLinkRepository(databaseConnection)
	githubTokenRepository repository.GithubTokenRepository = repository.NewGithubTokenRepository(databaseConnection)
	userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)
	scrapRepository	   repository.ScrapRepository       = repository.NewScrapRepository(databaseConnection)

	// Services
	linkService        service.LinkService        = service.NewLinkService(linkRepository)
	githubTokenService service.GithubTokenService = service.NewGithubTokenService(githubTokenRepository)
	userService        service.UserService        = service.NewUserService(userRepository)
	jwtService         service.JWTService         = service.NewJWTService()
	scrapService	   service.ScrapService       = service.NewScrapService(scrapRepository)

	// Controllers
	linkController        controller.LinkController        = controller.NewLinkController(linkService)
	githubTokenController controller.GithubTokenController = controller.NewGithubTokenController(githubTokenService)
	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	scrapController       controller.ScrapController       = controller.NewScrapController(scrapService)
)

var (
	linkApi = api.NewLinkAPI(linkController)
	userApi = api.NewUserAPI(userController)
	githubApi = api.NewGithubAPI(githubTokenController)
	scrapApi = api.NewScrapApi(scrapController)
)

func init() {

	var routes = []schemas.Route{
		{
			Path: "/auth/register",
			Method: "POST",
			Handler: userApi.Register,
			Description: "Register a new user",
			Product: []string{"application/json"},
			Tags: []string{"auth"},
			ParamQueryType: "formData",
			Params: map[string]string{
				"username": "string",
				"email": "string",
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
			Path: "/auth/login",
			Method: "POST",
			Handler: userApi.Login,
			Description: "Authenticate a user and provide a JWT to authorize API calls",
			Product: []string{"application/json"},
			Tags: []string{"auth"},
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
	// GenerateSwaggerDocs(routes)
	impactSwaggerFiles(routes)

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
