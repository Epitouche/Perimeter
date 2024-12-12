package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
)

type ServiceApi struct {
	controller controller.ServiceController
}

func NewServiceApi(
	controller controller.ServiceController,
	apiRoutes *gin.RouterGroup,
) *ServiceApi {
	apiRoutes = apiRoutes.Group("/service")
	api := ServiceApi{
		controller: controller,
	}
	apiRoutes = apiRoutes.Group("/info", middlewares.AuthorizeJWT())
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
//	@Security		Bearer
//	@Security		bearerAuth
//	@Success		200	{object}	[]schemas.Service
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/service/info/ [get]
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
