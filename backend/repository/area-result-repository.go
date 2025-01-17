package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

// AreaResultRepository defines the interface for interacting with area result data.
// It provides methods to save, update, delete, and retrieve area results.
//
// Methods:
//   - Save(action schemas.AreaResult): Persists a new area result.
//   - Update(action schemas.AreaResult): Updates an existing area result.
//   - Delete(action schemas.AreaResult): Removes an area result.
//   - FindAll() []schemas.AreaResult: Retrieves all area results.
//   - FindByAreaId(userID uint64) []schemas.AreaResult: Retrieves area results by a specific area ID.
type AreaResultRepository interface {
	Save(action schemas.AreaResult)
	Update(action schemas.AreaResult)
	Delete(action schemas.AreaResult)
	FindAll() []schemas.AreaResult
	FindByAreaId(userID uint64) []schemas.AreaResult
}

// areaResultRepository is a struct that provides access to the database for area results.
// It contains a single field, db, which is a pointer to a Database schema.
type areaResultRepository struct {
	db *schemas.Database
}

// NewAreaResultRepository creates a new instance of AreaResultRepository.
// It performs an automatic migration for the AreaResult schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - An instance of AreaResultRepository.
func NewAreaResultRepository(conn *gorm.DB) AreaResultRepository {
	err := conn.AutoMigrate(&schemas.AreaResult{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &areaResultRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// Save stores the given AreaResult action in the database.
// It panics if there is an error during the creation process.
//
// Parameters:
//   - action: The AreaResult schema instance to be saved.
//
// Example:
//
//	repo.Save(action)
//
// Panics:
//   - If there is an error during the database operation.
func (repo *areaResultRepository) Save(action schemas.AreaResult) {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

// Update updates an existing AreaResult record in the database.
// It takes an AreaResult schema as input and saves the changes to the database.
// If an error occurs during the save operation, it panics with the error message.
//
// Parameters:
//   - action: schemas.AreaResult - The AreaResult schema containing the updated data.
func (repo *areaResultRepository) Update(action schemas.AreaResult) {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

// Delete removes the specified AreaResult from the database.
// It takes an AreaResult schema as a parameter and deletes the corresponding record.
// If an error occurs during the deletion process, it will panic with the error message.
//
// Parameters:
//
//	action (schemas.AreaResult): The AreaResult schema to be deleted from the database.
func (repo *areaResultRepository) Delete(action schemas.AreaResult) {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

// FindAll retrieves all AreaResult records from the database, including their associated Service records.
// It returns a slice of AreaResult structs.
// If there is an error during the database query, the function will panic with the error message.
func (repo *areaResultRepository) FindAll() []schemas.AreaResult {
	var action []schemas.AreaResult
	err := repo.db.Connection.Preload("Service").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

// FindByAreaId retrieves a list of AreaResult records from the database
// that match the given areaId. It returns a slice of AreaResult structs.
// If an error occurs during the database query, the function will panic.
//
// Parameters:
//   - areaId: The ID of the area to filter the AreaResult records.
//
// Returns:
//   - A slice of AreaResult structs that match the given areaId.
func (repo *areaResultRepository) FindByAreaId(areaId uint64) []schemas.AreaResult {
	var actions []schemas.AreaResult
	err := repo.db.Connection.Where(&schemas.AreaResult{AreaId: areaId}).
		Find(&actions)
	if err.Error != nil {
		panic(fmt.Errorf("failed to find action by service id: %w", err.Error))
	}
	return actions
}
