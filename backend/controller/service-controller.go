package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

// ServiceController defines the interface for handling service-related operations.
// It includes methods for retrieving information about services and providing
// JSON responses for the "About" endpoint.
type ServiceController interface {
	// AboutJSON returns information about the service in JSON format.
	// It takes a gin.Context as a parameter and returns an AboutJSON schema and an error.
	AboutJSON(ctx *gin.Context) (aboutJSON schemas.AboutJSON, err error)

	// GetServicesInfo retrieves information about all available services.
	// It returns a slice of Service schemas and an error.
	GetServicesInfo() (response []schemas.Service, err error)

	// GetServiceInfoById retrieves information about a specific service by its ID.
	// It takes a uint64 ID as a parameter and returns a Service schema and an error.
	GetServiceInfoById(id uint64) (response schemas.Service, err error)
}

// serviceController is a struct that holds references to various service interfaces.
// It includes:
// - service: an instance of ServiceService which provides service-related operations.
// - serviceAction: an instance of ActionService which handles actions related to services.
// - serviceReaction: an instance of ReactionService which manages reactions to service events.
type serviceController struct {
	service         service.ServiceService
	serviceAction   service.ActionService
	serviceReaction service.ReactionService
}

// NewServiceController creates a new instance of ServiceController with the provided services.
// Parameters:
//   - service: an instance of ServiceService to handle service-related operations.
//   - serviceAction: an instance of ActionService to handle action-related operations.
//   - serviceReaction: an instance of ReactionService to handle reaction-related operations.
//
// Returns:
//   - A new instance of ServiceController.
func NewServiceController(
	service service.ServiceService,
	serviceAction service.ActionService,
	serviceReaction service.ReactionService,
) ServiceController {
	return &serviceController{
		service:         service,
		serviceAction:   serviceAction,
		serviceReaction: serviceReaction,
	}
}

// AboutJSON generates a JSON response containing information about the client and server.
// It retrieves all services, formats them into JSON, and includes the client's IP address
// and the current server time in the response.
//
// Parameters:
//
//	ctx - the Gin context for the request
//
// Returns:
//
//	aboutJSON - a structured JSON response containing client and server information
//	err - an error object if any issues occur during processing
func (controller *serviceController) AboutJSON(
	ctx *gin.Context,
) (aboutJSON schemas.AboutJSON, err error) {
	allServicesJSON := []schemas.ServiceJSON{}
	allServices := controller.service.FindAll()

	for _, oneService := range allServices {
		allServicesJSON = append(allServicesJSON, schemas.ServiceJSON{
			Name:     schemas.ServiceName(oneService.Name),
			Action:   controller.serviceAction.GetAllServicesByServiceId(oneService.Id),
			Reaction: controller.serviceReaction.GetAllServicesByServiceId(oneService.Id),
		})
	}
	aboutJSON.Client.Host = ctx.ClientIP()
	aboutJSON.Server.CurrentTime = strconv.FormatInt(time.Now().Unix(), 10)
	aboutJSON.Server.Services = allServicesJSON
	return aboutJSON, nil
}

// GetServicesInfo retrieves information about services.
// It returns a slice of Service schemas and an error if the operation fails.
// If an error occurs, it wraps the original error with additional context.
func (controller *serviceController) GetServicesInfo() (response []schemas.Service, err error) {
	response, err = controller.service.GetServicesInfo()
	if err != nil {
		return nil, fmt.Errorf("can't get services info: %w", err)
	}
	return response, nil
}

// GetServiceInfoById retrieves the service information based on the provided service ID.
//
// Parameters:
//
//	id (uint64): The unique identifier of the service.
//
// Returns:
//
//	response (schemas.Service): The service information corresponding to the provided ID.
//	err (error): An error if the service information could not be retrieved.
func (controller *serviceController) GetServiceInfoById(
	id uint64,
) (response schemas.Service, err error) {
	response = controller.service.GetServiceById(id)
	return response, nil
}
