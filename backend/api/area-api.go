package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Epitouche/Perimeter/controller"
	"github.com/Epitouche/Perimeter/middlewares"
	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/service"
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
	api.UpdateUserArea(apiRoutes)
	api.DeleteUserArea(apiRoutes)
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
			fmt.Printf("Error: %v\n", err.Error())
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

// UpdateUserArea godoc
//
//	@Summary		update user area
//	@Description	update user area list
//	@Tags			Area
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.Area
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/area [put]
func (api *AreaApi) UpdateUserArea(apiRoutes *gin.RouterGroup) {
	apiRoutes.PUT("/", func(ctx *gin.Context) {
		response, err := api.controller.UpdateUserArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}

// DeleteUserArea godoc
//
//	@Summary		delete user area
//	@Description	delete user area list
//	@Tags			Area
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Security		bearerAuth
//	@Success		200	{object}	schemas.Area
//	@Failure		401	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/area [delete]
func (api *AreaApi) DeleteUserArea(apiRoutes *gin.RouterGroup) {
	apiRoutes.DELETE("/", func(ctx *gin.Context) {
		response, err := api.controller.DeleteUserArea(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
