package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area/controller"
	"area/middlewares"
	"area/schemas"
	"area/service"
)

type AreaApi struct {
	controller controller.AreaController
}

func NewAreaAPI(
	controller controller.AreaController,
	apiRoutes *gin.RouterGroup,
	serviceUser service.UserService,
) *AreaApi {
	apiRoutes = apiRoutes.Group("/area", middlewares.AuthorizeJWT(serviceUser))
	api := AreaApi{
		controller: controller,
	}
	api.CreateArea(apiRoutes)
	api.GetUserAreas(apiRoutes)
	return &api
}

// CreateArea godoc
//
//	@Summary		create area
//	@Description	create area
//	@Tags			Area
//	@Accept			json
//	@Produce		json
//	@Security		bearerAuth
//	@Param			payload	body		schemas.AreaMessage	true	"Area Payload"
//	@Success		200		{object}	schemas.Response
//	@Failure		500		{object}	schemas.ErrorResponse
//	@Router			/area [post]
func (api *AreaApi) CreateArea(apiRoutes *gin.RouterGroup) {
	apiRoutes.POST("/", func(ctx *gin.Context) {
		response, err := api.controller.CreateArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: response,
		})
	})
}

// GetUserAreas godoc
//
//	@Summary		get user areas
//	@Description	get user areas list
//	@Tags			Area
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Security		bearerAuth
//	@Success		200	{object}	[]schemas.Area
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/area [get]
func (api *AreaApi) GetUserAreas(apiRoutes *gin.RouterGroup) {
	apiRoutes.GET("/", func(ctx *gin.Context) {
		response, err := api.controller.GetUserAreas(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}

// GetUserAreas godoc
//
//	@Summary		update user area
//	@Description	update user area list
//	@Tags			Area
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Security		bearerAuth
//	@Success		200	{object}	[]schemas.Area
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/area [put]
func (api *AreaApi) UpdateUserArea(apiRoutes *gin.RouterGroup) {
	apiRoutes.PUT("/", func(ctx *gin.Context) {
		response, err := api.controller.GetUserAreas(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
