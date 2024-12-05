package controller

import (
	"fmt"
	"os"
	"time"

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
		return "", fmt.Errorf("GMAIL_CLIENT_ID is not set")
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return "", fmt.Errorf("BACKEND_PORT is not set")
	}

	// Generate the CSRF token
	state, err := tools.GenerateCSRFToken()
	if err != nil {
		return "", fmt.Errorf("unable to generate CSRF token")
	}

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the GitHub authorization URL
	redirectURI := "http://localhost:" + appPort + path
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
		return "", fmt.Errorf("missing code")
	}

	// state := credentials.State
	// latestCSRFToken, err := ctx.Cookie("latestCSRFToken")
	// if err != nil {
	// 	return "", fmt.Errorf("missing CSRF token")
	// }

	// if state != latestCSRFToken {
	// 	return "", fmt.Errorf("invalid CSRF token")
	// }

	gmailTokenResponse, err := controller.service.AuthGetServiceAccessToken(code, path)
	if err != nil {
		return "", fmt.Errorf("unable to get access token because %w", err)
	}

	userInfo, err := controller.service.GetUserInfo(gmailTokenResponse.AccessToken)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}

	newUser := schemas.User{
		Username: userInfo.Login,
		Email:    userInfo.Email,
	}

	token, err := controller.serviceUser.Login(newUser)
	if err == nil {
		return token, nil
	}

	token, newUserId, err := controller.serviceUser.Register(newUser)
	if err != nil {
		return "", fmt.Errorf("unable to register user because %w", err)
	}

	gmailService := controller.serviceService.FindByName(schemas.OpenWeatherMap)
	savedUser := controller.serviceUser.GetUserById(newUserId)

	newSpotifyToken := schemas.Token{
		Token:        gmailTokenResponse.AccessToken,
		RefreshToken: gmailTokenResponse.RefreshToken,
		ExpireAt:     time.Now().Add(time.Duration(gmailTokenResponse.ExpiresIn) * time.Second),
		Service:      gmailService,
		User:         savedUser,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newSpotifyToken)
	if err != nil {
		if err.Error() == "token already exists" {
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	savedUser.TokenId = tokenId

	controller.serviceUser.UpdateUserInfo(savedUser)
	return token, nil
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
