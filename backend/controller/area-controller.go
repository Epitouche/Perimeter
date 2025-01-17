package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

// AreaController defines the interface for managing user areas.
// It includes methods for creating, retrieving, updating, and deleting areas.
//
// Methods:
//   - CreateArea: Creates a new area and returns its ID or an error.
//   - GetUserAreas: Retrieves a list of areas associated with the user.
//   - UpdateUserArea: Updates an existing area and returns the updated area or an error.
//   - DeleteUserArea: Deletes an existing area and returns the deleted area or an error.
type AreaController interface {
	CreateArea(ctx *gin.Context) (string, error)
	GetUserAreas(ctx *gin.Context) (areaList []schemas.Area, err error)
	UpdateUserArea(ctx *gin.Context) (newArea schemas.Area, err error)
	DeleteUserArea(ctx *gin.Context) (newArea schemas.Area, err error)
}

// areaController is a struct that handles area-related operations.
// It uses an AreaService to perform business logic related to areas.
type areaController struct {
	service service.AreaService
}

// NewAreaController creates a new instance of AreaController with the provided AreaService.
// It returns a pointer to the newly created areaController struct.
//
// Parameters:
//   - service: An implementation of the AreaService interface.
//
// Returns:
//   - AreaController: A new instance of AreaController.
func NewAreaController(service service.AreaService) AreaController {
	return &areaController{
		service: service,
	}
}

// CreateArea handles the creation of a new area.
// It decodes the request body into an AreaMessage schema and retrieves the
// authorization token from the request header. It then calls the service layer
// to create the area with the provided data and token.
//
// Parameters:
//   - ctx: The Gin context which provides the request and response objects.
//
// Returns:
//   - A string representing the result of the area creation.
//   - An error if the request body cannot be decoded or if the service layer
//     returns an error.
func (controller *areaController) CreateArea(ctx *gin.Context) (string, error) {
	var result schemas.AreaMessage

	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	return controller.service.CreateArea(result, token)
}

// GetUserAreas retrieves the list of areas associated with the user.
// It extracts the authorization token from the request header and uses it
// to fetch the user's areas from the service.
//
// Parameters:
//
//	ctx - the Gin context which provides request-specific information.
//
// Returns:
//
//	areaList - a slice of Area schemas representing the user's areas.
//	err - an error if the operation fails, otherwise nil.
func (controller *areaController) GetUserAreas(
	ctx *gin.Context,
) (areaList []schemas.Area, err error) {
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	areaList, err = controller.service.GetUserAreas(token)
	if err != nil {
		return nil, fmt.Errorf("can't get user areas: %w", err)
	}
	return areaList, nil
}

// UpdateUserArea updates the area information for a user.
// It decodes the request body into a schemas.Area object and updates the user's area using the provided service.
// The function expects a Bearer token in the Authorization header for authentication.
//
// Parameters:
//
//	ctx - The Gin context which provides request and response handling.
//
// Returns:
//
//	newArea - The updated area information.
//	err - An error if the update process fails.
func (controller *areaController) UpdateUserArea(
	ctx *gin.Context,
) (newArea schemas.Area, err error) {
	var result schemas.Area

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return newArea, fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	newArea, err = controller.service.UpdateUserArea(token, result)
	if err != nil {
		return newArea, fmt.Errorf("can't get user areas: %w", err)
	}
	return newArea, nil
}

// DeleteUserArea handles the deletion of a user's area.
// It expects a JSON body containing the area ID and an Authorization header with a Bearer token.
// The function decodes the JSON body to extract the area ID, retrieves the Bearer token from the Authorization header,
// and calls the service layer to delete the user's area.
// If successful, it returns the updated area and nil error. Otherwise, it returns an error.
//
// Parameters:
// - ctx: The Gin context, which provides request and response handling.
//
// Returns:
// - newArea: The updated area after deletion.
// - err: An error if the deletion fails or if there are issues with decoding the request body or retrieving the token.
func (controller *areaController) DeleteUserArea(
	ctx *gin.Context,
) (newArea schemas.Area, err error) {
	var result struct{ Id uint64 }

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return newArea, fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	newArea, err = controller.service.DeleteUserArea(token, result)
	if err != nil {
		return newArea, fmt.Errorf("can't get user areas: %w", err)
	}
	return newArea, nil
}
