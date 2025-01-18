package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

// GithubController defines the interface for handling GitHub OAuth operations.
// It includes methods for redirecting to the GitHub OAuth service, handling
// the callback from the service, and retrieving user information.
//
// Methods:
//   - RedirectToService: Redirects the user to the GitHub OAuth service and returns the OAuth URL.
//   - HandleServiceCallback: Handles the callback from the GitHub OAuth service and returns a token or error.
//   - HandleServiceCallbackMobile: Handles the callback from the GitHub OAuth service for mobile clients and returns a token or error.
//   - GetUserInfo: Retrieves user information based on the provided context and returns user credentials or an error.
type GithubController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
}

// githubController is a struct that holds various service dependencies
// required for handling GitHub-related operations.
//
// Fields:
// - service: An instance of GithubService for interacting with GitHub APIs.
// - serviceUser: An instance of UserService for managing user-related operations.
// - serviceToken: An instance of TokenService for handling token-related operations.
// - serviceService: An instance of ServiceService for managing additional services.
type githubController struct {
	service        service.GithubService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

// NewGithubController creates a new instance of GithubController with the provided services.
// Parameters:
//   - service: an instance of GithubService to handle GitHub related operations.
//   - serviceUser: an instance of UserService to handle user related operations.
//   - serviceToken: an instance of TokenService to handle token related operations.
//   - serviceService: an instance of ServiceService to handle service related operations.
//
// Returns:
//   - GithubController: a new instance of GithubController.
func NewGithubController(
	service service.GithubService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
	serviceService service.ServiceService,
) GithubController {
	return &githubController{
		service:        service,
		serviceUser:    serviceUser,
		serviceToken:   serviceToken,
		serviceService: serviceService,
	}
}

// RedirectToService handles the redirection to the GitHub OAuth authorization page.
// It constructs the OAuth URL with the necessary scopes and returns it.
//
// Parameters:
//   - ctx: The Gin context for the current request.
//
// Returns:
//   - oauthURL: The constructed OAuth URL for GitHub authorization.
//   - err: An error if the redirection URL could not be constructed.
func (controller *githubController) RedirectToService(
	ctx *gin.Context,
) (oauthURL string, err error) {
	oauthURL, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Github,
		"https://github.com/login/oauth/authorize",
		"repo user user:email",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthURL, nil
}

// HandleServiceCallback handles the callback from the GitHub service.
// It binds the incoming request's credentials and processes the authentication code.
// If the code is missing or invalid, it returns an error.
// It also retrieves the Authorization header from the request and uses it to handle the service callback.
// If successful, it returns a bearer token; otherwise, it returns an error.
//
// Parameters:
//
//	ctx - The Gin context containing the request data.
//
// Returns:
//
//	A string containing the bearer token if successful, or an error if the callback handling fails.
func (controller *githubController) HandleServiceCallback(
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
		schemas.Github,
		controller.service.AuthGetServiceAccessToken,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	if err != nil {
		println(err.Error())
		return "", fmt.Errorf("unable to handle service callback because %w", err)
	}
	return bearer, nil
}

// HandleServiceCallbackMobile handles the callback from the mobile service for GitHub authentication.
// It binds the incoming request to the MobileTokenRequest schema and retrieves the Authorization header.
// Then, it calls the HandleServiceCallbackMobile method of the serviceService to process the callback.
//
// Parameters:
// - ctx: The Gin context for the incoming request.
//
// Returns:
// - A string representing the bearer token if successful.
// - An error if the binding of credentials or the service callback handling fails.
func (controller *githubController) HandleServiceCallbackMobile(
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
		schemas.Github,
		credentials,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	return bearer, err
}

// GetUserInfo retrieves user information from GitHub based on the provided authorization token.
// It extracts the token from the "Authorization" header, fetches user information using the token,
// and then retrieves additional user details from GitHub.
//
// Parameters:
//
//	ctx *gin.Context - The context of the HTTP request.
//
// Returns:
//
//	userInfo schemas.UserCredentials - The user's credentials including email and username.
//	err error - An error object if any error occurs during the process.
func (controller *githubController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.UserCredentials, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.serviceToken.GetTokenById(user.TokenId)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get token because %w", err)
	}

	githubUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = githubUserInfo.Email
	userInfo.Username = githubUserInfo.Username

	return userInfo, nil
}
