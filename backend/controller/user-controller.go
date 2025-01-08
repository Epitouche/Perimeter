package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Epitouche/Perimeter/schemas"
	"github.com/Epitouche/Perimeter/service"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
	GetUser(ctx *gin.Context) (userInfo schemas.UserCredentials, err error)
	GetUserAllInfo(ctx *gin.Context) (userInfo schemas.UserAllInfo, err error)
	UpdateUser(
		ctx *gin.Context,
	) (updatedUser schemas.User, err error)
	DeleteUser(
		ctx *gin.Context,
	) (updatedUser schemas.User, err error)
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

func (controller *userController) GetUser(
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

func (controller *userController) UpdateUser(
	ctx *gin.Context,
) (updatedUser schemas.User, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	var result schemas.User

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return updatedUser, fmt.Errorf("can't bind credentials: %w", err)
	}

	user, err := controller.userService.GetUserInfo(tokenString)
	if err != nil {
		return updatedUser, fmt.Errorf("unable to get user info because %w", err)
	}

	if result.Id == user.Id {
		err = controller.userService.UpdateUserInfo(result)
		if err != nil {
			return updatedUser, fmt.Errorf("unable to update user info because %w", err)
		}
		return result, nil
	} else {
		return updatedUser, errors.New("unable to update user info because not the right user")
	}
}

func (controller *userController) DeleteUser(
	ctx *gin.Context,
) (updatedUser schemas.User, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	var result schemas.User

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return updatedUser, fmt.Errorf("can't bind credentials: %w", err)
	}

	user, err := controller.userService.GetUserInfo(tokenString)
	if err != nil {
		return updatedUser, fmt.Errorf("unable to get user info because %w", err)
	}

	if result.Id == user.Id {
		err = controller.userService.DeleteUser(result)
		if err != nil {
			return updatedUser, fmt.Errorf("unable to update user info because %w", err)
		}
		return result, nil
	} else {
		return updatedUser, errors.New("unable to update user info because not the right user")
	}
}
