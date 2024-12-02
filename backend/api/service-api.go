package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"area/controller"
)

type ServiceApi struct {
	controller controller.ServiceController
}

func NewServiceApi(controller controller.ServiceController) *ServiceApi {
	return &ServiceApi{
		controller: controller,
	}
}

func (api *ServiceApi) AboutJson(ctx *gin.Context) {
	allServices, err := api.controller.AboutJson(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
