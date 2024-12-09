package api

import (
	"area/controller"
)

type TokenApi struct {
	controller controller.TokenController
}

func NewTokenApi(controller controller.TokenController) *TokenApi {
	return &TokenApi{
		controller: controller,
	}
}
