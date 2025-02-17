package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

// SpotifyAPI is a struct that provides an interface to interact with the Spotify API.
// It contains a controller of type SpotifyController which handles the business logic
// for communicating with Spotify's services.
type SpotifyAPI struct {
	controller controller.SpotifyController
}

// NewSpotifyAPI initializes a new SpotifyAPI instance, sets up the necessary routes,
// and returns a pointer to the SpotifyAPI instance.
//
// Parameters:
//   - controller: an instance of SpotifyController to handle Spotify-related operations.
//   - apiRoutes: a pointer to a gin.RouterGroup where the Spotify routes will be registered.
//   - serviceUser: an instance of UserService to handle user-related operations.
//
// Returns:
//   - A pointer to the initialized SpotifyAPI instance.
func NewSpotifyAPI(
	controller controller.SpotifyController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *SpotifyAPI {
	apiRoutes = apiRoutes.Group("/spotify")
	api := SpotifyAPI{
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
//	@Summary		Redirect To Service
//	@Description	give url to authenticate with spotify
//	@Tags			Spotify
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationURL
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/spotify/auth [get]
func (api *SpotifyAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, schemas.AuthenticationURL{URL: authURL})
		}
	})
}

// HandleServiceCallback godoc
//
//	@Summary		Handle Service Callback
//	@Description	give authentication token to web client
//	@Tags			Spotify
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/spotify/auth/callback [post]
func (api *SpotifyAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		spotify_token, err := api.controller.HandleServiceCallback(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: spotify_token})
		}
	})
}

// HandleServiceCallbackMobile godoc
//
//	@Summary		Handle Service Callback Mobile
//	@Description	give authentication token to mobile
//	@Tags			Spotify
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/spotify/auth/callback/mobile [post]
func (api *SpotifyAPI) HandleServiceCallbackMobile(apiRoutes *gin.RouterGroup) {
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
//	@Summary		Get User Info
//	@Description	give user info of spotify
//	@Tags			Spotify
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/spotify/info [get]
func (api *SpotifyAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		usetInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, usetInfo)
		}
	})
}
