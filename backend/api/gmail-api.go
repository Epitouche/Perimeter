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
		ctx.JSON(http.StatusOK, gin.H{"github_authentication_url": authURL})
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
	usetInfo, err := api.controller.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"id": usetInfo.Id, "name": usetInfo.Name, "login": usetInfo.Login, "email": usetInfo.Email, "avatar_url": usetInfo.AvatarUrl, "html_url": usetInfo.HtmlUrl, "type": usetInfo.Type}})
	}
}
