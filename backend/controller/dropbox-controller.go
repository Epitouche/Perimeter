package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

// DropboxController defines the interface for handling Dropbox-related operations.
// It includes methods for redirecting to the Dropbox OAuth service, handling OAuth callbacks,
// and retrieving user information, files, and folders.
//
// Methods:
// - RedirectToService(ctx *gin.Context) (oauthURL string, err error): Redirects the user to the Dropbox OAuth service.
// - HandleServiceCallback(ctx *gin.Context) (string, error): Handles the OAuth callback from Dropbox.
// - HandleServiceCallbackMobile(ctx *gin.Context) (string, error): Handles the OAuth callback from Dropbox for mobile clients.
// - GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error): Retrieves the user's Dropbox credentials.
// - GetUserFile(ctx *gin.Context) (userFile []string, err error): Retrieves the user's files from Dropbox.
// - GetUserFolder(ctx *gin.Context) (userFile []string, err error): Retrieves the user's folders from Dropbox.
type DropboxController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
	GetUserFile(ctx *gin.Context) (userFile []string, err error)
	GetUserFolder(ctx *gin.Context) (userFile []string, err error)
}

// dropboxController is a struct that holds various service dependencies
// required for handling Dropbox-related operations.
//
// Fields:
// - service: An instance of DropboxService to interact with Dropbox API.
// - serviceUser: An instance of UserService to manage user-related operations.
// - serviceToken: An instance of TokenService to handle token-related operations.
// - serviceService: An instance of ServiceService to manage service-related operations.
type dropboxController struct {
	service        service.DropboxService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

// NewDropboxController creates a new instance of DropboxController with the provided services.
// Parameters:
//   - service: an instance of DropboxService to handle Dropbox-related operations.
//   - serviceUser: an instance of UserService to handle user-related operations.
//   - serviceToken: an instance of TokenService to handle token-related operations.
//   - serviceService: an instance of ServiceService to handle service-related operations.
//
// Returns:
//   - DropboxController: a new instance of DropboxController.
func NewDropboxController(
	service service.DropboxService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
	serviceService service.ServiceService,
) DropboxController {
	return &dropboxController{
		service:        service,
		serviceUser:    serviceUser,
		serviceToken:   serviceToken,
		serviceService: serviceService,
	}
}

// RedirectToService generates an OAuth URL for Dropbox and returns it.
// It uses the serviceService to create the URL with the necessary scopes.
//
// Parameters:
// - ctx: The Gin context for the request.
//
// Returns:
// - oauthURL: The generated OAuth URL for Dropbox.
// - err: An error if the URL generation fails.
func (controller *dropboxController) RedirectToService(
	ctx *gin.Context,
) (oauthURL string, err error) {
	oauthURL, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Dropbox,
		"https://www.dropbox.com/oauth2/authorize",
		"account_info.read files.content.read files.content.write files.metadata.read profile email openid",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthURL, nil
}

// HandleServiceCallback handles the callback from the Dropbox service after authentication.
// It binds the incoming request's credentials, validates the authentication code, and processes
// the callback using the provided service methods.
//
// Parameters:
//   - ctx: The Gin context for the incoming request.
//
// Returns:
//   - A string representing the bearer token if the callback is handled successfully.
//   - An error if there is any issue during the process, such as binding credentials, missing
//     authentication code, or handling the service callback.
func (controller *dropboxController) HandleServiceCallback(
	ctx *gin.Context,
) (string, error) {
	var credentials schemas.CodeCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	code := credentials.Code
	if code == "" {
		return "", schemas.ErrMissingAuthenticationCode
	}

	authHeader := ctx.GetHeader("Authorization")

	bearer, err := controller.serviceService.HandleServiceCallback(
		code,
		authHeader,
		schemas.Dropbox,
		controller.service.AuthGetServiceAccessToken,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	if err != nil {
		return "", fmt.Errorf("unable to handle service callback because %w", err)
	}
	return bearer, nil
}

// HandleServiceCallbackMobile handles the callback from the Dropbox service for mobile clients.
// It binds the incoming request to the MobileTokenRequest schema and retrieves the Authorization header.
// Then, it calls the HandleServiceCallbackMobile method of the serviceService to process the callback.
//
// Parameters:
//   - ctx: The Gin context for the incoming request.
//
// Returns:
//   - string: The bearer token if the callback is successful.
//   - error: An error if the binding or callback handling fails.
func (controller *dropboxController) HandleServiceCallbackMobile(
	ctx *gin.Context,
) (string, error) {
	var credentials schemas.MobileTokenRequest
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")

	bearer, err := controller.serviceService.HandleServiceCallbackMobile(
		authHeader,
		schemas.Dropbox,
		credentials,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	return bearer, err
}

// GetUserInfo retrieves user information from Dropbox using the provided context.
// It expects an Authorization header with a Bearer token.
// The function performs the following steps:
// 1. Extracts the token from the Authorization header.
// 2. Retrieves user information using the token.
// 3. Retrieves the token associated with the user ID.
// 4. Retrieves Dropbox user information using the token.
// If any step fails, it returns an error with a descriptive message.
// On success, it returns the user's email and username.
func (controller *dropboxController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.UserCredentials, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return userInfo, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.serviceToken.GetTokenById(user.Id)
	if err != nil {
		return userInfo, fmt.Errorf("unable to get token because %w", err)
	}

	dropboxUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return userInfo, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = dropboxUserInfo.Email
	userInfo.Username = dropboxUserInfo.Username
	return userInfo, nil
}

// GetUserFile retrieves a list of user files from Dropbox.
//
// It extracts the authorization token from the request header, fetches user information,
// retrieves the Dropbox token associated with the user, and then fetches all folders and files
// from the user's Dropbox account. Finally, it processes and returns the list of file paths.
//
// Parameters:
//
//	ctx - The Gin context which provides request-specific information.
//
// Returns:
//
//	userFile - A slice of strings containing the paths of the user's files.
//	err - An error object if any error occurs during the process.
func (controller *dropboxController) GetUserFile(
	ctx *gin.Context,
) (userFile []string, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return userFile, fmt.Errorf("unable to get user info because %w", err)
	}

	DropboxToken, err := controller.serviceToken.GetTokenById(user.Id)
	if err != nil {
		return userFile, fmt.Errorf("unable to get token because %w", err)
	}

	dropboxAllFolderAndFileList, err := controller.service.GetUserAllFolderAndFileList(
		DropboxToken.Token,
	)
	if err != nil {
		return userFile, fmt.Errorf("unable to get user info because %w", err)
	}

	userFile = controller.service.GetPathDisplayDropboxEntry(
		controller.service.GetUserFileList(dropboxAllFolderAndFileList),
	)

	return userFile, nil
}

// GetUserFolder retrieves the list of user folders from Dropbox.
//
// It extracts the authorization token from the request header, fetches user information,
// retrieves the Dropbox token for the user, and then fetches all folders and files from Dropbox.
// Finally, it filters and returns the list of user folders.
//
// Parameters:
//
//	ctx - The Gin context which contains the request and response objects.
//
// Returns:
//
//	userFile - A slice of strings containing the paths of user folders.
//	err - An error object if any error occurs during the process.
func (controller *dropboxController) GetUserFolder(
	ctx *gin.Context,
) (userFile []string, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return userFile, fmt.Errorf("unable to get user info because %w", err)
	}

	DropboxToken, err := controller.serviceToken.GetTokenById(user.Id)
	if err != nil {
		return userFile, fmt.Errorf("unable to get token because %w", err)
	}

	dropboxAllFolderAndFileList, err := controller.service.GetUserAllFolderAndFileList(
		DropboxToken.Token,
	)
	if err != nil {
		return userFile, fmt.Errorf("unable to get user info because %w", err)
	}

	userFile = controller.service.GetPathDisplayDropboxEntry(
		controller.service.GetUserFolderList(dropboxAllFolderAndFileList),
	)

	return userFile, nil
}
