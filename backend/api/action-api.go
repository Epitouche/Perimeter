package api

import (
	"area/controller"
)

type ActionApi struct {
	controller controller.ActionController
}

func NewActionApi(controller controller.ActionController) *ActionApi {
	return &ActionApi{
		controller: controller,
	}
}
