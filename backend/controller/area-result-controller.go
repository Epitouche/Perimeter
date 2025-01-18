package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type AreaResultController interface {
	GetUserAreaResultsByAreaID(
		ctx *gin.Context,
		areaID uint64,
	) (areaList []schemas.AreaResult, err error)
}

// areaResultController is a controller that handles requests related to area results.
// It uses AreaResultService to manage area result operations and AreaService to manage area operations.
type areaResultController struct {
	service     service.AreaResultService
	serviceArea service.AreaService
}

// NewAreaResultController creates a new instance of AreaResultController with the provided services.
// Parameters:
//   - service: an instance of AreaResultService to handle area result operations.
//   - serviceArea: an instance of AreaService to handle area-related operations.
//
// Returns:
//   - An instance of AreaResultController.
func NewAreaResultController(
	service service.AreaResultService,
	serviceArea service.AreaService,
) AreaResultController {
	return &areaResultController{
		service:     service,
		serviceArea: serviceArea,
	}
}

// GetUserAreaResultsByAreaID retrieves the area results for a specific area ID that belongs to the user.
// It extracts the authorization token from the request header, fetches the user's areas using the token,
// and then searches for the specified area ID within the user's areas. If the area is found, it returns
// the area results; otherwise, it returns an error indicating that the area was not found.
//
// Parameters:
//   - ctx: The Gin context which provides request-specific information.
//   - areaID: The ID of the area for which to retrieve results.
//
// Returns:
//   - areaResultList: A list of area results for the specified area ID.
//   - err: An error if the area is not found or if there is an issue retrieving the user's areas.
func (controller *areaResultController) GetUserAreaResultsByAreaID(
	ctx *gin.Context,
	areaID uint64,
) (areaResultList []schemas.AreaResult, err error) {
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]

	areaList, err := controller.serviceArea.GetUserAreas(token)
	if err != nil {
		return nil, fmt.Errorf("can't get user areas: %w", err)
	}

	for _, area := range areaList {
		if area.Id == areaID {
			areaResultList = controller.service.FindByAreaID(areaID)
			return areaResultList, nil
		}
	}
	return areaResultList, fmt.Errorf("area not found")
}
