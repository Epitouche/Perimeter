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

// ReactionApi is a struct that provides an API for handling reactions.
// It contains a controller of type ReactionController which manages the
// business logic related to reactions.
type ReactionApi struct {
	controller controller.ReactionController
}

// NewReactionApi initializes a new ReactionApi instance, sets up the necessary routes,
// and applies the JWT authorization middleware.
//
// Parameters:
//   - controller: An instance of ReactionController to handle reaction-related operations.
//   - apiRoutes: A pointer to the gin.RouterGroup where the reaction routes will be registered.
//   - serviceUser: An instance of UserService used for JWT authorization.
//
// Returns:
//   - A pointer to the initialized ReactionApi instance.
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

// GetReactionInfoByReactionID godoc
//
//	@Summary		get reaction info of reaction id
//	@Description	get reaction info of reaction id
//	@Tags			Reaction
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.Reaction
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/reaction/info/reaction/:idReaction [get]
func (api *ReactionApi) GetRereactionInfoByReactionID(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/reaction/:idReaction", func(ctx *gin.Context) {
		idReaction := ctx.Param("idReaction")

		idInt, err := strconv.ParseUint(idReaction, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
				Error: err.Error(),
			})

			return
		}

		response, err := api.controller.GetReactionByReactionID(idInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
