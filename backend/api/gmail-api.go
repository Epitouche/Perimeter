package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type GmailAPI struct {
	controller controller.GmailController
}

func NewGmailAPI(controller controller.GmailController, apiRoutes *gin.RouterGroup) *GmailAPI {
	apiRoutes = apiRoutes.Group("/gmail")
	api := GmailAPI{
		controller: controller,
	}
	api.RedirectToService(apiRoutes)
	api.HandleServiceCallback(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT())
	api.GetUserInfo(apiRoutesInfo)
	return &api
}

// HandleServiceCallback godoc
// @Summary give url to authenticate with gmail
// @Description give url to authenticate with gmail
// @Tags Gmail
// @Accept json
// @Produce json
// @Success 200 {object} schemas.AuthenticationUrl
// @Failure 500 {object} schemas.ErrorRespose
// @Router /gmail/auth [get]
func (api *GmailAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx, apiRoutes.BasePath()+"/auth/callback")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, schemas.ErrorRespose{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, schemas.AuthenticationUrl{Url: authURL})
		}
	})
}

// HandleServiceCallback godoc
// @Summary give url to authenticate with gmail
// @Description give url to authenticate with gmail
// @Tags Gmail
// @Accept json
// @Produce json
// @Success 200 {object} schemas.JWT
// @Failure 500 {object} schemas.ErrorRespose
// @Router /gmail/auth/callback [post]
func (api *GmailAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		gmail_token, err := api.controller.HandleServiceCallback(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: gmail_token})
		}
	})
}

// GetUserInfo godoc
// @Summary give user info of gmail
// @Description give user info of gmail
// @Tags Gmail
// @Accept json
// @Produce json
// @Success 200 {object} schemas.UserCredentials
// @Failure 500 {object} schemas.ErrorRespose
// @Router /gmail/info/user [get]
func (api *GmailAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/user", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
