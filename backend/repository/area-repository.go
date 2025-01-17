package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

// AreaRepository defines the interface for interacting with area data.
// It provides methods to save, update, delete, and retrieve area records.
//
// Methods:
//   - SaveArea(area schemas.Area) (areaID uint64, err error): Saves a new area and returns its ID.
//   - Save(area schemas.Area) error: Saves a new area.
//   - Update(area schemas.Area) error: Updates an existing area.
//   - Delete(area schemas.Area) error: Deletes an existing area.
//   - FindAll() (areas []schemas.Area, err error): Retrieves all areas.
//   - FindByUserId(userID uint64) (areas []schemas.Area, err error): Retrieves areas associated with a specific user ID.
//   - FindById(id uint64) (area schemas.Area, err error): Retrieves an area by its ID.
type AreaRepository interface {
	SaveArea(area schemas.Area) (areaID uint64, err error)
	Save(area schemas.Area) error
	Update(area schemas.Area) error
	Delete(area schemas.Area) error
	FindAll() (areas []schemas.Area, err error)
	FindByUserId(userID uint64) (areas []schemas.Area, err error)
	FindById(id uint64) (area schemas.Area, err error)
}

// areaRepository is a struct that provides methods to interact with the area-related data in the database.
// It holds a reference to the Database schema, which is used to perform database operations.
type areaRepository struct {
	db *schemas.Database
}

// NewAreaRepository creates a new instance of AreaRepository.
// It performs an automatic migration for the Area schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
// It returns an implementation of the AreaRepository interface.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - An implementation of the AreaRepository interface.
func NewAreaRepository(conn *gorm.DB) AreaRepository {
	err := conn.AutoMigrate(&schemas.Area{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &areaRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// SaveArea saves the given area action to the repository and returns the ID of the saved area.
// If an error occurs during the save operation, it returns an error.
//
// Parameters:
//   - action: The area action to be saved.
//
// Returns:
//   - areaID: The ID of the saved area.
//   - err: An error if the save operation fails.
func (repo *areaRepository) SaveArea(action schemas.Area) (areaID uint64, err error) {
	err = repo.Save(action)
	if err != nil {
		return 0, fmt.Errorf("failed to save area: %w", err)
	}
	result := repo.db.Connection.Last(&action)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to save area: %w", err)
	}
	return action.Id, nil
}

// Save stores the given Area action in the database.
// It prints the ActionRefreshRate of the action to the console for debugging purposes.
// If the save operation fails, it returns an error indicating the failure.
//
// Parameters:
//
//	action - the Area action to be saved.
//
// Returns:
//
//	error - an error if the save operation fails, otherwise nil.
func (repo *areaRepository) Save(action schemas.Area) error {
	fmt.Printf("refreshtime: %v\n", action.ActionRefreshRate)
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		return fmt.Errorf("failed to save area: %w", err.Error)
	}
	return nil
}

// Update updates an existing area record in the database.
// It first searches for the area by its ID. If the area is found,
// it updates the record with the new data provided in the action parameter.
// If the area is not found or if there is an error during the update process,
// it returns an error.
//
// Parameters:
//   - action: schemas.Area - The area data to be updated.
//
// Returns:
//   - error: An error object if the update fails, otherwise nil.
func (repo *areaRepository) Update(action schemas.Area) error {
	var area schemas.Area
	err := repo.db.Connection.Where(&schemas.Area{Id: action.Id}).First(&area).Error
	if err != nil {
		return fmt.Errorf("failed to find area: %w", err)
	}
	if area.Id == action.Id {
		err = repo.db.Connection.Save(&action).Error
		if err != nil {
			return fmt.Errorf("failed to update area: %w", err)
		}
	}
	return nil
}

// Delete removes an area record from the database.
// It takes an Area schema as input and returns an error if the deletion fails.
//
// Parameters:
//
//	action (schemas.Area): The area schema to be deleted.
//
// Returns:
//
//	error: An error object if the deletion fails, otherwise nil.
func (repo *areaRepository) Delete(action schemas.Area) error {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		return fmt.Errorf("failed to delete area: %w", err.Error)
	}
	return nil
}

// FindAll retrieves all areas from the database, including their associated services.
// It returns a slice of Area schemas and an error if the operation fails.
// If the operation is successful, the error will be nil.
func (repo *areaRepository) FindAll() (areas []schemas.Area, err error) {
	err = repo.db.Connection.Preload("Service").Find(&areas).Error
	if err != nil {
		return areas, fmt.Errorf("failed to find all areas: %w", err)
	}
	return areas, nil
}

// FindByUserId retrieves a list of areas associated with a specific user ID.
// It preloads the User, Action.Service, and Reaction.Service associations
// to ensure related data is also fetched. If an error occurs during the
// database query, the function will panic with a formatted error message.
//
// Parameters:
//   - userID: The ID of the user whose areas are to be retrieved.
//
// Returns:
//   - areas: A slice of Area schemas associated with the given user ID.
//   - err: An error object if any issues occur during the query.
func (repo *areaRepository) FindByUserId(userID uint64) (areas []schemas.Area, err error) {
	err = repo.db.Connection.
		Preload("User").
		Preload("Action.Service").
		Preload("Reaction.Service").
		Where(&schemas.Area{UserId: userID}).
		Find(&areas).Error
	if err != nil {
		panic(fmt.Errorf("failed to find areas by user id: %w", err))
	}

	return areas, nil
}

// FindById retrieves an Area by its ID from the database, including its associated Action and Reaction.
// It returns the Area and an error if the operation fails.
//
// Parameters:
//   - id: the ID of the Area to retrieve.
//
// Returns:
//   - area: the retrieved Area, including its associated Action and Reaction.
//   - err: an error if the operation fails, otherwise nil.
func (repo *areaRepository) FindById(id uint64) (area schemas.Area, err error) {
	err = repo.db.Connection.Where(&schemas.Area{Id: id}).First(&area).Error
	var actionResult schemas.Action
	repo.db.Connection.Where(&schemas.Action{Id: area.ActionId}).First(&actionResult)
	area.Action = actionResult
	var reactionResult schemas.Reaction
	repo.db.Connection.Where(&schemas.Reaction{Id: area.ReactionId}).First(&reactionResult)
	area.Reaction = reactionResult

	if err != nil {
		return area, fmt.Errorf("failed to find action by id: %w", err)
	}
	return area, nil
}
