package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

// MicrosoftController defines the interface for handling Microsoft OAuth service interactions.
// It includes methods for redirecting to the OAuth service, handling the callback from the service,
// and retrieving user information.
//
// Methods:
// - RedirectToService(ctx *gin.Context) (oauthURL string, err error): Redirects the user to the Microsoft OAuth service.
// - HandleServiceCallback(ctx *gin.Context) (string, error): Handles the callback from the Microsoft OAuth service.
// - HandleServiceCallbackMobile(ctx *gin.Context) (string, error): Handles the callback from the Microsoft OAuth service for mobile clients.
// - GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error): Retrieves user information from the Microsoft OAuth service.
type MicrosoftController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
}

// microsoftController is a struct that holds various service interfaces
// required for handling Microsoft-related operations.
//
// Fields:
// - service: An interface for Microsoft service operations.
// - serviceUser: An interface for user service operations.
// - serviceToken: An interface for token service operations.
// - serviceService: An interface for general service operations.
type microsoftController struct {
	service        service.MicrosoftService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

// NewMicrosoftController creates a new instance of MicrosoftController with the provided services.
// Parameters:
//   - service: an instance of MicrosoftService to handle Microsoft-specific operations.
//   - serviceUser: an instance of UserService to manage user-related operations.
//   - serviceToken: an instance of TokenService to handle token-related operations.
//   - serviceService: an instance of ServiceService to manage general service operations.
//
// Returns:
//   - MicrosoftController: a new instance of MicrosoftController initialized with the provided services.
func NewMicrosoftController(
	service service.MicrosoftService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
	serviceService service.ServiceService,
) MicrosoftController {
	return &microsoftController{
		service:        service,
		serviceUser:    serviceUser,
		serviceToken:   serviceToken,
		serviceService: serviceService,
	}
}

// RedirectToService generates an OAuth URL for redirecting to the Microsoft service authorization page.
// It uses the serviceService to create the URL with the necessary scopes for accessing Microsoft services.
//
// Parameters:
//   - ctx: The Gin context for the current request.
//
// Returns:
//   - oauthURL: The generated OAuth URL for the Microsoft service.
//   - err: An error if the URL generation fails.
func (controller *microsoftController) RedirectToService(
	ctx *gin.Context,
) (oauthURL string, err error) {
	oauthURL, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Microsoft,
		"https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		"Mail.ReadWrite, Mail.Read, User.Read, Mail.Send, offline_access, calendars.Read, calendars.ReadWrite",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthURL, nil
}

// HandleServiceCallback handles the callback from the Microsoft service.
// It binds the incoming request to the CodeCredentials schema and retrieves the authorization code.
// If the code is missing, it returns an error indicating the missing authentication code.
// It then retrieves the Authorization header from the request context.
// The function calls the HandleServiceCallback method of the serviceService with the necessary parameters
// to handle the service callback and obtain a bearer token.
// If successful, it returns the bearer token; otherwise, it returns an error.
//
// Parameters:
//
//	ctx - The Gin context containing the request data.
//
// Returns:
//
//	A string representing the bearer token if successful, or an error if the callback handling fails.
func (controller *microsoftController) HandleServiceCallback(
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
		schemas.Microsoft,
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

// HandleServiceCallbackMobile handles the callback from a mobile service.
// It binds the incoming request to a MobileTokenRequest schema and retrieves the
// Authorization header. It then calls the HandleServiceCallbackMobile method of the
// serviceService to process the callback.
//
// Parameters:
// - ctx: The Gin context for the request.
//
// Returns:
// - A string representing the bearer token.
// - An error if the binding or service callback handling fails.
func (controller *microsoftController) HandleServiceCallbackMobile(
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
		schemas.Microsoft,
		credentials,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	return bearer, err
}

// GetUserInfo retrieves user information based on the provided context.
// It extracts the authorization token from the request header, fetches user information
// using the token, retrieves the corresponding token from the token service, and then
// fetches the Microsoft user information using the retrieved token.
//
// Parameters:
//
//	ctx - The context of the request, which contains the authorization header.
//
// Returns:
//
//	userInfo - The user credentials containing the email and username.
//	err - An error if any step in the process fails, with a descriptive message.
func (controller *microsoftController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.UserCredentials, err error) {
	println("get user info")
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		println("error 1")
		return userInfo, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.serviceToken.GetTokenById(user.Id)
	if err != nil {
		println("error 2")
		return userInfo, fmt.Errorf("unable to get token because %w", err)
	}

	microsoftUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		println("error 3")
		return userInfo, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = microsoftUserInfo.Email
	userInfo.Username = microsoftUserInfo.Username
	return userInfo, nil
}
