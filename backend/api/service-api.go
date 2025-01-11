package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/schemas"
)

// ServiceApi represents the API layer for the service, handling HTTP requests
// and delegating them to the appropriate service controller.
type ServiceApi struct {
	controller controller.ServiceController
}

// NewServiceApi initializes a new instance of ServiceApi, sets up the necessary
// routes for the service API, and returns a pointer to the created ServiceApi instance.
//
// Parameters:
//   - controller: an instance of ServiceController that handles the business logic for the service API.
//   - apiRoutes: a pointer to a gin.RouterGroup where the service API routes will be registered.
//
// Returns:
//   - A pointer to the initialized ServiceApi instance.
func NewServiceApi(
	controller controller.ServiceController,
	apiRoutes *gin.RouterGroup,
) *ServiceApi {
	apiRoutes = apiRoutes.Group("/service")
	api := ServiceApi{
		controller: controller,
	}
	apiRoutes = apiRoutes.Group("/info")
	api.GetServicesInfo(apiRoutes)
	api.GetServiceInfoById(apiRoutes)
	return &api
}

func (api *ServiceApi) AboutJSON(ctx *gin.Context) {
	aboutJSON, err := api.controller.AboutJSON(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
			Error: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, aboutJSON)
	}
}

// GetServicesInfo godoc
//
//	@Summary		get service info
//	@Description	get service info of service id
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Success		200	{object}	[]schemas.Service
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/service/info [get]
func (api *ServiceApi) GetServicesInfo(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		response, err := api.controller.GetServicesInfo()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}

// GetServiceInfoById godoc
//
//	@Summary		get service info
//	@Description	get service info of service id
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Param			id	path		int	true	"Service ID"
//	@Success		200	{object}	schemas.Service
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/service/info/:id [get]
func (api *ServiceApi) GetServiceInfoById(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &schemas.ErrorResponse{
				Error: err.Error(),
			})

			return
		}

		response, err := api.controller.GetServiceInfoById(idInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
