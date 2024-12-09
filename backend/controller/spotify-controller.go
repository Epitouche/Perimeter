package controller

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
	"area/tools"
)

type SpotifyController interface {
	RedirectToService(ctx *gin.Context) (string, error)
	HandleServiceCallback(ctx *gin.Context) (string, error)
	HandleServiceCallbackMobile(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
}

type spotifyController struct {
	service        service.SpotifyService
	serviceUser    service.UserService
	serviceToken   service.TokenService
	serviceService service.ServiceService
}

func NewSpotifyController(
	service service.SpotifyService,
	serviceUser service.UserService,
	serviceToken service.TokenService,
	serviceService service.ServiceService,
) SpotifyController {
	return &spotifyController{
		service:        service,
		serviceUser:    serviceUser,
		serviceToken:   serviceToken,
		serviceService: serviceService,
	}
}

func (controller *spotifyController) RedirectToService(
	ctx *gin.Context,
) (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientID == "" {
		return "", schemas.ErrSpotifyClientIdNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return "", schemas.ErrBackendPortNotSet
	}

	// Generate the CSRF token
	state, err := tools.GenerateCSRFToken()
	if err != nil {
		return "", fmt.Errorf("unable to generate CSRF token because %w", err)
	}

	// Store the CSRF token in session (you can replace this with a session library or in-memory storage)
	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	// Construct the Spotify authorization URL
	redirectURI := "http://localhost:8081/services/spotify"
	authURL := "https://accounts.spotify.com/authorize" +
		"?response_type=code" +
		"&client_id=" + clientID +
		"&scope=user-read-private user-read-email user-modify-playback-state" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}

func (controller *spotifyController) HandleServiceCallback(
	ctx *gin.Context,
) (string, error) {
	var credentials schemas.CodeCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	code := credentials.Code
	if code == "" {
		return "", schemas.ErrMissingCode
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
	spotifyToken := schemas.Token{}
	var bearerToken string

	spotifyTokenResponse, err := controller.service.AuthGetServiceAccessToken(code)
	if err != nil {
		println(fmt.Errorf("unable to get access token because %w", err))
		return "", fmt.Errorf("unable to get access token because %w", err)
	}
	spotifyToken.Token = spotifyTokenResponse.AccessToken
	spotifyToken.RefreshToken = spotifyTokenResponse.RefreshToken
	spotifyToken.ExpireAt = time.Now().
		Add(time.Duration(spotifyTokenResponse.ExpiresIn) * time.Second)

	if len(authHeader) > len(schemas.BearerTokenType) {
		bearerToken = authHeader[len(schemas.BearerTokenType):]

		newUser, err = controller.serviceUser.GetUserInfo(bearerToken)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
	} else {
		userInfo, err := controller.service.GetUserInfo(spotifyToken.Token)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
		newUser = schemas.User{
			Username: userInfo.DisplayName,
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

	spotifyService := controller.serviceService.FindByName(schemas.Spotify)

	newSpotifyToken := schemas.Token{
		Token:        spotifyToken.Token,
		RefreshToken: spotifyToken.RefreshToken,
		ExpireAt:     spotifyToken.ExpireAt,
		Service:      spotifyService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newSpotifyToken)
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

func (controller *spotifyController) HandleServiceCallbackMobile(
	ctx *gin.Context,
) (string, error) {

	var credentials schemas.CodeCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	code := credentials.Code
	if code == "" {
		return "", schemas.ErrMissingCode
	}

	println("4444444444444")
	println("code verify", credentials.CodeVerifier)
	codeVerifier := credentials.CodeVerifier
	if codeVerifier == "" {
		return "", schemas.ErrMissingCodeVerifier
	}

	println("11111111111111111111")
	authHeader := ctx.GetHeader("Authorization")
	newUser := schemas.User{}
	spotifyToken := schemas.Token{}
	var bearerToken string

	println("222222222222222222222")
	spotifyTokenResponse, err := controller.service.AuthGetServiceAccessTokenMobile(
		code,
		codeVerifier,
	)
	if err != nil {
		println(fmt.Errorf("unable to get access token because %w", err))
		return "", fmt.Errorf("unable to get access token because %w", err)
	}
	spotifyToken.Token = spotifyTokenResponse.AccessToken
	spotifyToken.RefreshToken = spotifyTokenResponse.RefreshToken
	spotifyToken.ExpireAt = time.Now().
		Add(time.Duration(spotifyTokenResponse.ExpiresIn) * time.Second)

	println("3333333333333333333333")

	if len(authHeader) > len(schemas.BearerTokenType) {
		bearerToken = authHeader[len(schemas.BearerTokenType):]

		newUser, err = controller.serviceUser.GetUserInfo(bearerToken)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
	} else {
		userInfo, err := controller.service.GetUserInfo(spotifyToken.Token)
		if err != nil {
			return "", fmt.Errorf("unable to get user info because %w", err)
		}
		newUser = schemas.User{
			Username: userInfo.DisplayName,
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

	spotifyService := controller.serviceService.FindByName(schemas.Spotify)

	newSpotifyToken := schemas.Token{
		Token:        spotifyToken.Token,
		RefreshToken: spotifyToken.RefreshToken,
		ExpireAt:     spotifyToken.ExpireAt,
		Service:      spotifyService,
		User:         newUser,
	}

	// Save the access token in the database
	tokenId, err := controller.serviceToken.SaveToken(newSpotifyToken)
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

func (controller *spotifyController) GetUserInfo(
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

	spotifyUserInfo, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = spotifyUserInfo.Email
	userInfo.Username = spotifyUserInfo.DisplayName
	return userInfo, nil
}
