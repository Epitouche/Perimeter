package service

import (
	"area/repository"
	"area/schemas"
)

// ReactionService defines the interface for managing reactions.
// It provides methods to find, save, and retrieve reaction data.
type ReactionService interface {
	// FindAll retrieves all reactions.
	// Returns a slice of Reaction and an error if any.
	FindAll() (reactions []schemas.Reaction, err error)

	// SaveAllReaction saves all reactions.
	SaveAllReaction()

	// FindById retrieves a reaction by its ID.
	// It takes a reactionId of type uint64 as input and returns a schemas.Reaction and an error.
	// If the reaction is found, it returns the reaction and a nil error.
	// If an error occurs during the retrieval, it returns an empty reaction and the error.
	// Takes a reaction ID as input and returns the corresponding Reaction and an error if any.
	FindById(reactionId uint64) (reaction schemas.Reaction, err error)

	// GetReactionsInfo retrieves reaction information by ID.
	// Takes an ID as input and returns a slice of Reaction and an error if any.
	GetReactionsInfo(id uint64) (response []schemas.Reaction, err error)

	// GetAllServicesByServiceId retrieves all services by service ID.
	// Takes a service ID as input and returns a slice of ReactionJSON.
	GetAllServicesByServiceId(serviceId uint64) (reactionJSON []schemas.ReactionJSON)
}

// ServiceReaction defines the interface for retrieving reaction information.
// It contains a single method GetServiceReactionInfo which returns a slice of Reaction schemas.
type ServiceReaction interface {
	GetServiceReactionInfo() []schemas.Reaction
}

// reactionService is a struct that provides methods to interact with the reaction repository
// and other related services.
//
// Fields:
// - repository: An instance of ReactionRepository used to perform CRUD operations on reactions.
// - serviceService: An instance of ServiceService used to interact with other services.
type reactionService struct {
	repository     repository.ReactionRepository
	serviceService ServiceService
}

// NewReactionService creates a new instance of ReactionService with the provided
// ReactionRepository and ServiceService. It initializes the service by calling
// SaveAllReaction method and returns the newly created ReactionService.
//
// Parameters:
//   - repository: an instance of ReactionRepository to interact with reaction data.
//   - serviceService: an instance of ServiceService to provide additional services.
//
// Returns:
//   - ReactionService: a new instance of ReactionService.
func NewReactionService(
	repository repository.ReactionRepository,
	serviceService ServiceService,
) ReactionService {
	newService := &reactionService{
		repository:     repository,
		serviceService: serviceService,
	}
	newService.SaveAllReaction()
	return newService
}

// FindAll retrieves all reactions from the repository.
// It returns a slice of Reaction schemas and an error if any occurs during the retrieval process.
func (service *reactionService) FindAll() (reactions []schemas.Reaction, err error) {
	reactions, err = service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return reactions, nil
}

// GetAllServicesByServiceId retrieves all reactions associated with a given service ID.
// It takes a service ID as input and returns a slice of ReactionJSON objects.
// If an error occurs while fetching the reactions, it prints an error message.
//
// Parameters:
//   - serviceId: The ID of the service for which reactions are to be retrieved.
//
// Returns:
//   - reactionJSON: A slice of ReactionJSON objects containing the name and description of each reaction.
func (service *reactionService) GetAllServicesByServiceId(
	serviceId uint64,
) (reactionJSON []schemas.ReactionJSON) {
	allRectionForService, err := service.repository.FindByServiceId(serviceId)
	if err != nil {
		println("Error when get all reactions by service id")
	}
	for _, oneReaction := range allRectionForService {
		reactionJSON = append(reactionJSON, schemas.ReactionJSON{
			Name:        oneReaction.Name,
			Description: oneReaction.Description,
		})
	}
	return reactionJSON
}

// SaveAllReaction iterates over all services obtained from the serviceService,
// checks if they implement the ServiceReaction interface, and if so, processes
// their reactions. For each reaction, it attempts to find an existing reaction
// by name in the repository. If no existing reaction is found, it saves the new
// reaction to the repository. Errors encountered during these operations are
// logged to the console.
func (service *reactionService) SaveAllReaction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceReaction, ok := services.(ServiceReaction); ok {
			reactions := serviceReaction.GetServiceReactionInfo()
			for _, reaction := range reactions {
				reactionByName, err := service.repository.FindByName(reaction.Name)
				if err != nil {
					println("Error when find reaction by name")
				}
				if len(reactionByName) == 0 {
					err = service.repository.Save(reaction)
					if err != nil {
						println("Error when save reaction")
					}
				}
			}
		} else {
			println("ServiceReaction interface not implemented")
		}
	}
}

func (service *reactionService) FindById(reactionId uint64) (reaction schemas.Reaction, err error) {
	reaction, err = service.repository.FindById(reactionId)
	if err != nil {
		return reaction, err
	}
	return reaction, nil
}

// GetReactionsInfo retrieves the reactions information for a given service ID.
// It returns a slice of Reaction schemas and an error if any occurred during the retrieval.
//
// Parameters:
//   - id: The unique identifier of the service.
//
// Returns:
//   - response: A slice of Reaction schemas containing the reactions information.
//   - err: An error if any occurred during the retrieval process.
func (service *reactionService) GetReactionsInfo(
	id uint64,
) (response []schemas.Reaction, err error) {
	response, err = service.repository.FindByServiceId(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
