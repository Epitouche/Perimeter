package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type ActionApi struct {
	controller controller.ActionController
}

func NewActionApi(controller controller.ActionController, apiRoutes *gin.RouterGroup) *ActionApi {
	apiRoutes = apiRoutes.Group("/action", middlewares.AuthorizeJWT())
	api := ActionApi{
		controller: controller,
	}
	apiRoutes = apiRoutes.Group("/info")
	api.GetActionsInfo(apiRoutes)
	return &api
}

func (api *ActionApi) GetActionsInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		response, err := api.controller.GetActionsInfo(idInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
}
