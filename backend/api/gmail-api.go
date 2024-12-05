package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
)

type GmailAPI struct {
	controller controller.GmailController
}

func NewGmailAPI(controller controller.GmailController) *GmailAPI {
	return &GmailAPI{
		controller: controller,
	}
}

func (api *GmailAPI) RedirectToService(ctx *gin.Context, path string) {
	authURL, err := api.controller.RedirectToService(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"authentication_url": authURL})
	}
}

func (api *GmailAPI) HandleServiceCallback(ctx *gin.Context, path string) {
	github_token, err := api.controller.HandleServiceCallback(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": github_token})
	}
}

func (api *GmailAPI) GetUserInfo(ctx *gin.Context) {
	userInfo, err := api.controller.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"login": userInfo.Login, "email": userInfo.Email}})
	}
}
