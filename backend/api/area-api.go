package api

import (
	"github.com/gin-gonic/gin"

	"area/controller"
)

type AreaApi struct {
	controller controller.AreaController
}

func NewAreAPI(controller controller.AreaController) *AreaApi {
	return &AreaApi{
		controller: controller,
	}
}

func (api *AreaApi) CreateArea(ctx *gin.Context) {
	response, err := api.controller.CreateArea(ctx)
	if err != nil {
		ctx.JSON(500, "status: error")
		return
	}
	ctx.JSON(200, response)
}
