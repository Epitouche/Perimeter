package repository

import (
	"gorm.io/gorm"

	"area/schemas"
)

// ActionRepository defines the interface for performing CRUD operations on actions.
// It includes methods for saving, updating, deleting, and finding actions based on various criteria.
type ActionRepository interface {
	// Save persists a new action in the repository.
	// Returns an error if the operation fails.
	Save(action schemas.Action) error

	// Update modifies an existing action in the repository.
	// Returns an error if the operation fails.
	Update(action schemas.Action) error

	// Delete removes an action from the repository.
	// Returns an error if the operation fails.
	Delete(action schemas.Action) error

	// FindAll retrieves all actions from the repository.
	// Returns a slice of actions and an error if the operation fails.
	FindAll() (action []schemas.Action, err error)

	// FindByName retrieves actions from the repository that match the given name.
	// Returns a slice of actions and an error if the operation fails.
	FindByName(actionName string) (action []schemas.Action, err error)

	// FindByServiceId retrieves actions from the repository that match the given service ID.
	// Returns a slice of actions and an error if the operation fails.
	FindByServiceId(serviceId uint64) (action []schemas.Action, err error)

	// FindById retrieves an action from the repository that matches the given action ID.
	// Returns the action and an error if the operation fails.
	FindById(actionId uint64) (action schemas.Action, err error)

	// FindByServiceByName retrieves actions from the repository that match the given service ID and action name.
	// Returns a slice of actions and an error if the operation fails.
	FindByServiceByName(serviceId uint64, actionName string) (action []schemas.Action, err error)
}

// actionRepository is a struct that provides methods to interact with the actions
// stored in the database. It holds a reference to the Database schema.
type actionRepository struct {
	db *schemas.Database
}

// NewActionRepository creates a new instance of ActionRepository.
// It performs an automatic migration for the Action schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
// The function returns an implementation of the ActionRepository interface.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - An implementation of the ActionRepository interface.
func NewActionRepository(conn *gorm.DB) ActionRepository {
	err := conn.AutoMigrate(&schemas.Action{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &actionRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// Save inserts a new action record into the database.
// It takes an Action schema as input and returns an error if the operation fails.
//
// Parameters:
//
//	action (schemas.Action): The action entity to be saved.
//
// Returns:
//
//	error: An error object if the save operation fails, otherwise nil.
func (repo *actionRepository) Save(action schemas.Action) error {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// Update updates an existing action in the database.
// It takes an action of type schemas.Action as input and returns an error if the update fails.
// If the update is successful, it returns nil.
func (repo *actionRepository) Update(action schemas.Action) error {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// Delete removes the specified action from the database.
// It takes an action of type schemas.Action as a parameter and returns an error if the deletion fails.
// If the deletion is successful, it returns nil.
func (repo *actionRepository) Delete(action schemas.Action) error {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAll retrieves all actions from the database, preloading the associated Service.
// It returns a slice of Action schemas and an error if any occurs during the database query.
func (repo *actionRepository) FindAll() (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

// FindByName retrieves a list of actions from the database that match the given action name.
// It preloads the "Service" association for each action.
// Parameters:
//   - actionName: The name of the action to search for.
//
// Returns:
//   - actions: A slice of Action structs that match the given action name.
//   - err: An error if the database query fails, otherwise nil.
func (repo *actionRepository) FindByName(actionName string) (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{Name: actionName}).
		Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

// FindByServiceId retrieves a list of actions associated with a given service ID.
// It preloads the "Service" field in the Action schema and filters the actions
// based on the provided service ID.
//
// Parameters:
//   - serviceId: The ID of the service for which actions are to be retrieved.
//
// Returns:
//   - actions: A slice of Action schemas associated with the given service ID.
//   - err: An error object if there is an issue with the database query, otherwise nil.
func (repo *actionRepository) FindByServiceId(
	serviceId uint64,
) (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{ServiceId: serviceId}).
		Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

// FindByServiceByName retrieves a list of actions associated with a specific service ID and action name.
// It preloads the "Service" relation and filters the actions based on the provided service ID and action name.
//
// Parameters:
// - serviceId: The ID of the service to filter actions by.
// - actionName: The name of the action to filter by.
//
// Returns:
// - actions: A slice of Action schemas that match the provided service ID and action name.
// - err: An error if the database query fails, otherwise nil.
func (repo *actionRepository) FindByServiceByName(
	serviceId uint64,
	actionName string,
) (actions []schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{ServiceId: serviceId, Name: actionName}).
		Find(&actions)

	if errDatabase.Error != nil {
		return actions, errDatabase.Error
	}
	return actions, nil
}

// FindById retrieves an action from the database by its ID.
// It preloads the associated Service for the action.
// Parameters:
//   - actionId: The ID of the action to retrieve.
//
// Returns:
//   - action: The retrieved action.
//   - err: An error if the retrieval fails, otherwise nil.
func (repo *actionRepository) FindById(actionId uint64) (action schemas.Action, err error) {
	errDatabase := repo.db.Connection.Preload("Service").
		Where(&schemas.Action{Id: actionId}).
		First(&action)

	if errDatabase.Error != nil {
		return action, errDatabase.Error
	}
	return action, nil
}
