package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

// GithubAPI is a struct that provides an interface to interact with the GithubController.
// It contains a single field, controller, which is an instance of GithubController.
type GithubAPI struct {
	controller controller.GithubController
}

// NewGithubAPI initializes a new GithubAPI instance, sets up the necessary routes,
// and returns a pointer to the created GithubAPI instance.
//
// Parameters:
//   - controller: An instance of GithubController to handle GitHub-related operations.
//   - apiRoutes: A pointer to a gin.RouterGroup where the GitHub API routes will be registered.
//   - serviceUser: An instance of UserService to handle user-related operations.
//
// Returns:
//   - A pointer to the initialized GithubAPI instance.
func NewGithubAPI(
	controller controller.GithubController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *GithubAPI {
	apiRoutes = apiRoutes.Group("/github")
	api := GithubAPI{
		controller: controller,
	}
	api.RedirectToService(apiRoutes)
	api.HandleServiceCallback(apiRoutes)
	api.HandleServiceCallbackMobile(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT(serviceUser))
	api.GetUserInfo(apiRoutesInfo)
	return &api
}

// RedirectToService godoc
//
//	@Summary		give url to authenticate with github
//	@Description	give url to authenticate with github
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationURL
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/github/auth [get]
func (api *GithubAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, &schemas.AuthenticationURL{URL: authURL})
		}
	})
}

// HandleServiceCallback godoc
//
//	@Summary		give url to authenticate with github
//	@Description	give url to authenticate with github
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		schemas.CodeCredentials	true	"Callback Payload"
//	@Success		200		{object}	schemas.JWT
//	@Failure		500		{object}	schemas.ErrorResponse
//	@Router			/github/auth/callback [post]
func (api *GithubAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		github_token, err := api.controller.HandleServiceCallback(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: github_token})
		}
	})
}

// HandleServiceCallbackMobile godoc
//
//	@Summary		give authentication token to mobile
//	@Description	give authentication token to mobile
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/github/auth/callback/mobile [post]
func (api *GithubAPI) HandleServiceCallbackMobile(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback/mobile", func(ctx *gin.Context) {
		spotify_token, err := api.controller.HandleServiceCallbackMobile(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: spotify_token})
		}
	})
}

// GetUserInfo godoc
//
//	@Summary		give user info of github
//	@Description	give user info of github
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/github/info [get]
func (api *GithubAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
