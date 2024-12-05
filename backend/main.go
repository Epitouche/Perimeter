package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"area/api"
	"area/controller"
	"area/database"
	"area/docs"
	"area/repository"
	"area/schemas"
	"area/service"
)

// @Summary ping example
// @Description do ping to check if the server is running
// @Tags ping route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Router /ping [get]
func ping(router *gin.RouterGroup) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "pong",
		})
	})
}

func setupRouter() *gin.Engine {
	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		panic("BACKEND_PORT is not set")
	}

	router := gin.Default()
	router.Use(cors.Default())

	docs.SwaggerInfo.Title = "Area API"
	docs.SwaggerInfo.Description = "Area - Automation API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + appPort
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)

	// Database connection
	databaseConnection := database.Connection()

	// Repositories
	githubRepository := repository.NewGithubRepository(databaseConnection)
	gmailRepository := repository.NewGmailRepository(databaseConnection)
	spotifyRepository := repository.NewSpotifyRepository(databaseConnection)
	timerRepository := repository.NewTimerRepository()
	userRepository := repository.NewUserRepository(databaseConnection)
	serviceRepository := repository.NewServiceRepository(databaseConnection)
	actionRepository := repository.NewActionRepository(databaseConnection)
	reactionRepository := repository.NewReactionRepository(databaseConnection)
	areaRepository := repository.NewAreaRepository(databaseConnection)
	tokenRepository := repository.NewTokenRepository(databaseConnection)

	// Services
	githubService := service.NewGithubService(githubRepository)
	gmailService := service.NewGmailService(gmailRepository)
	spotifyService := service.NewSpotifyService(spotifyRepository)
	timerService := service.NewTimerService(timerRepository)
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository, jwtService)
	serviceService := service.NewServiceService(serviceRepository, timerService)
	actionService := service.NewActionService(actionRepository, serviceService)
	reactionService := service.NewReactionService(reactionRepository, serviceService)
	areaService := service.NewAreaService(areaRepository, serviceService, actionService, reactionService, userService)
	tokenService := service.NewTokenService(tokenRepository)

	// Controllers
	spotifyController := controller.NewSpotifyController(
		spotifyService,
		userService,
		tokenService,
		serviceService,
	)
	githubController := controller.NewGithubController(
		githubService,
		userService,
		tokenService,
		serviceService,
	)
	gmailController := controller.NewGmailController(
		gmailService,
		userService,
		tokenService,
		serviceService,
	)
	userController := controller.NewUserController(userService, jwtService)
	serviceController := controller.NewServiceController(
		serviceService,
		actionService,
		reactionService,
	)
	actionController := controller.NewActionController(actionService)
	reactionController := controller.NewReactionController(reactionService)
	areaController := controller.NewAreaController(areaService)
	tokenController := controller.NewTokenController(tokenService)

	// API routes
	serviceAPI := api.NewServiceApi(serviceController)
	api.NewActionApi(actionController)
	api.NewReactionApi(reactionController)
	areaAPI := api.NewAreAPI(areaController)
	api.NewTokenApi(tokenController)

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		area := apiRoutes.Group("/area")
		{
			area.POST("/", func(c *gin.Context) {
				areaAPI.CreateArea(c)
			})
		}

	}
	ping(apiRoutes)
	api.NewUserApi(userController, apiRoutes)
	api.NewSpotifyAPI(spotifyController, apiRoutes)
	api.NewGmailAPI(gmailController, apiRoutes)
	api.NewGithubAPI(githubController, apiRoutes)

	// basic about.json route
	router.GET("/about.json", serviceAPI.AboutJson)

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
	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		panic("BACKEND_PORT is not set")
	}

	err := router.Run(":" + appPort)
	if err != nil {
		panic("Error when running the server")
	}
}
