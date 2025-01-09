package api

import (
	"github.com/gin-gonic/gin"

	"area/controller"
)

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
) *AreaResultApi {
	return &AreaResultApi{
		controller: controller,
	}
}
