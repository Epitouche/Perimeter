package api

import (
	"github.com/gin-gonic/gin"

	"area/controller"
)

type AreaResultApi struct {
	controller controller.AreaResultController
}

func NewAreaResultAPI(controller controller.AreaResultController, apiRoutes *gin.RouterGroup) *AreaResultApi {
	return &AreaResultApi{
		controller: controller,
	}
}
