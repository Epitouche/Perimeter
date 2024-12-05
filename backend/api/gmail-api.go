package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
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
// @Tags gmail route
// @Accept json
// @Produce json
// @Success 200 {string} Bearer token
// @Failure 500 {object} schemas.Response
// @Router /gmail/auth [get]
func (api *GmailAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx, apiRoutes.BasePath()+"/auth/callback")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"authentication_url": authURL})
		}
	})
}

// HandleServiceCallback godoc
// @Summary give url to authenticate with gmail
// @Description give url to authenticate with gmail
// @Tags gmail route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /gmail/auth/callback [get]
func (api *GmailAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth/callback", func(ctx *gin.Context) {
		gmail_token, err := api.controller.HandleServiceCallback(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"access_token": gmail_token})
		}
	})
}

// GetUserInfo godoc
// @Summary give user info of gmail
// @Description give user info of gmail
// @Tags gmail route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /gmail/info/user [get]
func (api *GmailAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/user", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"login": userInfo.Login, "email": userInfo.Email}})
		}
	})
}
