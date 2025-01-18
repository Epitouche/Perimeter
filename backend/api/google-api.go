package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

// GoogleAPI is a struct that provides an interface to interact with the Gmail service.
// It contains a controller of type GoogleController which handles the core logic
// for managing Gmail-related operations.
type GoogleAPI struct {
	controller controller.GoogleController
}

// NewGoogleAPI initializes a new GoogleAPI instance, sets up the necessary routes, and returns the instance.
// It configures the following routes:
// - /gmail: Base route for Gmail API.
// - /gmail/info: Route for getting user information, protected by JWT authorization middleware.
//
// Parameters:
// - controller: An instance of GoogleController to handle Gmail-related operations.
// - apiRoutes: A gin.RouterGroup to define the API routes.
// - serviceUser: An instance of UserService to handle user-related operations.
//
// Returns:
// - A pointer to the initialized GoogleAPI instance.
func NewGoogleAPI(
	controller controller.GoogleController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *GoogleAPI {
	apiRoutes = apiRoutes.Group("/google")
	api := GoogleAPI{
		controller: controller,
	}
	api.RedirectToService(apiRoutes)
	api.HandleServiceCallback(apiRoutes)
	api.HandleServiceCallbackMobile(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT(serviceUser))
	api.GetUserInfo(apiRoutesInfo)
	return &api
}

// HandleServiceCallback godoc
//
//	@Summary		Redirect To Service
//	@Description	give url to authenticate with gmail
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationURL
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/gmail/auth [get]
func (api *GoogleAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, schemas.AuthenticationURL{URL: authURL})
		}
	})
}

// HandleServiceCallback godoc
//
//	@Summary		Handle Service Callback
//	@Description	give url to authenticate with gmail
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/gmail/auth/callback [post]
func (api *GoogleAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		gmail_token, err := api.controller.HandleServiceCallback(
			ctx,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: gmail_token})
		}
	})
}

// HandleServiceCallbackMobile godoc
//
//	@Summary		Handle Service Callback Mobile
//	@Description	give authentication token to mobile
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/gmail/auth/callback/mobile [post]
func (api *GoogleAPI) HandleServiceCallbackMobile(apiRoutes *gin.RouterGroup) {
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
//	@Description	give user info of gmail
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/gmail/info [get]
func (api *GoogleAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
