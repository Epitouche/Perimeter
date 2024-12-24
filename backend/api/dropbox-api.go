package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type DropboxAPI struct {
	controller controller.DropboxController
}

func NewDropboxAPI(
	controller controller.DropboxController,
	apiRoutes *gin.RouterGroup,
) *DropboxAPI {
	apiRoutes = apiRoutes.Group("/dropbox")
	api := DropboxAPI{
		controller: controller,
	}
	api.RedirectToService(apiRoutes)
	api.HandleServiceCallback(apiRoutes)
	api.HandleServiceCallbackMobile(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT())
	api.GetUserInfo(apiRoutesInfo)
	return &api
}

// HandleServiceCallback godoc
//
//	@Summary		give url to authenticate with dropbox
//	@Description	give url to authenticate with dropbox
//	@Tags			Dropbox
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationURL
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/dropbox/auth [get]
func (api *DropboxAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
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
//	@Summary		give url to authenticate with dropbox
//	@Description	give url to authenticate with dropbox
//	@Tags			Dropbox
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/dropbox/auth/callback [post]
func (api *DropboxAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		dropbox_token, err := api.controller.HandleServiceCallback(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: dropbox_token})
		}
	})
}

// HandleServiceCallbackMobile godoc
//
//	@Summary		give url to authenticate with dropbox
//	@Description	give url to authenticate with dropbox
//	@Tags			Dropbox
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Param			payload	body		schemas.CodeCredentials	true	"Callback Payload"
//	@Success		200		{object}	schemas.JWT
//	@Failure		500		{object}	schemas.ErrorResponse
//	@Router			/dropbox/auth/callback/mobile [post]
func (api *DropboxAPI) HandleServiceCallbackMobile(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback/mobile", func(ctx *gin.Context) {
		token, err := api.controller.HandleServiceCallbackMobile(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: token})
		}
	})
}

// GetUserInfo godoc
//
//	@Summary		give user info of dropbox
//	@Description	give user info of dropbox
//	@Tags			Dropbox
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/dropbox/info [get]
func (api *DropboxAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
