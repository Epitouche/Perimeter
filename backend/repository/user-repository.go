package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

// UserRepository defines the interface for user repository operations.
// It provides methods to save, update, delete, and find users.
//
// Save saves a new user to the repository.
// Update updates an existing user in the repository.
// Delete removes a user from the repository.
// FindAll retrieves all users from the repository.
// FindByEmail finds users by their email address.
// FindByUserName finds users by their username.
// FindById finds a user by their unique ID.
type UserRepository interface {
	Save(user schemas.User) (err error)
	Update(user schemas.User) (err error)
	Delete(user schemas.User) (err error)
	FindAll() (users []schemas.User, err error)
	FindByEmail(email string) (users []schemas.User, err error)
	FindByUserName(username string) (users []schemas.User, err error)
	FindById(id uint64) (user schemas.User, err error)
}

// Define a struct that embeds `*schemas.Database` and implements `UserRepository`.
type userRepository struct {
	db *schemas.Database
}

// NewUserRepository creates a new instance of UserRepository.
// It performs an automatic migration for the User schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
// It returns a UserRepository with the initialized database connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - UserRepository: An initialized UserRepository with the database connection.
func NewUserRepository(conn *gorm.DB) UserRepository {
	err := conn.AutoMigrate(&schemas.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &userRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// Save stores a new user in the database.
// It takes a schemas.User object as input and returns an error if the operation fails.
// If the user is successfully saved, it returns nil.
//
// Parameters:
//
//	user (schemas.User): The user object to be saved.
//
// Returns:
//
//	error: An error object if the save operation fails, otherwise nil.
func (repo *userRepository) Save(user schemas.User) (err error) {
	err = repo.db.Connection.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

// Update updates an existing user in the database.
// It takes a schemas.User object as input and returns an error if the update fails.
// If the update is successful, it returns nil.
//
// Parameters:
//   - user: The user object containing updated information.
//
// Returns:
//   - error: An error object if the update fails, otherwise nil.
func (repo *userRepository) Update(user schemas.User) (err error) {
	err = repo.db.Connection.Save(&user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete removes the specified user from the database.
// It returns an error if the deletion fails.
//
// Parameters:
//   - user: The user entity to be deleted.
//
// Returns:
//   - err: An error object if the deletion fails, otherwise nil.
func (repo *userRepository) Delete(user schemas.User) (err error) {
	err = repo.db.Connection.Delete(&user).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// FindAll retrieves all user records from the database.
// It returns a slice of User schemas and an error if the operation fails.
// If the operation is successful, the error will be nil.
func (repo *userRepository) FindAll() (users []schemas.User, err error) {
	err = repo.db.Connection.Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("failed to find all users: %w", err)
	}
	return users, nil
}

// FindByEmail retrieves a list of users from the database that match the given email address.
// It returns a slice of User schemas and an error if the operation fails.
//
// Parameters:
//   - email: The email address to search for.
//
// Returns:
//   - users: A slice of User schemas that match the given email address.
//   - err: An error if the operation fails, otherwise nil.
func (repo *userRepository) FindByEmail(email string) (users []schemas.User, err error) {
	err = repo.db.Connection.Where(&schemas.User{Email: email}).Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("failed to find user by email: %w", err)
	}
	return users, nil
}

// FindByUserName retrieves a list of users from the database that match the given username.
// It returns a slice of User schemas and an error if the operation fails.
//
// Parameters:
//   - username: The username to search for in the database.
//
// Returns:
//   - users: A slice of User schemas that match the given username.
//   - err: An error if the operation fails, otherwise nil.
func (repo *userRepository) FindByUserName(username string) (users []schemas.User, err error) {
	err = repo.db.Connection.Where(&schemas.User{Username: username}).Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("failed to find user by username: %w", err)
	}
	return users, nil
}

// FindById retrieves a user from the database by their ID.
// It returns the user and an error if the user could not be found or if there was an issue with the database query.
//
// Parameters:
//   - id: The ID of the user to retrieve.
//
// Returns:
//   - user: The user retrieved from the database.
//   - err: An error if the user could not be found or if there was an issue with the database query.
func (repo *userRepository) FindById(id uint64) (user schemas.User, err error) {
	err = repo.db.Connection.Where(&schemas.User{Id: id}).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("failed to find user by id: %w", err)
	}
	return user, nil
}
