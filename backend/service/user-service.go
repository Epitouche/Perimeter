package service

import (
	"fmt"
	"regexp"
	"strconv"

	"area/database"
	"area/repository"
	"area/schemas"
)

// UserService defines the interface for user-related operations.
// It includes methods for user authentication, registration, information retrieval, and updates.
//
// Methods:
// - Login: Authenticates a user and returns a JWT token, user ID, and an error if any.
// - Register: Registers a new user and returns a JWT token, user ID, and an error if any.
// - GetUserInfo: Retrieves user information based on a provided token and returns the user info and an error if any.
// - UpdateUserInfo: Updates the information of an existing user and returns an error if any.
// - GetUserById: Retrieves user information based on a provided user ID and returns the user info and an error if any.
// - DeleteUser: Deletes an existing user and returns an error if any.
type UserService interface {
	Login(user schemas.User) (jwtToken string, userID uint64, err error)
	Register(newUser schemas.User) (jwtToken string, userID uint64, err error)
	GetUserInfo(token string) (userInfo schemas.User, err error)
	UpdateUserInfo(newUser schemas.User) (err error)
	GetUserById(userID uint64) (user schemas.User, err error)
	DeleteUser(newUser schemas.User) (err error)
}

// userService is a struct that provides user-related services.
// It contains the following fields:
// - authorizedUsername: a string representing the authorized username.
// - authorizedPassword: a string representing the authorized password.
// - repository: an instance of UserRepository for accessing user data.
// - serviceJWT: an instance of JWTService for handling JWT operations.
type userService struct {
	authorizedUsername string
	authorizedPassword string
	repository         repository.UserRepository
	serviceJWT         JWTService
}

// NewUserService creates a new instance of UserService with the provided
// UserRepository and JWTService. It initializes the userService with default
// authorized credentials and returns the UserService interface.
//
// Parameters:
//   - userRepository: an instance of UserRepository to interact with user data.
//   - serviceJWT: an instance of JWTService to handle JWT operations.
//
// Returns:
//   - UserService: an interface representing the user service.
func NewUserService(userRepository repository.UserRepository, serviceJWT JWTService) UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository:         userRepository,
		serviceJWT:         serviceJWT,
	}
}

// isValidEmail checks if the provided email string is in a valid email format.
// It uses a regular expression to match the email pattern.
// The function returns true if the email is valid, otherwise false.
//
// Parameters:
//   - email: A string representing the email address to be validated.
//
// Returns:
//   - bool: true if the email is valid, false otherwise.
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// Login authenticates a user based on the provided credentials.
// It first attempts to find the user by username. If no user is found, it returns an error indicating invalid credentials.
// If a user is found, it checks if the provided password matches the stored password for regular users.
// If the password matches, it generates and returns a JWT token along with the user ID.
// If the user is an OAuth2.0 user, it checks if the email matches and if the user has a valid token ID.
// If both conditions are met, it generates and returns a JWT token along with the user ID.
// If no valid user is found, it returns an error indicating the user was not found.
//
// Parameters:
// - newUser: schemas.User - The user credentials for login.
//
// Returns:
// - jwtToken: string - The generated JWT token if authentication is successful.
// - userID: uint64 - The ID of the authenticated user.
// - err: error - An error if authentication fails.
func (service *userService) Login(
	newUser schemas.User,
) (jwtToken string, userID uint64, err error) {
	userWiththisUserName, err := service.repository.FindByUserName(newUser.Username)
	if err != nil {
		return "", 0, err
	}
	if len(userWiththisUserName) == 0 {
		return "", 0, schemas.ErrInvalidCredentials
	}
	// regular user
	for _, user := range userWiththisUserName {
		if database.DoPasswordsMatch(user.Password, newUser.Password) {
			return service.serviceJWT.GenerateToken(
				strconv.FormatUint(user.Id, 10),
				user.Username,
				false,
			), user.Id, nil
		}
	}

	// Oauth2.0 user
	for _, user := range userWiththisUserName {
		if user.Email == newUser.Email {
			if user.TokenId != 0 {
				return service.serviceJWT.GenerateToken(
					strconv.FormatUint(user.Id, 10),
					user.Username,
					false,
				), user.Id, nil
			}
		}
	}

	return "", 0, schemas.ErrUserNotFound
}

