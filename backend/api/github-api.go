package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/schemas"
)

type GithubAPI struct {
	controller controller.GithubController
}

func NewGithubAPI(controller controller.GithubController) *GithubAPI {
	return &GithubAPI{
		controller: controller,
	}
}

// @Summary give url to authenticate with github
// @Description give url to authenticate with github
// @Tags github route
// @Accept json
// @Produce json
// @Success 200 {string} Bearer token
// @Error 500 {object} schemas.Response
// @Router /github/auth [get]
func (api *GithubAPI) RedirectToService(ctx *gin.Context, path string) {
	authURL, err := api.controller.RedirectToService(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
			Error: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"authentication_url": authURL})
	}
}

func (api *GithubAPI) HandleServiceCallback(ctx *gin.Context, path string) {
	github_token, err := api.controller.HandleServiceCallback(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
			Error: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": github_token})
	}
}

func (api *GithubAPI) GetUserInfo(ctx *gin.Context) {
	usetInfo, err := api.controller.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
			Error: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user_info": gin.H{"id": usetInfo.Id, "name": usetInfo.Name, "login": usetInfo.Login, "email": usetInfo.Email, "avatar_url": usetInfo.AvatarUrl, "html_url": usetInfo.HtmlUrl, "type": usetInfo.Type}})
	}
}
