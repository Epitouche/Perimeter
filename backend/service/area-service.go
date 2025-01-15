package service

import (
	"encoding/json"
	"fmt"
	"reflect"

	"area/repository"
	"area/schemas"
)

type AreaService interface {
	FindAll() (areas []schemas.Area, err error)
	CreateArea(result schemas.AreaMessage, token string) (string, error)
	InitArea(areaStartValue schemas.Area)
	AreaExist(id uint64) bool
	GetUserAreas(token string) ([]schemas.Area, error)
	UpdateUserArea(token string, areaToUpdate schemas.Area) (updatedArea schemas.Area, err error)
	DeleteUserArea(
		token string,
		areaToDelete struct{ Id uint64 },
	) (deletedArea schemas.Area, err error)
}

// areaService is a struct that provides various services related to areas.
// It includes the following fields:
// - repository: An instance of AreaRepository for data access operations.
// - actionService: An instance of ActionService for handling actions.
// - reactionService: An instance of ReactionService for handling reactions.
// - serviceUser: An instance of UserService for user-related operations.
// - serviceService: An instance of ServiceService for service-related operations.
// - areaResultService: An instance of AreaResultService for handling area results.
type areaService struct {
	repository        repository.AreaRepository
	actionService     ActionService
	reactionService   ReactionService
	serviceUser       UserService
	serviceService    ServiceService
	areaResultService AreaResultService
}

// NewAreaService creates a new instance of AreaService with the provided dependencies.
// It initializes the areaService struct with the given repository, actionService,
// reactionService, serviceUser, serviceService, and areaResultService.
//
// Parameters:
//   - repository: an instance of AreaRepository for data access.
//   - serviceService: an instance of ServiceService for service-related operations.
//   - actionService: an instance of ActionService for action-related operations.
//   - reactionService: an instance of ReactionService for reaction-related operations.
//   - serviceUser: an instance of UserService for user-related operations.
//   - areaResultService: an instance of AreaResultService for area result-related operations.
//
// Returns:
//   - AreaService: a new instance of AreaService initialized with the provided dependencies.
func NewAreaService(
	repository repository.AreaRepository,
	serviceService ServiceService,
	actionService ActionService,
	reactionService ReactionService,
	serviceUser UserService,
	areaResultService AreaResultService,
) AreaService {
	newService := areaService{
		repository:        repository,
		actionService:     actionService,
		reactionService:   reactionService,
		serviceUser:       serviceUser,
		serviceService:    serviceService,
		areaResultService: areaResultService,
	}
	return &newService
}

// FindAll retrieves all areas from the repository.
// It returns a slice of Area schemas and an error if any occurs during the retrieval process.
// If an error occurs, it wraps the original error with additional context.
func (service *areaService) FindAll() (areas []schemas.Area, err error) {
	areas, err = service.repository.FindAll()
	if err != nil {
		return areas, fmt.Errorf("error when get all areas: %w", err)
	}
	return areas, nil
}

// compareMaps compares two maps with string keys and interface{} values.
// It returns true if both maps have the same length and the same keys with values of the same type.
// If the lengths of the maps are different or if any key in map1 does not exist in map2
// or if the types of corresponding values are different, it returns false.
//
// Parameters:
//   - map1: The first map to compare.
//   - map2: The second map to compare.
//
// Returns:
//   - bool: true if the maps are equal in length and have the same keys with values of the same type, false otherwise.
func compareMaps(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		value2, ok := map2[key]
		if !ok || reflect.TypeOf(value1) != reflect.TypeOf(value2) {
			return false
		}
	}
	return true
}

