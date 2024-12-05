package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
	GetUserInfo(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
}

type userController struct {
	userService  service.UserService
	jWtService   service.JWTService
	tokenService service.TokenService
}

func NewUserController(userService service.UserService,
	jWtService service.JWTService, tokenService service.TokenService,
) UserController {
	return &userController{
		userService:  userService,
		jWtService:   jWtService,
		tokenService: tokenService,
	}
}

func (controller *userController) Login(ctx *gin.Context) (string, error) {
	var credentials schemas.LoginCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}

	newUser := schemas.User{
		Username: credentials.Username,
		Password: credentials.Password,
	}

	token, err := controller.userService.Login(newUser)
	if err != nil {
		return "", fmt.Errorf("can't login user: %w", err)
	}
	return token, nil
}

func (controller *userController) Register(ctx *gin.Context) (string, error) {
	var credentials schemas.RegisterCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", fmt.Errorf("can't bind credentials: %w", err)
	}
	if len(credentials.Username) < 4 {
		return "", fmt.Errorf("username must be at least 4 characters long" + credentials.Username)
	}
	if len(credentials.Password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters long")
	}
	if len(credentials.Email) < 4 {
		return "", fmt.Errorf("email must be at least 4 characters long")
	}

	newUser := schemas.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: credentials.Password,
	}
	token, newUserId, err := controller.userService.Register(newUser)
	if err != nil {
		return "", fmt.Errorf("can't register user: %w", err)
	}
	print(newUserId)
	return token, nil
}

func (controller *userController) GetUserInfo(
	ctx *gin.Context,
) (userInfo schemas.UserCredentials, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.userService.GetUserInfo(tokenString)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	token, err := controller.tokenService.GetTokenById(user.Id)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get token because %w", err)
	}

	spotifyUserInfo, err := controller.userService.GetUserInfo(token.Token)
	if err != nil {
		return schemas.UserCredentials{}, fmt.Errorf("unable to get user info because %w", err)
	}

	userInfo.Email = spotifyUserInfo.Email
	userInfo.Username = spotifyUserInfo.Username
	return userInfo, nil
}
