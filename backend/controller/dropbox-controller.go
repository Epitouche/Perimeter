package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type DropboxController interface {
	RedirectToService(ctx *gin.Context) (oauthURL string, err error)
	HandleServiceCallback(ctx *gin.Context, path string) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context, path string) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
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
	path string,
) (string, error) {
	var credentials schemas.TokenCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	token := credentials.Token
	if token == "" {
		return "", schemas.ErrMissingAuthenticationCode
	}

	email := credentials.Email
	if email == "" {
		return "", schemas.ErrMissingAuthenticationCode
	}

	username := credentials.Username
	if username == "" {
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
	newUser := schemas.User{
		Username: username,
		Email:    email,
	}
	dropboxToken := schemas.Token{}
	dropboxToken.Token = token
	var bearerToken string

	if len(authHeader) > len("Bearer ") {
		bearerToken = authHeader[len("Bearer "):]
	} else {

		bearerTokenLogin, _, err := controller.serviceUser.Login(newUser)
		if err == nil {
			return bearerTokenLogin, nil
		}

		bearerTokenRegister, newUserId, err := controller.serviceUser.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		bearerToken = bearerTokenRegister
		newUser = controller.serviceUser.GetUserById(newUserId)
	}

	dropboxService := controller.serviceService.FindByName(schemas.Dropbox)

	newDropboxToken := schemas.Token{
		Token:        dropboxToken.Token,
		RefreshToken: dropboxToken.RefreshToken,
		ExpireAt:     dropboxToken.ExpireAt,
		Service:      dropboxService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newDropboxToken)
	if err != nil {
		if errors.Is(err, schemas.ErrTokenAlreadyExists) {
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	if len(authHeader) == 0 {
		newUser.TokenId = tokenId

		err = controller.serviceUser.UpdateUserInfo(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to update user info because %w", err)
		}
	}
	return bearerToken, nil
}

func (controller *dropboxController) GetUserInfo(
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

	dropboxUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = dropboxUserInfo.Email
	userInfo.Username = dropboxUserInfo.Username
	return userInfo, nil
}