// CreateArea creates a new area with the provided action and reaction options.
// It validates the provided options against the default options for the specified action and reaction.
// If the options are valid, it saves the new area to the repository and initializes it.
//
// Parameters:
//   - result: schemas.AreaMessage containing the action and reaction options, title, and description.
//   - token: string representing the user's authentication token.
//
// Returns:
//   - string: A success message if the area is created successfully.
//   - error: An error message if any step in the process fails.
func (service *areaService) CreateArea(result schemas.AreaMessage, token string) (string, error) {
	var actionOption, reactionOption json.RawMessage

	println("action option: ", string(result.ActionOption))
	if err := json.Unmarshal(result.ActionOption, &actionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal action option: %w", err)
	}

	if err := json.Unmarshal(result.ReactionOption, &reactionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal reaction option: %w", err)
	}

	user, err := service.serviceUser.GetUserInfo(token)
	if err != nil {
		return "", fmt.Errorf("can't get user info: %w", err)
	}

	areaAction, err := service.actionService.FindById(result.ActionId)
	if err != nil {
		return "", fmt.Errorf("can't find action by id: %w", err)
	}
	areaReaction, err := service.reactionService.FindById(result.ReactionId)
	if err != nil {
		return "", fmt.Errorf("can't find reaction by id: %w", err)
	}

	// check if the json key are the same as default areaAction.Option, json value can be different
	var defaultActionOption, providedActionOption map[string]interface{}
	if err := json.Unmarshal(areaAction.Option, &defaultActionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal default action option: %w", err)
	}
	if err := json.Unmarshal(result.ActionOption, &providedActionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal provided action option: %w", err)
	}
	if !compareMaps(defaultActionOption, providedActionOption) {
		return "", fmt.Errorf("action option does not match default option type")
	}

	var defaultReactionOption, providedReactionOption map[string]interface{}
	if err := json.Unmarshal(areaReaction.Option, &defaultReactionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal default reaction option: %w", err)
	}
	if err := json.Unmarshal(result.ReactionOption, &providedReactionOption); err != nil {
		return "", fmt.Errorf("can't unmarshal provided reaction option: %w", err)
	}
	if !compareMaps(defaultReactionOption, providedReactionOption) {
		return "", fmt.Errorf("reaction option does not match default option type")
	}

	defaultVariavle := struct{}{}
	defaultStorageVariable, err := json.Marshal(defaultVariavle)
	if err != nil {
		return "", fmt.Errorf("can't marshal default storage variable: %w", err)
	}

	newArea := schemas.Area{
		User:            user,
		ActionOption:    result.ActionOption,
		ReactionOption:  result.ReactionOption,
		Title:           result.Title,
		Description:     result.Description,
		Enable:          true,
		Action:          areaAction,
		Reaction:        areaReaction,
		StorageVariable: defaultStorageVariable,
	}

	id, error := service.repository.SaveArea(newArea)
	if error != nil {
		return "", fmt.Errorf("can't save area: %w", error)
	}

	newArea.Id = id
	service.InitArea(newArea)
	return "Area created successfully", nil
}

// AreaExist checks if an area with the given ID exists in the repository.
// It returns true if the area exists, otherwise it returns false.
//
// Parameters:
//
//	id (uint64): The ID of the area to check.
//
// Returns:
//
//	bool: True if the area exists, false otherwise.
func (service *areaService) AreaExist(id uint64) bool {
	_, err := service.repository.FindById(id)
	return err == nil
}

// InitArea initializes the area with the given start value and starts two goroutines.
// The first goroutine continuously checks if the area exists and is enabled, then performs the action associated with the area.
// The second goroutine waits for the action result from the first goroutine, performs the reaction associated with the area, and saves the result.
//
// Parameters:
//   - areaStartValue: The initial value of the area to be initialized.
//
// The function uses a channel to communicate between the two goroutines.
func (service *areaService) InitArea(areaStartValue schemas.Area) {
	channelArea := make(chan string)
	println("go routine action " + areaStartValue.Action.Name)
	println("reaction " + areaStartValue.Reaction.Name)
	go func(areaStartValue schemas.Area, channelArea chan string) {
		// get the action with the id
		for service.AreaExist(areaStartValue.Id) {
			area, err := service.repository.FindById(areaStartValue.Id)
			if err != nil {
				println("error")
				return
			}
			// println(area.Action.Name)
			action := service.serviceService.FindActionByName(area.Action.Name)
			if action == nil {
				println("action not found")
				return
			}

			if area.Enable {
				action(channelArea, area.ActionOption, area)
			}
		}
		println("clear")
		channelArea <- "response to clear"
	}(areaStartValue, channelArea)
	// area
	fmt.Printf("go routine area %+v\n", areaStartValue)
	go func(areaStartValue schemas.Area, channelArea chan string) {
		// check if the area is in the databse
		for service.AreaExist(areaStartValue.Id) {
			// check if the area is enable in the databse
			area, err := service.repository.FindById(areaStartValue.Id)
			if err != nil {
				return
			}

			reaction := service.serviceService.FindReactionByName(area.Reaction.Name)

			if area.Enable {
				resultAction := <-channelArea
				resultReaction := reaction(area.ReactionOption, area)
				service.areaResultService.Save(schemas.AreaResult{
					Area:   area,
					Result: resultReaction,
				})
				println("result action")
				println(resultAction)
				println("result reaction")
				println(resultReaction)
			}
		}
	}(areaStartValue, channelArea)
}

