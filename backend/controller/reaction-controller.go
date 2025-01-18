package controller

import (
	"fmt"

	"area/schemas"
	"area/service"
)

// ReactionController defines the interface for handling reactions-related operations.
// It includes methods for retrieving reactions information and specific reaction details.
type ReactionController interface {
	// GetReactionsInfo retrieves a list of reactions based on the provided ID.
	// Parameters:
	//   id - The unique identifier for the reactions.
	// Returns:
	//   response - A slice of Reaction schemas containing the reactions information.
	//   err - An error object if an error occurred during the retrieval process.
	GetReactionsInfo(id uint64) (response []schemas.Reaction, err error)

	// GetReactionByReactionID retrieves a specific reaction based on the provided reaction ID.
	// Parameters:
	//   id - The unique identifier for the reaction.
	// Returns:
	//   response - A Reaction schema containing the reaction details.
	//   err - An error object if an error occurred during the retrieval process.
	GetReactionByReactionID(id uint64) (response schemas.Reaction, err error)
}

// reactionController is a controller that handles HTTP requests related to reactions.
// It uses a ReactionService to perform business logic operations.
type reactionController struct {
	service service.ReactionService
}

// NewReactionController creates a new instance of ReactionController with the provided ReactionService.
// It returns a pointer to the newly created reactionController struct.
//
// Parameters:
//   - service: an instance of ReactionService that will be used by the ReactionController.
//
// Returns:
//   - ReactionController: a new instance of ReactionController.
func NewReactionController(service service.ReactionService) ReactionController {
	return &reactionController{
		service: service,
	}
}

// GetReactionsInfo retrieves the reaction information for a given ID.
// It returns a slice of Reaction schemas and an error if the operation fails.
//
// Parameters:
//   - id: The unique identifier for the reactions to be retrieved.
//
// Returns:
//   - response: A slice of Reaction schemas containing the reaction information.
//   - err: An error if there is an issue retrieving the reaction information.
func (controller *reactionController) GetReactionsInfo(
	id uint64,
) (response []schemas.Reaction, err error) {
	response, err = controller.service.GetReactionsInfo(id)
	if err != nil {
		return nil, fmt.Errorf("can't get reactions info: %w", err)
	}
	return response, nil
}

// GetReactionByReactionID retrieves a reaction by its ID.
// It takes a uint64 ID as input and returns a schemas.Reaction response and an error.
// If the reaction is found, it returns the reaction and a nil error.
// If there is an error during the retrieval, it returns an empty reaction and an error
// with a message indicating the failure reason.
func (controller *reactionController) GetReactionByReactionID(
	id uint64,
) (response schemas.Reaction, err error) {
	response, err = controller.service.FindById(id)
	if err != nil {
		return response, fmt.Errorf("unable to get actions info because %w", err)
	}
	return response, nil
}
