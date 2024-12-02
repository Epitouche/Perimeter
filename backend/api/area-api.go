package api

import (
	"area/controller"
)

type Areapi struct {
	controller controller.AreaController
}

func NewAreapi(controller controller.AreaController) *Areapi {
	return &Areapi{
		controller: controller,
	}
}