// containsArea checks if a given area is present in a list of areas.
// It takes a slice of schemas.Area and a single schemas.Area as input parameters.
// It returns true if the area is found in the list, otherwise it returns false.
func containsArea(areas []schemas.Area, area schemas.Area) bool {
	for _, a := range areas {
		if a.Id == area.Id {
			return true
		}
	}
	return false
}

// GetUserAreas retrieves the areas associated with a user based on the provided token.
// It first fetches the user information using the token, and then finds the areas
// associated with the user's ID.
//
// Parameters:
//   - token: A string representing the user's authentication token.
//
// Returns:
//   - []schemas.Area: A slice of Area objects associated with the user.
//   - error: An error object if there is any issue in fetching user information or areas.
func (service *areaService) GetUserAreas(token string) ([]schemas.Area, error) {
	user, err := service.serviceUser.GetUserInfo(token)
	if err != nil {
		return nil, fmt.Errorf("can't get user info: %w", err)
	}
	areas, err := service.repository.FindByUserId(user.Id)
	if err != nil {
		return nil, fmt.Errorf("can't find areas by user id: %w", err)
	}
	return areas, nil
}

// UpdateUserArea updates the area information for a user based on the provided token and area details.
// It retrieves the user information using the token, finds the user's areas, and updates the specified area if it exists.
//
// Parameters:
//   - token: A string representing the user's authentication token.
//   - areaToUpdate: A schemas.Area object containing the details of the area to be updated.
//
// Returns:
//   - updatedArea: A schemas.Area object representing the updated area.
//   - err: An error object if any error occurs during the process.
//
// Possible errors:
//   - If the user information cannot be retrieved using the token.
//   - If the user's areas cannot be found by user ID.
//   - If the area to be updated cannot be found by its ID.
//   - If the area to be updated does not belong to the user.
//   - If the area update operation fails.
func (service *areaService) UpdateUserArea(
	token string,
	areaToUpdate schemas.Area,
) (updatedArea schemas.Area, err error) {
	user, err := service.serviceUser.GetUserInfo(token)
	if err != nil {
		return updatedArea, fmt.Errorf("can't get user info: %w", err)
	}
	userArea, err := service.repository.FindByUserId(user.Id)
	if err != nil {
		return updatedArea, fmt.Errorf("can't find areas by user id: %w", err)
	}
	areaToUpdateDatabase, err := service.repository.FindById(areaToUpdate.Id)
	if err != nil {
		return updatedArea, fmt.Errorf("can't find areas by user id: %w", err)
	}
	if containsArea(userArea, areaToUpdateDatabase) {
		err = service.repository.Update(areaToUpdate)
		if err != nil {
			return updatedArea, fmt.Errorf("can't update area: %w", err)
		}
		return areaToUpdateDatabase, nil
	} else {
		return updatedArea, fmt.Errorf("area not found")
	}
}

// DeleteUserArea deletes a user area based on the provided token and area ID.
// It first retrieves the user information using the provided token, then fetches
// the areas associated with the user. If the area to be deleted is found within
// the user's areas, it deletes the area from the repository.
//
// Parameters:
//   - token: A string representing the user's authentication token.
//   - areaToDelete: A struct containing the ID of the area to be deleted.
//
// Returns:
//   - deletedArea: The deleted area if the operation is successful.
//   - err: An error if any issues occur during the process, such as failing to
//     retrieve user information, finding the area, or deleting the area.
func (service *areaService) DeleteUserArea(
	token string,
	areaToDelete struct{ Id uint64 },
) (deletedArea schemas.Area, err error) {
	user, err := service.serviceUser.GetUserInfo(token)
	if err != nil {
		return deletedArea, fmt.Errorf("can't get user info: %w", err)
	}
	userAreas, err := service.repository.FindByUserId(user.Id)
	if err != nil {
		return deletedArea, fmt.Errorf("can't find areas by user id: %w", err)
	}
	areaToDeleteDatabase, err := service.repository.FindById(areaToDelete.Id)
	if err != nil {
		return deletedArea, fmt.Errorf("can't find areas by user id: %w", err)
	}
	if containsArea(userAreas, areaToDeleteDatabase) {
		err = service.repository.Delete(areaToDeleteDatabase)
		if err != nil {
			return deletedArea, fmt.Errorf("can't update area: %w", err)
		}
		return areaToDeleteDatabase, nil
	} else {
		return deletedArea, fmt.Errorf("area not found")
	}
}
