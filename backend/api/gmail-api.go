package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Epitouche/Perimeter/controller"
	"github.com/Epitouche/Perimeter/middlewares"
	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/service"
)

type GmailAPI struct {
	controller controller.GmailController
}

func NewGmailAPI(
	controller controller.GmailController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *GmailAPI {
	apiRoutes = apiRoutes.Group("/gmail")
	api := GmailAPI{
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
//	@Summary		give url to authenticate with gmail
//	@Description	give url to authenticate with gmail
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationURL
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/gmail/auth [get]
func (api *GmailAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
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
//	@Summary		give url to authenticate with gmail
//	@Description	give url to authenticate with gmail
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/gmail/auth/callback [post]
func (api *GmailAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
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
//	@Summary		give authentication token to mobile
//	@Description	give authentication token to mobile
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Param			payload			body		schemas.CodeCredentials	true	"Callback Payload"
//	@Param			Authorization	header		string					false	"Bearer token"
//	@Success		200				{object}	schemas.JWT
//	@Failure		500				{object}	schemas.ErrorResponse
//	@Router			/gmail/auth/callback/mobile [post]
func (api *GmailAPI) HandleServiceCallbackMobile(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback/mobile", func(ctx *gin.Context) {
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Printf("Failed to read body: %v\n", err)
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
				Error: "Invalid request body",
			})
			return
		}

		// Print the body
		fmt.Printf("body: %s\n", string(bodyBytes))
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
//	@Summary		give user info of gmail
//	@Description	give user info of gmail
//	@Tags			Gmail
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.UserCredentials
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/gmail/info [get]
func (api *GmailAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
