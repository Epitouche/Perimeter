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

func (api *AreaApi) GetArea(ctx *gin.Context) {
	ctx.JSON(200, "status: success")
}
