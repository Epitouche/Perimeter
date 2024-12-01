package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"area/api"
	"area/controller"
	"area/database"
	"area/docs"
	"area/middlewares"
	"area/repository"
	"area/schemas"
	"area/service"
)

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

	// Database connection
	databaseConnection := database.Connection()

	// Repositories
	githubTokenRepository := repository.NewGithubTokenRepository(databaseConnection)
	userRepository := repository.NewUserRepository(databaseConnection)
	serviceRepository := repository.NewServiceRepository(databaseConnection)
	actionRepository := repository.NewActionRepository(databaseConnection)
	reactionRepository := repository.NewReactionRepository(databaseConnection)

	// Services
	githubTokenService := service.NewGithubTokenService(githubTokenRepository)
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository, jwtService)
	serviceService := service.NewServiceService(serviceRepository)
	actionService := service.NewActionService(actionRepository, serviceService)
	reactionService := service.NewReactionService(reactionRepository, serviceService)

	// Controllers
	githubTokenController := controller.NewGithubTokenController(githubTokenService, userService)
	userController := controller.NewUserController(userService, jwtService)
	serviceController := controller.NewServiceController(serviceService)
	actionController := controller.NewActionController(actionService)
	reactionController := controller.NewReactionController(reactionService)

	userAPI := api.NewUserApi(userController)

	GithubAPI := api.NewGithubAPI(githubTokenController)

	api.NewServiceApi(serviceController)
	api.NewActionApi(actionController)
	api.NewReactionApi(reactionController)

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		// User Auth
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", userAPI.Login)
			auth.POST("/register", userAPI.Register)
		}

		// Github
		github := apiRoutes.Group("/github")
		{
			github.GET("/auth", func(c *gin.Context) {
				GithubAPI.RedirectToGithub(c, github.BasePath()+"/auth/callback")
			})

			github.GET("/auth/callback", func(c *gin.Context) {
				GithubAPI.HandleGithubTokenCallback(c, github.BasePath()+"/auth/callback")
			})

			githubInfo := github.Group("/info", middlewares.AuthorizeJWT())
			{
				githubInfo.GET("/user", GithubAPI.GetUserInfo)
			}
		}
	}

	// basic about.json route
	router.GET("/about.json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"client": map[string]string{
				"host": ctx.ClientIP(),
			},
			"server": map[string]string{
				"current_time": fmt.Sprintf("%d", time.Now().Unix()),
				"services":     "area",
			},
		})
	})

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

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization.
func main() {
	router := setupRouter()

	// Listen and Server in 0.0.0.0:8000
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		panic("APP_PORT is not set")
	}

	err := router.Run(":" + appPort)
	if err != nil {
		panic("Error when running the server")
	}
}
