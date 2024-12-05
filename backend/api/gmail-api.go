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

func (api *GmailAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth/callback", func(ctx *gin.Context) {
		github_token, err := api.controller.HandleServiceCallback(ctx, apiRoutes.BasePath()+"/auth/callback")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"access_token": github_token})
		}
	})
}

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
