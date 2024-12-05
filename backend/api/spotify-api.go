package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type SpotifyAPI struct {
	controller controller.SpotifyController
}

func NewSpotifyAPI(
	controller controller.SpotifyController,
	apiRoutes *gin.RouterGroup,
) *SpotifyAPI {
	apiRoutes = apiRoutes.Group("/spotify")
	api := SpotifyAPI{
		controller: controller,
	}
	api.RedirectToService(apiRoutes)
	api.HandleServiceCallback(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT())
	api.GetUserInfo(apiRoutesInfo)
	return &api
}

// RedirectToService godoc
// @Summary give url to authenticate with spotify
// @Description give url to authenticate with spotify
// @Tags spotify route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /spotify/auth [get]
func (api *SpotifyAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx, apiRoutes.BasePath()+"/auth/callback")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"authentication_url": authURL})
		}
	})
}

// HandleServiceCallback godoc
// @Summary give url to authenticate with spotify
// @Description give url to authenticate with spotify
// @Tags spotify route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /spotify/auth/callback [get]
func (api *SpotifyAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth/callback", func(ctx *gin.Context) {
		spotify_token, err := api.controller.HandleServiceCallback(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"access_token": spotify_token})
		}
	})
}

// GetUserInfo godoc
// @Summary give user info of spotify
// @Description give user info of spotify
// @Tags spotify route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /spotify/info/user [get]
func (api *SpotifyAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/user", func(ctx *gin.Context) {
		usetInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"name": usetInfo.Login, "email": usetInfo.Email}})
		}
	})
}
