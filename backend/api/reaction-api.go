package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

type ReactionApi struct {
	controller controller.ReactionController
}

func NewReactionApi(
	controller controller.ReactionController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *ReactionApi {
	apiRoutes = apiRoutes.Group("/reaction", middlewares.AuthorizeJWT(serviceUser))
	api := ReactionApi{
		controller: controller,
	}
	apiRoutes = apiRoutes.Group("/info")
	api.GetReactionsInfo(apiRoutes)
	return &api
}

// GetReactionsInfo godoc
//
//	@Summary		get reaction info
//	@Description	get reaction info of service id
//	@Tags			Reaction
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Security		bearerAuth
//	@Param			id	path		int	true	"Service ID"
//	@Success		200	{object}	[]schemas.Reaction
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/reaction/info/:id [get]
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
