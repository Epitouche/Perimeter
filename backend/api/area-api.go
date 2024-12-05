package api

import (
	"github.com/gin-gonic/gin"

	"area/controller"
)

type AreaApi struct {
	controller controller.AreaController
}

func NewAreAPI(controller controller.AreaController, apiRoutes *gin.RouterGroup) *AreaApi {
	apiRoutes = apiRoutes.Group("/area")
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
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, response)
	})
}
