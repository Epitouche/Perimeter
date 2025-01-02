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
	GetUserAllInfo(ctx *gin.Context) (userInfo schemas.UserAllInfo, err error)
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

	token, _, err := controller.userService.Login(newUser)
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

	if len(credentials.Username) < schemas.UsernameMinimumLength {
		return "", schemas.ErrUsernameTooShort
	}

	if len(credentials.Password) < schemas.PasswordMinimumLength {
		return "", schemas.ErrPasswordTooShort
	}

	if len(credentials.Email) < schemas.EmailMinimumLength {
		return "", schemas.ErrEmailTooShort
	}

	newUser := schemas.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	token, _, err := controller.userService.Register(newUser)
	if err != nil {
		return "", fmt.Errorf("can't register user: %w", err)
	}
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

	userInfo.Email = user.Email
	userInfo.Username = user.Username
	return userInfo, nil
}

func (controller *userController) GetUserAllInfo(
	ctx *gin.Context,
) (userInfo schemas.UserAllInfo, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.userService.GetUserInfo(tokenString)
	if err != nil {
		return userInfo, fmt.Errorf("unable to get user info because %w", err)
	}

	tokens, err := controller.tokenService.GetTokenByUserId(user.Id)
	if err != nil {
		return userInfo, fmt.Errorf("unable to get tokens info because %w", err)
	}

	userInfo.User = user
	userInfo.Tokens = tokens
	return userInfo, nil
}
