package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

// ServiceRepository defines the interface for managing services in the repository.
// It provides methods to save, update, delete, and retrieve services.
//
// Methods:
//   - Save(service schemas.Service) error: Persists a new service to the repository.
//   - Update(service schemas.Service) error: Updates an existing service in the repository.
//   - Delete(service schemas.Service) error: Removes a service from the repository.
//   - FindAll() (services []schemas.Service, err error): Retrieves all services from the repository.
//   - FindAllByName(name schemas.ServiceName) (services []schemas.Service, err error): Retrieves all services with the specified name from the repository.
//   - FindByName(name schemas.ServiceName) (service schemas.Service, err error): Retrieves a service with the specified name from the repository.
//   - FindById(id uint64) (service schemas.Service, err error): Retrieves a service with the specified ID from the repository.
type ServiceRepository interface {
	Save(service schemas.Service) error
	Update(service schemas.Service) error
	Delete(service schemas.Service) error
	FindAll() (services []schemas.Service, err error)
	FindAllByName(name schemas.ServiceName) (services []schemas.Service, err error)
	FindByName(name schemas.ServiceName) (service schemas.Service, err error)
	FindById(id uint64) (service schemas.Service, err error)
}

// serviceRepository is a struct that provides methods to interact with the database.
// It contains a single field, db, which is a pointer to a schemas.Database instance.
type serviceRepository struct {
	db *schemas.Database
}

// NewServiceRepository creates a new instance of ServiceRepository.
// It performs an automatic migration for the Service schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - ServiceRepository: A new instance of ServiceRepository with the database connection initialized.
func NewServiceRepository(conn *gorm.DB) ServiceRepository {
	err := conn.AutoMigrate(&schemas.Service{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &serviceRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// Save stores the given service in the database.
// It returns an error if the operation fails.
//
// Parameters:
//
//	service (schemas.Service): The service entity to be saved.
//
// Returns:
//
//	error: An error object if the save operation fails, otherwise nil.
func (repo *serviceRepository) Save(service schemas.Service) error {
	err := repo.db.Connection.Create(&service).Error
	if err != nil {
		return fmt.Errorf("failed to save service: %w", err)
	}
	return nil
}

// Update updates an existing service in the repository.
// It takes a Service schema as input and returns an error if the update fails.
// If the update is successful, it returns nil.
//
// Parameters:
//   - service: The Service schema to be updated.
//
// Returns:
//   - error: An error if the update operation fails, otherwise nil.
func (repo *serviceRepository) Update(service schemas.Service) error {
	err := repo.db.Connection.Save(&service).Error
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	return nil
}

// Delete removes the specified service from the database.
// It returns an error if the deletion fails.
//
// Parameters:
//   - service: The service entity to be deleted.
//
// Returns:
//   - error: An error object if the deletion fails, otherwise nil.
func (repo *serviceRepository) Delete(service schemas.Service) error {
	err := repo.db.Connection.Delete(&service).Error
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	return nil
}

// FindAll retrieves all services from the database.
// It returns a slice of Service schemas and an error if the operation fails.
// If the retrieval is successful, the error will be nil.
func (repo *serviceRepository) FindAll() (services []schemas.Service, err error) {
	err = repo.db.Connection.Find(&services).Error
	if err != nil {
		return services, fmt.Errorf("failed to get all services: %w", err)
	}
	return services, nil
}

// FindAllByName retrieves all services that match the given name from the database.
// It returns a slice of Service objects and an error if the operation fails.
//
// Parameters:
//   - name: The name of the services to be retrieved.
//
// Returns:
//   - services: A slice of Service objects that match the given name.
//   - err: An error if the operation fails, otherwise nil.
func (repo *serviceRepository) FindAllByName(
	name schemas.ServiceName,
) (services []schemas.Service, err error) {
	err = repo.db.Connection.Where(&schemas.Service{Name: name}).Find(&services).Error
	if err != nil {
		return services, fmt.Errorf("failed to get all services by name: %w", err)
	}
	return services, nil
}

// FindByName retrieves a service by its name from the database.
// It returns the service and an error if the operation fails.
//
// Parameters:
//   - name: the name of the service to be retrieved.
//
// Returns:
//   - service: the service object retrieved from the database.
//   - err: an error if the operation fails, otherwise nil.
func (repo *serviceRepository) FindByName(
	name schemas.ServiceName,
) (service schemas.Service, err error) {
	err = repo.db.Connection.Where(&schemas.Service{Name: name}).First(&service).Error
	if err != nil {
		return service, fmt.Errorf("failed to get service by name: %w", err)
	}
	return service, nil
}

// FindById retrieves a service by its ID from the database.
// It returns the service and an error if the operation fails.
//
// Parameters:
//   - id: The ID of the service to retrieve.
//
// Returns:
//   - service: The service with the specified ID.
//   - err: An error if the service could not be retrieved.
func (repo *serviceRepository) FindById(id uint64) (service schemas.Service, err error) {
	err = repo.db.Connection.Where(&schemas.Service{Id: id}).First(&service).Error
	if err != nil {
		return service, fmt.Errorf("failed to get service by id: %w", err)
	}
	return service, nil
}