// Register registers a new user in the system.
// It takes a schemas.User object as input and returns a JWT token, the user ID, and an error if any occurred.
//
// Parameters:
//   - newUser: schemas.User - The user object containing the new user's details.
//
// Returns:
//   - jwtToken: string - The generated JWT token for the new user.
//   - userID: uint64 - The ID of the newly registered user.
//   - err: error - An error object if any error occurred during the registration process.
//
// Possible errors:
//   - schemas.ErrEmailAlreadyExist: If a user with the same email already exists.
//   - schemas.ErrInvalidEmail: If the provided email is not valid.
//   - schemas.ErrHashingPassword: If there was an error hashing the user's password.
//   - Other errors returned by the repository methods.
func (service *userService) Register(
	newUser schemas.User,
) (jwtToken string, userID uint64, err error) {
	// email validation
	if !isValidEmail(newUser.Email) {
		return "", 0, schemas.ErrInvalidEmail
	}

	userWiththisEmail, err := service.repository.FindByEmail(newUser.Email)
	if err != nil {
		return "", 0, err
	}

	if len(userWiththisEmail) != 0 {
		return "", 0, schemas.ErrEmailAlreadyExist
	}

	println(newUser.Username)

	// username validation
	userWiththisUserName, err := service.repository.FindByUserName(newUser.Username)
	if err != nil {
		return "", 0, err
	}

	if len(userWiththisUserName) != 0 {
		return "", 0, schemas.ErrUsernameAlreadyExist
	}

	// store user
	if newUser.Password != "" {
		hashedPassword, err := database.HashPassword(newUser.Password)
		if err != nil {
			return "", 0, schemas.ErrHashingPassword
		}
		newUser.Password = hashedPassword
	}

	err = service.repository.Save(newUser)
	if err != nil {
		return "", 0, err
	}

	userTemp, err := service.repository.FindByUserName(newUser.Username)
	if err != nil {
		return "", 0, err
	}
	newUser.Id = userTemp[0].Id

	return service.serviceJWT.GenerateToken(
		strconv.FormatUint(newUser.Id, 10),
		newUser.Username,
		false,
	), newUser.Id, nil
}

// GetUserInfo retrieves user information based on the provided JWT token.
// It first extracts the user ID from the token using the serviceJWT.GetUserIdfromJWTToken method.
// Then, it fetches the user information from the repository using the extracted user ID.
// If any error occurs during these operations, it returns an empty user object and the error.
//
// Parameters:
//   - token: A string representing the JWT token.
//
// Returns:
//   - userInfo: A schemas.User object containing the user information.
//   - err: An error object if any error occurs during the process.
func (service *userService) GetUserInfo(token string) (userInfo schemas.User, err error) {
	userId, err := service.serviceJWT.GetUserIdfromJWTToken(token)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get user info because %w", err)
	}
	userInfo, err = service.repository.FindById(userId)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to get user info because %w", err)
	}
	return userInfo, nil
}

// UpdateUserInfo updates the information of an existing user in the repository.
// It takes a schemas.User object as input, which contains the new user information.
// It returns an error if the update operation fails.
func (service *userService) UpdateUserInfo(newUser schemas.User) (err error) {
	return service.repository.Update(newUser)
}

// GetUserById retrieves a user by their unique ID.
// Parameters:
//   - userID: The unique identifier of the user to retrieve.
//
// Returns:
//   - user: The user object corresponding to the provided ID.
//   - err: An error object if an error occurred during the retrieval process.
func (service *userService) GetUserById(userID uint64) (user schemas.User, err error) {
	return service.repository.FindById(userID)
}

// DeleteUser deletes a user from the repository.
// It takes a schemas.User object as input and returns an error if the deletion fails.
//
// Parameters:
//
//	newUser (schemas.User): The user object to be deleted.
//
// Returns:
//
//	error: An error object if the deletion fails, otherwise nil.
func (service *userService) DeleteUser(newUser schemas.User) (err error) {
	return service.repository.Delete(newUser)
}
