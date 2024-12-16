package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type GmailController interface {
	RedirectToService(ctx *gin.Context) (string, error)
	HandleServiceCallback(ctx *gin.Context, path string) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context, path string) (string, error)
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
) (oauthUrl string, err error) {
	oauthUrl, err = controller.serviceService.RedirectToServiceOauthPage(
		schemas.Gmail,
		"https://accounts.google.com/o/oauth2/v2/auth",
		"https://mail.google.com/ profile email",
	)
	if err != nil {
		return "", fmt.Errorf("unable to redirect to service oauth page because %w", err)
	}
	return oauthUrl, nil
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
	gmailToken := schemas.Token{}
	gmailToken.Token = token
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

	gmailService := controller.serviceService.FindByName(schemas.Gmail)

	newGmailToken := schemas.Token{
		Token:        gmailToken.Token,
		RefreshToken: gmailToken.RefreshToken,
		ExpireAt:     gmailToken.ExpireAt,
		Service:      gmailService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newGmailToken)
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
