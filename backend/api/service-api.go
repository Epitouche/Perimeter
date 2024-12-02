package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"area/controller"
)

type ServiceApi struct {
	serviceController controller.ServiceController
}

func NewServiceApi(serviceController controller.ServiceController) *ServiceApi {
	return &ServiceApi{
		serviceController: serviceController,
	}
}

func (api *ServiceApi) AboutJson(ctx *gin.Context) {
	allServices, err := api.serviceController.AboutJson(ctx)

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
