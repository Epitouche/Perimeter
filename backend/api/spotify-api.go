package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
)

type SpotifyAPI struct {
	controller controller.SpotifyController
}

func NewSpotifyAPI(controller controller.SpotifyController) *SpotifyAPI {
	return &SpotifyAPI{
		controller: controller,
	}
}

func (api *SpotifyAPI) RedirectToService(ctx *gin.Context, path string) {
	authURL, err := api.controller.RedirectToService(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"authentication_url": authURL})
	}
}

func (api *SpotifyAPI) HandleServiceCallback(ctx *gin.Context, path string) {
	spotify_token, err := api.controller.HandleServiceCallback(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": spotify_token})
	}
}

func (api *SpotifyAPI) GetUserInfo(ctx *gin.Context) {
	usetInfo, err := api.controller.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"name": usetInfo.Login, "email": usetInfo.Email}})
	}
}
