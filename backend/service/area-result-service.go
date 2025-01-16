package service

import (
	"area/repository"
	"area/schemas"
)

// AreaResultService defines the interface for managing area results.
// It provides methods to save a new area result, retrieve all area results,
// and find area results by a specific area ID.
type AreaResultService interface {
	// Save stores a new area result.
	// Parameters:
	//   newAreaResult - the area result to be saved.
	Save(newAreaResult schemas.AreaResult)

	// FindAll retrieves all area results.
	// Returns:
	//   A slice of all area results.
	FindAll() []schemas.AreaResult

	// FindByAreaID retrieves area results by a specific area ID.
	// Parameters:
	//   areaID - the ID of the area to find results for.
	// Returns:
	//   A slice of area results for the specified area ID.
	FindByAreaID(areaID uint64) []schemas.AreaResult
}

// areaResultService is a service that provides operations related to area results.
// It interacts with the AreaResultRepository to perform CRUD operations and business logic.
type areaResultService struct {
	repository repository.AreaResultRepository
}

// NewAreaResultService creates a new instance of AreaResultService with the provided repository.
// It initializes an areaResultService struct with the given repository and returns a pointer to it.
//
// Parameters:
//   - repository: an instance of AreaResultRepository that will be used by the service.
//
// Returns:
//   - AreaResultService: a new instance of AreaResultService.
func NewAreaResultService(
	repository repository.AreaResultRepository,
) AreaResultService {
	newService := areaResultService{
		repository: repository,
	}
	return &newService
}

// Save stores a new AreaResult in the repository.
//
// Parameters:
//
//	newAreaResult - the AreaResult schema to be saved.
func (service *areaResultService) Save(newAreaResult schemas.AreaResult) {
	service.repository.Save(newAreaResult)
}

// FindAll retrieves all area results from the repository.
// It returns a slice of AreaResult schemas.
func (service *areaResultService) FindAll() []schemas.AreaResult {
	return service.repository.FindAll()
}

// FindByAreaID retrieves a list of AreaResult schemas based on the provided area ID.
//
// Parameters:
//   - areaID: The unique identifier of the area.
//
// Returns:
//   - A slice of AreaResult schemas that match the given area ID.
func (service *areaResultService) FindByAreaID(areaID uint64) []schemas.AreaResult {
	return service.repository.FindByAreaId(areaID)
}
