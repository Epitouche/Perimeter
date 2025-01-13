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

// ActionApi represents the API layer for handling actions.
// It contains a reference to the ActionController which manages the business logic.
type ActionApi struct {
	controller controller.ActionController
}

// NewActionApi initializes a new instance of ActionApi with the provided controller,
// sets up the API routes with the necessary middleware, and registers the action info routes.
//
// Parameters:
//   - controller: An instance of ActionController to handle action-related operations.
//   - apiRoutes: A pointer to a gin.RouterGroup where the action routes will be registered.
//   - serviceUser: An instance of UserService used for JWT authorization middleware.
//
// Returns:
//   - A pointer to the initialized ActionApi instance.
func NewActionApi(
	controller controller.ActionController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *ActionApi {
	apiRoutes = apiRoutes.Group("/action", middlewares.AuthorizeJWT(serviceUser))
	api := ActionApi{
		controller: controller,
	}
	apiRoutes = apiRoutes.Group("/info")
	api.GetActionsInfo(apiRoutes)

	return &api
}

// GetActionsInfo godoc
//
//	@Summary		get action info
//	@Description	get action info of service id
//	@Tags			Action
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	[]schemas.Action
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/action/info/:id [get]
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
