package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Epitouche/Perimeter/api"
	"github.com/Epitouche/Perimeter/controller"
	"github.com/Epitouche/Perimeter/database"
	"github.com/Epitouche/Perimeter/docs"
	"github.com/Epitouche/Perimeter/repository"
	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/service"
)

// ping godoc
//
//	@Summary		ping example
//	@Description	do ping to check if the server is running
//	@Tags			ping route
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.Response
//	@Router			/ping [get]
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
	dropboxRepository := repository.NewDropboxRepository(databaseConnection)
	timerRepository := repository.NewTimerRepository(databaseConnection)
	openweathermapRepository := repository.NewOpenweathermapRepository(databaseConnection)
	userRepository := repository.NewUserRepository(databaseConnection)
	serviceRepository := repository.NewServiceRepository(databaseConnection)
	actionRepository := repository.NewActionRepository(databaseConnection)
	reactionRepository := repository.NewReactionRepository(databaseConnection)
	areaRepository := repository.NewAreaRepository(databaseConnection)
	tokenRepository := repository.NewTokenRepository(databaseConnection)
	areaResultRepository := repository.NewAreaResultRepository(databaseConnection)

	// Services
	githubService := service.NewGithubService(
		githubRepository,
		serviceRepository,
		areaRepository,
		tokenRepository,
	)
	gmailService := service.NewGmailService(
		gmailRepository,
		serviceRepository,
		areaRepository,
		tokenRepository,
	)
	spotifyService := service.NewSpotifyService(
		spotifyRepository,
		serviceRepository,
		areaRepository,
		tokenRepository,
	)
	dropboxService := service.NewDropboxService(
		dropboxRepository,
		serviceRepository,
		areaRepository,
		tokenRepository,
	)
	timerService := service.NewTimerService(timerRepository, serviceRepository)
	openweathermapService := service.NewOpenweathermapService(
		openweathermapRepository,
		serviceRepository,
	)
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository, jwtService)
	serviceService := service.NewServiceService(
		serviceRepository,
		timerService,
		spotifyService,
		gmailService,
		githubService,
		dropboxService,
		openweathermapService,
	)
	actionService := service.NewActionService(actionRepository, serviceService)
	reactionService := service.NewReactionService(reactionRepository, serviceService)
	areaResultService := service.NewAreaResultService(areaResultRepository)
	areaService := service.NewAreaService(
		areaRepository,
		serviceService,
		actionService,
		reactionService,
		userService,
		areaResultService,
	)
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
	dropboxController := controller.NewDropboxController(
		dropboxService,
		userService,
		tokenService,
		serviceService,
	)
	userController := controller.NewUserController(userService, jwtService, tokenService)
	serviceController := controller.NewServiceController(
		serviceService,
		actionService,
		reactionService,
	)
	actionController := controller.NewActionController(actionService)
	reactionController := controller.NewReactionController(reactionService)
	areaController := controller.NewAreaController(areaService)
	tokenController := controller.NewTokenController(tokenService)
	areaResultController := controller.NewAreaResultController(areaResultService)

	// API routes
	api.NewActionApi(actionController, apiRoutes, userService)
	api.NewReactionApi(reactionController, apiRoutes, userService)
	api.NewTokenApi(tokenController)

	ping(apiRoutes)
	serviceAPI := api.NewServiceApi(serviceController, apiRoutes)
	api.NewUserApi(userController, apiRoutes, userService)
	api.NewSpotifyAPI(spotifyController, apiRoutes, userService)
	api.NewGmailAPI(gmailController, apiRoutes, userService)
	api.NewGithubAPI(githubController, apiRoutes, userService)
	api.NewDropboxAPI(dropboxController, apiRoutes, userService)
	api.NewAreaAPI(areaController, apiRoutes, userService)
	api.NewAreaResultAPI(areaResultController, apiRoutes)

	// basic about.json route
	router.GET("/about.json", serviceAPI.AboutJSON)

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

// func init() {
// err := .Load()
// if err != nil {
// 	panic("Error loading .env file")
// }
// }

// @securityDefinitions.apiKey	bearerAuth
// @in							header
// @name						Authorization
// @description				Use "Bearer <token>" as the format for the Authorization header.
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
