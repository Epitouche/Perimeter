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
//
//	@Summary		give url to authenticate with github
//	@Description	give url to authenticate with github
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.AuthenticationUrl
//	@Failure		500	{object}	schemas.ErrorRespose
//	@Router			/github/auth [get]
func (api *GithubAPI) RedirectToService(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/auth", func(ctx *gin.Context) {
		authURL, err := api.controller.RedirectToService(ctx, apiRoutes.BasePath()+"/auth/callback")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, &schemas.AuthenticationUrl{Url: authURL})
		}
	})
}

// HandleServiceCallback godoc
//
//	@Summary		give url to authenticate with github
//	@Description	give url to authenticate with github
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		schemas.CodeCredentials	true	"Callback Payload"
//	@Success		200		{object}	schemas.JWT
//	@Failure		500		{object}	schemas.ErrorRespose
//	@Router			/github/auth/callback [post]
func (api *GithubAPI) HandleServiceCallback(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/auth/callback", func(ctx *gin.Context) {
		github_token, err := api.controller.HandleServiceCallback(
			ctx,
			apiRoutes.BasePath()+"/auth/callback",
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, &schemas.JWT{Token: github_token})
		}
	})
}

// GetUserInfo godoc
//
//	@Summary		give user info of github
//	@Description	give user info of github
//	@Tags			Github
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Success		200				{object}	schemas.UserCredentials
//	@Failure		401				{object}	schemas.ErrorRespose
//	@Failure		500				{object}	schemas.ErrorRespose
//	@Router			/github/info/user [get]
func (api *GithubAPI) GetUserInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/user", func(ctx *gin.Context) {
		userInfo, err := api.controller.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, userInfo)
		}
	})
}
