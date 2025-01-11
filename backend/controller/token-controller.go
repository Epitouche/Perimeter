package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area/schemas"
	"area/service"
)

type TokenController interface {
	DeleteUserToken(
		ctx *gin.Context,
	) (newToken schemas.Token, err error)
}

type tokenController struct {
	service service.TokenService
}

func NewTokenController(
	service service.TokenService,
) TokenController {
	return &tokenController{
		service: service,
	}
}

func (controller *tokenController) DeleteUserToken(
	ctx *gin.Context,
) (newToken schemas.Token, err error) {
	var result struct{ Id uint64 }

	err = json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		println(fmt.Errorf("can't bind credentials: %w", err))
		return newToken, fmt.Errorf("can't bind credentials: %w", err)
	}

	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]
	newToken, err = controller.service.DeleteUserToken(token, result)
	if err != nil {
		return newToken, fmt.Errorf("can't get user areas: %w", err)
	}
	return newToken, nil
}
