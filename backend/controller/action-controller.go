package controller

import (
	"fmt"

	"area/schemas"
	"area/service"
)

// ActionController defines the interface for handling actions-related operations.
// It provides methods to retrieve information about actions and specific action details.
type ActionController interface {
	// GetActionsInfo retrieves a list of actions information based on the provided ID.
	// Parameters:
	//   - id: The unique identifier for the actions.
	// Returns:
	//   - response: A slice of Action schemas containing the actions information.
	//   - err: An error object if an error occurred during the retrieval process.
	GetActionsInfo(id uint64) (response []schemas.Action, err error)

	// GetActionByActionID retrieves the details of a specific action based on the provided action ID.
	// Parameters:
	//   - id: The unique identifier for the action.
	// Returns:
	//   - response: An Action schema containing the action details.
	//   - err: An error object if an error occurred during the retrieval process.
	GetActionByActionID(id uint64) (response schemas.Action, err error)
}

// actionController is a struct that handles actions by utilizing the provided ActionService.
// It serves as a controller in the application's architecture, managing the flow of data
// between the service layer and the client.
type actionController struct {
	service service.ActionService
}

// NewActionController creates a new instance of ActionController with the provided ActionService.
// It returns an ActionController interface that can be used to manage actions.
//
// Parameters:
//   - service: an instance of ActionService that provides the necessary business logic for actions.
//
// Returns:
//   - ActionController: an interface for managing actions.
func NewActionController(service service.ActionService) ActionController {
	return &actionController{
		service: service,
	}
}

// GetActionsInfo retrieves information about actions based on the provided ID.
// It returns a slice of Action schemas and an error if the retrieval fails.
//
// Parameters:
//   - id: The unique identifier for the actions to retrieve.
//
// Returns:
//   - response: A slice of Action schemas containing the actions information.
//   - err: An error if there is an issue retrieving the actions information.
func (controller *actionController) GetActionsInfo(
	id uint64,
) (response []schemas.Action, err error) {
	response, err = controller.service.GetActionsInfo(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get actions info because %w", err)
	}
	return response, nil
}

// GetActionByActionID retrieves an action by its ID.
// It takes an action ID as a parameter and returns the corresponding action schema and an error, if any.
// If the action is not found or there is an issue with the retrieval, an error is returned with a descriptive message.
func (controller *actionController) GetActionByActionID(
	id uint64,
) (response schemas.Action, err error) {
	response, err = controller.service.FindById(id)
	if err != nil {
		return response, fmt.Errorf("unable to get actions info because %w", err)
	}
	return response, nil
}
