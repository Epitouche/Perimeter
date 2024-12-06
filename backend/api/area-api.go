package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type AreaApi struct {
	controller controller.AreaController
}

func NewAreAPI(controller controller.AreaController, apiRoutes *gin.RouterGroup) *AreaApi {
	apiRoutes = apiRoutes.Group("/area", middlewares.AuthorizeJWT())
	api := AreaApi{
		controller: controller,
	}
	api.CreateArea(apiRoutes)
	return &api
}

func (api *AreaApi) CreateArea(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/", func(ctx *gin.Context) {
		response, err := api.controller.CreateArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorRespose{
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
}
