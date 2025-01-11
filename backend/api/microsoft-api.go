package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

// MicrosoftAPI is a struct that provides an interface to interact with the MicrosoftController.
// It serves as a bridge between the API layer and the controller layer, facilitating the handling
// of requests and responses related to Microsoft services.
type MicrosoftAPI struct {
	controller controller.MicrosoftController
}

// NewMicrosoftAPI initializes a new MicrosoftAPI instance, sets up the necessary routes,
// and returns a pointer to the MicrosoftAPI instance.
//
// Parameters:
//   - controller: An instance of MicrosoftController to handle the API logic.
//   - apiRoutes: A pointer to a gin.RouterGroup where the Microsoft API routes will be registered.
//   - serviceUser: An instance of UserService to handle user-related operations.
//
// Returns:
//   - A pointer to the initialized MicrosoftAPI instance.
func NewMicrosoftAPI(
	controller controller.MicrosoftController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *MicrosoftAPI {
	apiRoutes = apiRoutes.Group("/microsoft")
	api := MicrosoftAPI{
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
//	@Summary		give url to authenticate with microsoft
//	@Description	give url to authenticate with microsoft
//	@Tags			Microsoft
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationURL
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/microsoft/auth [get]
func (api *MicrosoftAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
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
//	@Summary		give url to authenticate with microsoft
//	@Description	give url to authenticate with microsoft
//	@Tags			Microsoft
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/microsoft/auth/callback [post]
func (api *MicrosoftAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		microsoft_token, err := api.controller.HandleServiceCallback(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: microsoft_token})
		}
	})
}

// HandleServiceCallbackMobile godoc
//
//	@Summary		give url to authenticate with microsoft
//	@Description	give url to authenticate with microsoft
//	@Tags			Microsoft
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Param			payload	body		schemas.CodeCredentials	true	"Callback Payload"
//	@Success		200		{object}	schemas.JWT
//	@Failure		500		{object}	schemas.ErrorResponse
//	@Router			/microsoft/auth/callback/mobile [post]
func (api *MicrosoftAPI) HandleServiceCallbackMobile(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback/mobile", func(ctx *gin.Context) {
		spotify_token, err := api.controller.HandleServiceCallbackMobile(ctx)
		if err != nil {
			println("--------------------")
			println(err.Error())
			println("--------------------")
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
//	@Summary		give user info of microsoft
//	@Description	give user info of microsoft
//	@Tags			Microsoft
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/microsoft/info [get]
func (api *MicrosoftAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
