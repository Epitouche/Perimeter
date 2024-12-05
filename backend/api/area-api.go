package api

import (
	"area/controller"
)

type AreaApi struct {
	controller controller.AreaController
}

func NewAreapi(controller controller.AreaController) *AreaApi {
	return &AreaApi{
		controller: controller,
	}
}
