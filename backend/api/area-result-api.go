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

// AreaResultApi is a struct that provides an API for handling area result operations.
// It contains a controller of type AreaResultController which manages the business logic
// for area result-related actions.
type AreaResultApi struct {
	controller controller.AreaResultController
}

// NewAreaResultAPI creates a new instance of AreaResultApi with the provided controller and API routes.
// It initializes the AreaResultApi struct with the given AreaResultController.
//
// Parameters:
//   - controller: An instance of AreaResultController that handles the business logic for area results.
//   - apiRoutes: A pointer to a gin.RouterGroup that defines the API routes for the area results.
//
// Returns:
//   - A pointer to an initialized AreaResultApi struct.
func NewAreaResultAPI(
	controller controller.AreaResultController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *AreaResultApi {
	apiRoutes = apiRoutes.Group("/area-result", middlewares.AuthorizeJWT(serviceUser))
	api := AreaResultApi{
		controller: controller,
	}
	api.GetUserAreaResultsByAreaID(apiRoutes)
	return &api
}

// GetUserAreaResultsByAreaID godoc
//
//	@Summary		Get User Area Results By Area ID
//	@Description	get user areas results list by area id
//	@Tags			AreaResults
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	[]schemas.AreaResult
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/area-result/:id [get]
func (api *AreaResultApi) GetUserAreaResultsByAreaID(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
				Error: err.Error(),
			})

			return
		}

		response, err := api.controller.GetUserAreaResultsByAreaID(ctx, idInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
