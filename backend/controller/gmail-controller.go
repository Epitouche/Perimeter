package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type GmailController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context, path string) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
}

type gmailController struct {
	service        service.GmailService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

func NewGmailController(
	service service.GmailService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
	serviceService service.ServiceService,
) GmailController {
	return &gmailController{
		service:        service,
		serviceUser:    serviceUser,
		serviceToken:   serviceToken,
		serviceService: serviceService,
	}
}

func (controller *gmailController) RedirectToService(
	ctx *gin.Context,
) (oauthURL string, err error) {
	oauthURL, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Gmail,
		"https://accounts.google.com/o/oauth2/v2/auth",
		"https://mail.google.com/ profile email",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthURL, nil
}

func (controller *gmailController) HandleServiceCallback(
	ctx *gin.Context,
	path string,
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

	// state := credentials.State
	// latestCSRFToken, err := ctx.Cookie("latestCSRFToken")
	// if err != nil {
	// 	return "", fmt.Errorf("missing CSRF token")
	// }

	// if state != latestCSRFToken {
	// 	return "", fmt.Errorf("invalid CSRF token")
	// }

	authHeader := ctx.GetHeader("Authorization")

	bearer, err := controller.serviceService.HandleServiceCallback(
		code,
		authHeader,
		schemas.Gmail,
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

func (controller *gmailController) HandleServiceCallbackMobile(
	ctx *gin.Context,
) (string, error) {
	var credentials schemas.MobileTokenRequest
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	bearer, err := controller.serviceService.HandleServiceCallbackMobile(
		schemas.Gmail,
		credentials,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	return bearer, err
}

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
