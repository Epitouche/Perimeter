package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

type TokenApi struct {
	controller controller.TokenController
}

func NewTokenApi(controller controller.TokenController, apiRoutes *gin.RouterGroup, serviceUser service.UserService) *TokenApi {
	apiRoutes = apiRoutes.Group("/token", middlewares.AuthorizeJWT(serviceUser))
	api := TokenApi{
		controller: controller,
	}
	api.DeleteUserToken(apiRoutes)
	return &api
}

// DeleteUserToken godoc
//
//	@Summary		delete user token
//	@Description	delete user token list
//	@Tags			Token
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.Token
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/token [delete]
func (api *TokenApi) DeleteUserToken(apiRoutes *gin.RouterGroup) {
	apiRoutes.DELETE("/", func(ctx *gin.Context) {
		response, err := api.controller.DeleteUserToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
