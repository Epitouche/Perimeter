package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type DropboxController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
	GetUserFile(ctx *gin.Context) (userFile []schemas.DropboxFile, err error)
}

type dropboxController struct {
	service        service.DropboxService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

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

func (controller *dropboxController) RedirectToService(
	ctx *gin.Context,
) (oauthURL string, err error) {
	oauthURL, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Dropbox,
		"https://www.dropbox.com/oauth2/authorize",
		"account_info.read profile email openid",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthURL, nil
}

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

func (controller *dropboxController) HandleServiceCallbackMobile(
	ctx *gin.Context,
) (string, error) {
	var credentials schemas.MobileTokenRequest
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	bearer, err := controller.serviceService.HandleServiceCallbackMobile(
		schemas.Dropbox,
		credentials,
		controller.serviceUser,
		controller.service.GetUserInfo,
		controller.serviceToken,
	)
	return bearer, err
}

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

func (controller *dropboxController) GetUserFile(
	ctx *gin.Context,
) (userFile []schemas.DropboxFile, err error) {
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

	dropboxFile, err := controller.service.GetUserFileList(DropboxToken.Token)
	if err != nil {
		return userFile, fmt.Errorf("unable to get user info because %w", err)
	}

	userFile = dropboxFile
	return userFile, nil
}
