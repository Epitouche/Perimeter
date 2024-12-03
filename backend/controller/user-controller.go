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
}

type userController struct {
	userService service.UserService
	jWtService  service.JWTService
}

func NewUserController(userService service.UserService,
	jWtService service.JWTService,
) UserController {
	return &userController{
		userService: userService,
		jWtService:  jWtService,
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
	token, err := controller.userService.Register(newUser)
	if err != nil {
		return "", fmt.Errorf("can't register user: %w", err)
	}
	return token, nil
}
