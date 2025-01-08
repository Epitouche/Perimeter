package api

import (
	"github.com/Epitouche/Perimeter/controller"
)

type TokenApi struct {
	controller controller.TokenController
}

func NewTokenApi(controller controller.TokenController) *TokenApi {
	return &TokenApi{
		controller: controller,
	}
}
