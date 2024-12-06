package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type ReactionApi struct {
	controller controller.ReactionController
}

func NewReactionApi(
	controller controller.ReactionController,
	apiRoutes *gin.RouterGroup,
) *ReactionApi {
	apiRoutes = apiRoutes.Group("/reaction", middlewares.AuthorizeJWT())
	api := ReactionApi{
		controller: controller,
	}
	apiRoutes = apiRoutes.Group("/info")
	api.GetReactionsInfo(apiRoutes)
	return &api
}

func (api *ReactionApi) GetReactionsInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		response, err := api.controller.GetReactionsInfo(idInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
}
