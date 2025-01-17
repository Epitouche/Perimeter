package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

// ReactionRepository defines the interface for managing reactions in the repository.
// It provides methods to save, update, delete, and retrieve reactions based on various criteria.
type ReactionRepository interface {
	// Save stores a new reaction in the repository.
	// Returns an error if the operation fails.
	Save(reaction schemas.Reaction) error

	// Update modifies an existing reaction in the repository.
	// Returns an error if the operation fails.
	Update(reaction schemas.Reaction) error

	// Delete removes a reaction from the repository.
	// Returns an error if the operation fails.
	Delete(reaction schemas.Reaction) error

	// FindAll retrieves all reactions from the repository.
	// Returns a slice of reactions and an error if the operation fails.
	FindAll() (reactions []schemas.Reaction, err error)

	// FindByName retrieves reactions from the repository that match the given action name.
	// Returns a slice of reactions and an error if the operation fails.
	FindByName(actionName string) (reactions []schemas.Reaction, err error)

	// FindByServiceId retrieves reactions from the repository that match the given service ID.
	// Returns a slice of reactions and an error if the operation fails.
	FindByServiceId(serviceId uint64) (reactions []schemas.Reaction, err error)

	// FindByServiceByName retrieves reactions from the repository that match the given service ID and action name.
	// Returns a slice of reactions and an error if the operation fails.
	FindByServiceByName(
		serviceID uint64,
		actionName string,
	) (reactions []schemas.Reaction, err error)

	// FindById retrieves a reaction from the repository that matches the given action ID.
	// Returns the reaction and an error if the operation fails.
	FindById(actionId uint64) (reaction schemas.Reaction, err error)
}

// reactionRepository is a struct that provides methods to interact with the reactions data in the database.
// It holds a reference to the Database schema which is used to perform database operations.
type reactionRepository struct {
	db *schemas.Database
}

// NewReactionRepository creates a new instance of ReactionRepository.
// It performs an automatic migration for the Reaction schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
// It returns a pointer to a reactionRepository struct initialized with the given database connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - ReactionRepository: A new instance of ReactionRepository.
func NewReactionRepository(conn *gorm.DB) ReactionRepository {
	err := conn.AutoMigrate(&schemas.Reaction{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &reactionRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// Save stores a given reaction in the database.
// It returns an error if the operation fails.
//
// Parameters:
//   - reaction: the Reaction schema object to be saved.
//
// Returns:
//   - error: an error object if the save operation fails, otherwise nil.
func (repo *reactionRepository) Save(reaction schemas.Reaction) error {
	err := repo.db.Connection.Create(&reaction)
	if err.Error != nil {
		return fmt.Errorf("failed to save reaction: %w", err.Error)
	}
	return nil
}

// Update updates an existing reaction in the database.
// It takes a schemas.Reaction object as input and returns an error if the update fails.
// If the update is successful, it returns nil.
//
// Parameters:
//   - reaction: the schemas.Reaction object to be updated in the database.
//
// Returns:
//   - error: an error object if the update fails, or nil if the update is successful.
func (repo *reactionRepository) Update(reaction schemas.Reaction) error {
	err := repo.db.Connection.Save(&reaction)
	if err.Error != nil {
		return fmt.Errorf("failed to update reaction: %w", err.Error)
	}
	return nil
}

// Delete removes a reaction record from the database.
// It takes a schemas.Reaction object as input and returns an error if the deletion fails.
//
// Parameters:
//
//	reaction (schemas.Reaction): The reaction object to be deleted.
//
// Returns:
//
//	error: An error object if the deletion fails, otherwise nil.
func (repo *reactionRepository) Delete(reaction schemas.Reaction) error {
	err := repo.db.Connection.Delete(&reaction).Error
	if err != nil {
		return fmt.Errorf("failed to delete reaction: %w", err)
	}
	return nil
}

// FindAll retrieves all reactions from the database, preloading the associated Service.
// It returns a slice of Reaction schemas and an error if the operation fails.
// If the operation is successful, the error will be nil.
func (repo *reactionRepository) FindAll() (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").Find(&reactions).Error
	if err != nil {
		return reactions, fmt.Errorf("failed to find all reactions: %w", err)
	}
	return reactions, nil
}

// FindByName retrieves a list of reactions from the database that match the given action name.
// It preloads the associated Service for each reaction.
// Parameters:
//   - actionName: The name of the action to search for.
//
// Returns:
//   - reactions: A slice of Reaction structs that match the given action name.
//   - err: An error if the query fails, otherwise nil.
func (repo *reactionRepository) FindByName(
	actionName string,
) (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").
		Where(&schemas.Reaction{Name: actionName}).
		Find(&reactions).
		Error
	if err != nil {
		return reactions, fmt.Errorf("failed to find reaction by name: %w", err)
	}
	return reactions, nil
}

// FindByServiceId retrieves all reactions associated with a given service ID.
// It preloads the "Service" association and filters the reactions by the provided service ID.
// If an error occurs during the database query, the function will panic with a formatted error message.
//
// Parameters:
//   - serviceId: The ID of the service for which to find reactions.
//
// Returns:
//   - reactions: A slice of Reaction schemas associated with the given service ID.
//   - err: An error object if any issues occur during the query.
func (repo *reactionRepository) FindByServiceId(
	serviceId uint64,
) (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").Where(&schemas.Reaction{ServiceId: serviceId}).
		Find(&reactions).Error
	if err != nil {
		panic(fmt.Errorf("failed to find reaction by service id: %w", err))
	}
	return reactions, nil
}

// FindByServiceByName retrieves a list of reactions from the database
// that match the given service ID and action name. It preloads the
// associated Service for each reaction.
//
// Parameters:
//   - serviceID: The ID of the service to filter reactions by.
//   - actionName: The name of the action to filter reactions by.
//
// Returns:
//   - reactions: A slice of Reaction schemas that match the given criteria.
//   - err: An error if the query fails, otherwise nil.
//
// Note:
//   - This function will panic if an error occurs during the database query.
func (repo *reactionRepository) FindByServiceByName(
	serviceID uint64,
	actionName string,
) (reactions []schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").
		Where(&schemas.Reaction{ServiceId: serviceID, Name: actionName}).
		Find(&reactions).
		Error
	if err != nil {
		panic(fmt.Errorf("failed to find reaction by service name: %w", err))
	}
	return reactions, nil
}

// FindById retrieves a reaction from the database by its ID.
// It preloads the associated Service for the reaction.
// Parameters:
//   - actionId: the ID of the reaction to retrieve.
//
// Returns:
//   - reaction: the retrieved reaction object.
//   - err: an error object if the operation fails, otherwise nil.
func (repo *reactionRepository) FindById(actionId uint64) (reaction schemas.Reaction, err error) {
	err = repo.db.Connection.Preload("Service").
		Where(&schemas.Reaction{Id: actionId}).
		First(&reaction).
		Error
	if err != nil {
		return reaction, fmt.Errorf("failed to find reaction by id: %w", err)
	}
	return reaction, nil
}
