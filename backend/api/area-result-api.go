package api

import (
	"github.com/gin-gonic/gin"

	"github.com/Epitouche/Perimeter/controller"
)

type AreaResultApi struct {
	controller controller.AreaResultController
}

func NewAreaResultAPI(
	controller controller.AreaResultController,
	apiRoutes *gin.RouterGroup,
) *AreaResultApi {
	return &AreaResultApi{
		controller: controller,
	}
}
