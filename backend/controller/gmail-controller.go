package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
	"area/tools"
)

type GmailController interface {
	RedirectToService(ctx *gin.Context, path string) (string, error)
	HandleServiceCallback(ctx *gin.Context, path string) (string, error)
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
	path string,
) (string, error) {
	clientID := os.Getenv("GMAIL_CLIENT_ID")
	if clientID == "" {
		return "", schemas.ErrGmailClientIdNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return "", schemas.ErrGmailSecretNotSet
	}

	// Generate the CSRF token
	state, err := tools.GenerateCSRFToken()
	if err != nil {
		return "", fmt.Errorf("unable to generate CSRF token because %w", err)
	}

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the GitHub authorization URL
	redirectURI := "http://localhost:8081/services/gmail"
	authURL := "https://accounts.google.com/o/oauth2/v2/auth" +
		"?client_id=" + clientID +
		"&response_type=code" +
		"&scope=https://mail.google.com/ profile email" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
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
	newUser := schemas.User{}
	gmailToken := schemas.Token{}
	bearerToken := ""

	gmailTokenResponse, err := controller.service.AuthGetServiceAccessToken(code, path)
	if err != nil {
		return "", fmt.Errorf("unable to get access token because %w", err)
	}
	gmailToken.Token = gmailTokenResponse.AccessToken
	gmailToken.RefreshToken = gmailTokenResponse.RefreshToken

	if len(authHeader) > len("Bearer ") {
		bearerToken = authHeader[len("Bearer "):]

		newUser, err = controller.serviceUser.GetUserInfo(bearerToken)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
	} else {

		userInfo, err := controller.service.GetUserInfo(gmailToken.Token)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
		newUser = schemas.User{
			Username: userInfo.Login,
			Email:    userInfo.Email,
		}

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

	newgmailToken := schemas.Token{
		Token:        gmailToken.Token,
		RefreshToken: gmailToken.RefreshToken,
		Service:      gmailService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newgmailToken)
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
	userInfo.Username = gmailUserInfo.Login
	return userInfo, nil
}
