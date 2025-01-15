package service

import (
	"fmt"

	"area/repository"
	"area/schemas"
)

// ActionService defines the interface for managing actions.
//
// Methods:
//   - FindAll: Retrieves all actions.
//   - SaveAllAction: Saves all actions.
//   - FindById: Finds an action by its ID.
//   - GetActionsInfo: Retrieves information about actions by ID.
//   - GetAllServicesByServiceId: Retrieves all services by service ID.
type ActionService interface {
	FindAll() (actions []schemas.Action, err error)
	SaveAllAction()
	FindById(actionId uint64) (action schemas.Action, err error)
	GetActionsInfo(id uint64) (response []schemas.Action, err error)
	GetAllServicesByServiceId(serviceId uint64) (actionJSON []schemas.ActionJSON)
}

// ServiceAction defines the interface for actions related to services.
// It includes a method to retrieve information about service actions.
type ServiceAction interface {
	GetServiceActionInfo() []schemas.Action
}

// actionService is a struct that provides services related to actions.
// It contains a repository for accessing action data and a serviceService
// for interacting with other services.
type actionService struct {
	repository     repository.ActionRepository
	serviceService ServiceService
}

// NewActionService creates a new instance of ActionService with the provided
// repository and serviceService. It initializes the action service by calling
// SaveAllAction method and returns the newly created ActionService instance.
//
// Parameters:
//   - repository: an instance of ActionRepository to interact with the data layer.
//   - serviceService: an instance of ServiceService to provide additional services.
//
// Returns:
//   - ActionService: a new instance of ActionService initialized with the provided
//     repository and serviceService.
func NewActionService(
	repository repository.ActionRepository,
	serviceService ServiceService,
) ActionService {
	newActionService := &actionService{
		repository:     repository,
		serviceService: serviceService,
	}
	newActionService.SaveAllAction()

	return newActionService
}

// FindAll retrieves all actions from the repository.
// It returns a slice of Action schemas and an error if any occurs during the retrieval process.
// If an error occurs, it wraps the original error with additional context.
func (service *actionService) FindAll() (actions []schemas.Action, err error) {
	actions, err = service.repository.FindAll()
	if err != nil {
		return actions, fmt.Errorf("error when get all actions: %w", err)
	}
	return actions, nil
}

// GetAllServicesByServiceId retrieves all actions associated with a given service ID.
// It takes a service ID as input and returns a slice of ActionJSON objects.
// If an error occurs while fetching the actions, it logs an error message.
//
// Parameters:
//   - serviceId: The ID of the service for which actions are to be retrieved.
//
// Returns:
//   - actionJSON: A slice of ActionJSON objects containing the name and description of each action.
func (service *actionService) GetAllServicesByServiceId(
	serviceId uint64,
) (actionJSON []schemas.ActionJSON) {
	allActionForService, err := service.repository.FindByServiceId(serviceId)
	if err != nil {
		fmt.Println("Error when get all actions by service id")
	}
	for _, oneAction := range allActionForService {
		actionJSON = append(actionJSON, schemas.ActionJSON{
			Name:        oneAction.Name,
			Description: oneAction.Description,
		})
	}

	return actionJSON
}

// SaveAllAction iterates through all services obtained from the serviceService,
// checks if each service implements the ServiceAction interface, and retrieves
// the service action information. For each action, it attempts to find an existing
// action by name in the repository. If no action is found, it saves the new action
// to the repository. Errors encountered during these operations are logged to the console.
func (service *actionService) SaveAllAction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceAction, ok := services.(ServiceAction); ok {
			actions := serviceAction.GetServiceActionInfo()
			for _, action := range actions {
				actionByName, err := service.repository.FindByName(action.Name)
				if err != nil {
					fmt.Println("Error when get action by name")
				}
				if len(actionByName) == 0 {
					err = service.repository.Save(action)
					if err != nil {
						fmt.Println("Error when save action")
					}
				}
			}
		} else {
			fmt.Println("Service is not ServiceAction")
		}
	}
}

// FindById retrieves an action by its ID from the repository.
// It returns the action and an error if any occurred during the retrieval process.
//
// Parameters:
//   - actionId: The ID of the action to be retrieved.
//
// Returns:
//   - action: The retrieved action.
//   - err: An error if any occurred during the retrieval process.
func (service *actionService) FindById(actionId uint64) (action schemas.Action, err error) {
	action, err = service.repository.FindById(actionId)
	if err != nil {
		return action, fmt.Errorf("error when get action by id: %w", err)
	}
	return action, nil
}

// GetActionsInfo retrieves the actions information for a given service ID.
// It returns a slice of Action schemas and an error if any occurred during the retrieval process.
//
// Parameters:
//   - id: The ID of the service for which to retrieve actions information.
//
// Returns:
//   - response: A slice of Action schemas containing the actions information.
//   - err: An error if any occurred during the retrieval process.
func (service *actionService) GetActionsInfo(id uint64) (response []schemas.Action, err error) {
	response, err = service.repository.FindByServiceId(id)
	if err != nil {
		return response, fmt.Errorf("error when get actions info: %w", err)
	}
	return response, nil
}
