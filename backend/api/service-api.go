package api

import (
	"fmt"
	"net/http"
	"time"

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
	return &api
}

func (api *ServiceApi) AboutJson(ctx *gin.Context) {
	allServices, err := api.controller.AboutJson(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.ErrorResponse{
			Error: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"client": map[string]string{
				"host": ctx.ClientIP(),
			},
			"server": map[string]any{
				"current_time": fmt.Sprintf("%d", time.Now().Unix()),
				"services":     allServices,
			},
		})
	}
}

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
