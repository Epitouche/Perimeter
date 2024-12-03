package controller

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
	"area/tools"
)

type SpotifyController interface {
	RedirectToService(ctx *gin.Context, path string) (string, error)
	HandleServiceCallback(ctx *gin.Context, path string) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.SpotifyUserInfo, err error)
}

type spotifyController struct {
	service      service.SpotifyService
	serviceUser  service.UserService
	serviceToken service.TokenService
}

func NewSpotifyController(
	service service.SpotifyService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
) SpotifyController {
	return &spotifyController{
		service:      service,
		serviceUser:  serviceUser,
		serviceToken: serviceToken,
	}
}

func (controller *spotifyController) RedirectToService(
	ctx *gin.Context,
	path string,
) (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientID == "" {
		return "", fmt.Errorf("SPOTIFY_CLIENT_ID is not set")
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

	// Construct the Spotify authorization URL
	redirectURI := "http://localhost:" + appPort + path
	authURL := "https://accounts.spotify.com/authorize" +
		"?response_type=code" +
		"&client_id=" + clientID +
		"&scope=user-read-private user-read-email" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}

func (controller *spotifyController) HandleServiceCallback(
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

	spotifyTokenResponse, err := controller.service.AuthGetServiceAccessToken(code, path)
	if err != nil {
		return "", fmt.Errorf("unable to get access token because %w", err)
	}

	newSpotifyToken := schemas.Token{
		Token:  spotifyTokenResponse.AccessToken,
		UserId: 1,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newSpotifyToken)
	userAlreadExists := false
	if err != nil {
		if err.Error() == "token already exists" {
			userAlreadExists = true
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	userInfo, err := controller.service.GetUserInfo(newSpotifyToken.Token)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}

	newUser := schemas.User{
		Username: userInfo.Login,
		Email:    userInfo.Email,
		TokenId:  tokenId,
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

func (controller *spotifyController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.SpotifyUserInfo, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.serviceUser.GetUserInfo(tokenString)
	if err != nil {
		return schemas.SpotifyUserInfo{}, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.serviceToken.GetTokenById(user.Id)
	if err != nil {
		return schemas.SpotifyUserInfo{}, fmt.Errorf("unable to get token because %w", err)
	}

	spotifyUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.SpotifyUserInfo{}, fmt.Errorf("unable to get user info because %w", err)
	}

	return spotifyUserInfo, nil
}
