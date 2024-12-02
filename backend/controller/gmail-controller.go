package controller

import (
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
	GetUserInfo(ctx *gin.Context) (userInfo schemas.GmailUserInfo, err error)
}

type gmailController struct {
	service     service.GmailService
	serviceUser service.UserService
}

func NewGmailController(
	service service.GmailService,
	serviceUser service.UserService,
) GmailController {
	return &gmailController{
		service:     service,
		serviceUser: serviceUser,
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
		"&scope=https://mail.google.com/" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}

func (controller *gmailController) HandleServiceCallback(
	ctx *gin.Context,
	path string,
) (string, error) {
	code := ctx.Query("code")
	if code == "" {
		return "", fmt.Errorf("missing code")
	}

	state := ctx.Query("state")
	latestCSRFToken, err := ctx.Cookie("latestCSRFToken")
	if err != nil {
		return "", fmt.Errorf("missing CSRF token")
	}

	if state != latestCSRFToken {
		return "", fmt.Errorf("invalid CSRF token")
	}

	githubTokenResponse, err := controller.service.AuthGetServiceAccessToken(code, path)
	if err != nil {
		return "", fmt.Errorf("unable to get access token because %w", err)
	}

	newGmailToken := schemas.GmailToken{
		AccessToken: githubTokenResponse.AccessToken,
		Scope:       githubTokenResponse.Scope,
		TokenType:   githubTokenResponse.TokenType,
	}

	// Save the access token in the database
	tokenId, err := controller.service.SaveToken(newGmailToken)
	userAlreadExists := false
	if err != nil {
		if err.Error() == "token already exists" {
			userAlreadExists = true
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	userInfo, err := controller.service.GetUserInfo(newGmailToken.AccessToken)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}

	newUser := schemas.User{
		Username: userInfo.Login,
		Email:    userInfo.Email,
		GithubId: tokenId,
	}

	if userAlreadExists {
		token, err := controller.serviceUser.Login(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to login user because %w", err)
		}
		return token, nil
	} else {
		token, err := controller.serviceUser.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		return token, nil
	}
}

func (controller *gmailController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.GmailUserInfo, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.service.GetTokenById(user.GithubId)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to get token because %w", err)
	}

	githubUserInfo, err := controller.service.GetUserInfo(token.AccessToken)
	if err != nil {
		return schemas.GmailUserInfo{}, fmt.Errorf("unable to get user info because %w", err)
	}

	return githubUserInfo, nil
}
