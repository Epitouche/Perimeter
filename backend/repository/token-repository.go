package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area/schemas"
)

// TokenRepository defines the interface for operations related to token management.
// It provides methods to save, update, delete, and retrieve tokens based on various criteria.
type TokenRepository interface {
	// Save stores a new token in the repository.
	// Returns an error if the operation fails.
	Save(token schemas.Token) error

	// Update modifies an existing token in the repository.
	// Returns an error if the operation fails.
	Update(token schemas.Token) error

	// Delete removes a token from the repository.
	// Returns an error if the operation fails.
	Delete(token schemas.Token) error

	// FindAll retrieves all tokens from the repository.
	// Returns a slice of tokens and an error if the operation fails.
	FindAll() (tokens []schemas.Token, err error)

	// FindByToken retrieves tokens that match the given token string.
	// Returns a slice of tokens and an error if the operation fails.
	FindByToken(token string) (tokens []schemas.Token, err error)

	// FindById retrieves a token by its unique identifier.
	// Returns the token and an error if the operation fails.
	FindById(id uint64) (token schemas.Token, err error)

	// FindByUserId retrieves tokens associated with a specific user ID.
	// Returns a slice of tokens and an error if the operation fails.
	FindByUserId(userID uint64) (tokens []schemas.Token, err error)

	// FindByUserIdAndServiceId retrieves a token associated with a specific user ID and service ID.
	// Returns the token and an error if the operation fails.
	FindByUserIdAndServiceId(id uint64, serviceId uint64) (token schemas.Token, err error)
}

// tokenRepository is a struct that provides methods to interact with the token data in the database.
// It holds a reference to the database connection through the db field.
type tokenRepository struct {
	db *schemas.Database
}

// NewTokenRepository creates a new instance of TokenRepository.
// It performs an automatic migration for the Token schema using the provided gorm.DB connection.
// If the migration fails, it panics with an error message.
// The function returns a TokenRepository with the initialized database connection.
//
// Parameters:
//   - conn: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - TokenRepository: An instance of TokenRepository with the database connection initialized.
func NewTokenRepository(conn *gorm.DB) TokenRepository {
	err := conn.AutoMigrate(&schemas.Token{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &tokenRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

// Save stores the given token in the database.
// It returns an error if the operation fails.
//
// Parameters:
//   - token: the token to be saved.
//
// Returns:
//   - error: an error if the token could not be saved, otherwise nil.
func (repo *tokenRepository) Save(token schemas.Token) error {
	err := repo.db.Connection.Create(&token).Error
	if err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}
	return nil
}

// Update updates an existing token in the database.
// It takes a schemas.Token object as input and returns an error if the update fails.
// If the update is successful, it returns nil.
//
// Parameters:
//   - token: the schemas.Token object to be updated in the database.
//
// Returns:
//   - error: an error object if the update fails, otherwise nil.
func (repo *tokenRepository) Update(token schemas.Token) error {
	err := repo.db.Connection.Save(&token).Error
	if err != nil {
		return fmt.Errorf("failed to update token: %w", err)
	}
	return nil
}

// Delete removes the specified token from the database.
// It returns an error if the deletion fails.
//
// Parameters:
//   - token: The token to be deleted.
//
// Returns:
//   - error: An error object if the deletion fails, otherwise nil.
func (repo *tokenRepository) Delete(token schemas.Token) error {
	err := repo.db.Connection.Delete(&token).Error
	if err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}
	return nil
}

// FindAll retrieves all tokens from the database.
// It returns a slice of Token schemas and an error if the operation fails.
// If the retrieval is successful, the error will be nil.
func (repo *tokenRepository) FindAll() (tokens []schemas.Token, err error) {
	err = repo.db.Connection.Find(&tokens).Error
	if err != nil {
		return tokens, fmt.Errorf("failed to find all tokens: %w", err)
	}
	return tokens, nil
}

// FindByToken retrieves a list of tokens from the database that match the given token string.
// It returns the list of matching tokens and an error if the operation fails.
//
// Parameters:
//   - token: The token string to search for in the database.
//
// Returns:
//   - tokens: A slice of Token objects that match the given token string.
//   - err: An error if the operation fails, otherwise nil.
func (repo *tokenRepository) FindByToken(token string) (tokens []schemas.Token, err error) {
	err = repo.db.Connection.Where(&schemas.Token{Token: token}).Find(&tokens).Error
	if err != nil {
		return tokens, fmt.Errorf("failed to find token by token: %w", err)
	}
	return tokens, nil
}

// FindById retrieves a token from the database by its ID.
// It returns the token and an error if the operation fails.
//
// Parameters:
//   - id: The ID of the token to retrieve.
//
// Returns:
//   - token: The retrieved token.
//   - err: An error if the token could not be found or another error occurred.
func (repo *tokenRepository) FindById(id uint64) (token schemas.Token, err error) {
	err = repo.db.Connection.Where(&schemas.Token{Id: id}).First(&token).Error
	if err != nil {
		return token, fmt.Errorf("failed to find token by id: %w", err)
	}
	return token, nil
}

// FindByUserId retrieves all tokens associated with a given user ID from the database.
// It preloads the associated User and Service entities for each token.
// Parameters:
//   - userID: The ID of the user whose tokens are to be retrieved.
//
// Returns:
//   - tokens: A slice of Token structs associated with the given user ID.
//   - err: An error object if the operation fails, otherwise nil.
func (repo *tokenRepository) FindByUserId(userID uint64) (tokens []schemas.Token, err error) {
	err = repo.db.Connection.
		Preload("User").
		Preload("Service").
		Where(&schemas.Token{UserId: userID}).Find(&tokens).Error
	if err != nil {
		return tokens, fmt.Errorf("failed to find token by user id: %w", err)
	}
	return tokens, nil
}

// FindByUserIdAndServiceId retrieves a token from the database based on the provided user ID and service ID.
// It returns the token if found, otherwise an error is returned.
//
// Parameters:
//   - id: The user ID to search for.
//   - serviceId: The service ID to search for.
//
// Returns:
//   - token: The token associated with the given user ID and service ID.
//   - err: An error if the token could not be found or another issue occurred during the query.
func (repo *tokenRepository) FindByUserIdAndServiceId(
	id uint64,
	serviceId uint64,
) (token schemas.Token, err error) {
	err = repo.db.Connection.Where(&schemas.Token{UserId: id, ServiceId: serviceId}).
		First(&token).
		Error
	if err != nil {
		return token, fmt.Errorf("failed to find token by user id and service id: %w", err)
	}
	return token, nil
}
