package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type GithubAPI struct {
	controller controller.GithubController
}

func NewGithubAPI(controller controller.GithubController, apiRoutes *gin.RouterGroup) *GithubAPI {
	apiRoutes = apiRoutes.Group("/github")
	api := GithubAPI{
		controller: controller,
	}
	api.RedirectToService(apiRoutes)
	api.HandleServiceCallback(apiRoutes)
	apiRoutesInfo := apiRoutes.Group("/info", middlewares.AuthorizeJWT())
	api.GetUserInfo(apiRoutesInfo)
	return &api
}

// RedirectToService godoc
// @Summary give url to authenticate with github
// @Description give url to authenticate with github
// @Tags github route
// @Accept json
// @Produce json
// @Success 200 {string} Bearer token
// @Failure 500 {object} schemas.Response
// @Router /github/auth [get]
func (api *GithubAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
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
// @Summary give url to authenticate with github
// @Description give url to authenticate with github
// @Tags github route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /github/auth/callback [get]
func (api *GithubAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth/callback", func(ctx *gin.Context) {
		github_token, err := api.controller.HandleServiceCallback(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"access_token": github_token})
		}
	})
}

// GetUserInfo godoc
// @Summary give user info of github
// @Description give user info of github
// @Tags github route
// @Accept json
// @Produce json
// @Success 200 {object} schemas.Response
// @Failure 500 {object} schemas.ErrorRespose
// @Router /github/info/user [get]
func (api *GithubAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/user", func(ctx *gin.Context) {
		usetInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"id": usetInfo.Id, "name": usetInfo.Name, "login": usetInfo.Login, "email": usetInfo.Email, "avatar_url": usetInfo.AvatarUrl, "html_url": usetInfo.HtmlUrl, "type": usetInfo.Type}})
		}
	})
}
