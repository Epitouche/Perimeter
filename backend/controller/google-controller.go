package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

// GoogleController defines the interface for handling Google OAuth2 authentication and user information retrieval.
// It includes methods for redirecting to the Google OAuth2 service, handling the callback from the service,
// and retrieving user information.
//
// Methods:
// - RedirectToService(ctx *gin.Context) (oauthURL string, err error): Redirects the user to the Google OAuth2 service.
// - HandleServiceCallback(ctx *gin.Context) (string, error): Handles the callback from the Google OAuth2 service and returns an authorization code or error.
// - HandleServiceCallbackMobile(ctx *gin.Context) (string, error): Handles the callback from the Google OAuth2 service for mobile clients and returns an authorization code or error.
// - GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error): Retrieves user information using the provided context and returns user credentials or an error.
type GoogleController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
}

// gmailController is a struct that holds various service interfaces required
// for handling Google-related operations. It includes the following services:
// - GoogleService: Handles interactions with Google APIs.
// - UserService: Manages user-related operations.
// - TokenService: Manages token-related operations.
// - ServiceService: Manages additional service-related operations.
type gmailController struct {
	service        service.GoogleService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

// NewGoogleController creates a new instance of GoogleController with the provided services.
// Parameters:
//   - service: an instance of GoogleService to handle Google-specific operations.
//   - serviceUser: an instance of UserService to manage user-related operations.
//   - serviceToken: an instance of TokenService to handle token-related operations.
//   - serviceService: an instance of ServiceService to manage general service operations.
//
// Returns:
//   - GoogleController: a new instance of GoogleController configured with the provided services.
func NewGoogleController(
	service service.GoogleService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
	serviceService service.ServiceService,
) GoogleController {
	return &gmailController{
		service:        service,
		serviceUser:    serviceUser,
		serviceToken:   serviceToken,
		serviceService: serviceService,
	}
}

// RedirectToService handles the redirection to the Google OAuth service.
// It constructs the OAuth URL and returns it for the client to redirect the user.
//
// Parameters:
//   - ctx: The Gin context for the current request.
//
// Returns:
//   - oauthURL: The URL to redirect the user to for Google OAuth authentication.
//   - err: An error if the redirection URL could not be constructed.
func (controller *gmailController) RedirectToService(
	ctx *gin.Context,
) (oauthURL string, err error) {
	oauthURL, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Google,
		"https://accounts.google.com/o/oauth2/v2/auth",
		"https://mail.google.com/ profile email",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthURL, nil
}

// HandleServiceCallback handles the callback from the Google service after authentication.
// It binds the incoming request to the CodeCredentials schema, extracts the authorization code,
// and uses it to obtain a bearer token from the Google service.
//
// Parameters:
//   - ctx: The Gin context which provides request-specific information.
//
// Returns:
//   - A string containing the bearer token if successful.
//   - An error if there is any issue during the process, such as binding the credentials,
//     missing authentication code, or handling the service callback.
func (controller *gmailController) HandleServiceCallback(
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
		schemas.Google,
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

// HandleServiceCallbackMobile handles the callback from the Google service for mobile clients.
// It binds the incoming request to the MobileTokenRequest schema and retrieves the authorization header.
// Then, it calls the HandleServiceCallbackMobile method of the serviceService to process the callback.
//
// Parameters:
//   - ctx: The Gin context for the incoming request.
//
// Returns:
//   - A string containing the bearer token if successful.
//   - An error if there is any issue during the process.
func (controller *gmailController) HandleServiceCallbackMobile(
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
		schemas.Google,
		credentials,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	return bearer, err
}

// GetUserInfo retrieves user information based on the provided context.
// It extracts the authorization token from the request header, fetches the user information
// using the token, retrieves the associated token by user ID, and then fetches the Gmail user
// information using the token. The function returns the user's email and username.
//
// Parameters:
//
//	ctx - the Gin context containing the request information.
//
// Returns:
//
//	userInfo - a schemas.UserCredentials struct containing the user's email and username.
//	err - an error if any occurred during the process of retrieving user information.
func (controller *gmailController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.UserCredentials, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.serviceToken.GetTokenById(user.Id)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get token because %w", err)
	}

	gmailUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = gmailUserInfo.Email
	userInfo.Username = gmailUserInfo.Username
	return userInfo, nil
}
