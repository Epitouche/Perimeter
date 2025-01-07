package controller

import (
	"github.com/Epitouche/Perimeter/service"
)

type TokenController interface{}

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
