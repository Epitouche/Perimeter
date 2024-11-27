package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

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

	var (
		// Database connection
		databaseConnection *gorm.DB = database.Connection()

		// Repositories
		linkRepository        repository.LinkRepository        = repository.NewLinkRepository(databaseConnection)
		githubTokenRepository repository.GithubTokenRepository = repository.NewGithubTokenRepository(databaseConnection)
		userRepository        repository.UserRepository        = repository.NewUserRepository(databaseConnection)
		serviceRepository     repository.ServiceRepository     = repository.NewServiceRepository(databaseConnection)

		// Services
		linkService        service.LinkService        = service.NewLinkService(linkRepository)
		githubTokenService service.GithubTokenService = service.NewGithubTokenService(githubTokenRepository)
		jwtService         service.JWTService         = service.NewJWTService()
		userService        service.UserService        = service.NewUserService(userRepository, jwtService)
		serviceService     service.ServiceService     = service.NewServiceService(serviceRepository)

		// Controllers
		linkController        controller.LinkController        = controller.NewLinkController(linkService)
		githubTokenController controller.GithubTokenController = controller.NewGithubTokenController(githubTokenService, userService)
		userController        controller.UserController        = controller.NewUserController(userService, jwtService)
		serviceController     controller.ServiceController     = controller.NewServiceController(serviceService)
	)

	linkApi := api.NewLinkApi(linkController)

	userApi := api.NewUserApi(userController)

	githubApi := api.NewGithubApi(githubTokenController)

	api.NewServiceApi(serviceController)

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

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	router := setupRouter()

	// basic about.json route
	router.GET("/about.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"client": map[string]string{
				"host": "localhost",
				"port": "3000",
			},
			"server": map[string]string{
				"current_time": "2021-09-01T00:00:00Z",
				"services":     "area",
			},
		})
	})

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
