package api

import (
	"area/controller"
)

type TimerAPI struct {
	controller controller.TimerController
}

func NewTimerAPI(controller controller.TimerController) *TimerAPI {
	return &TimerAPI{
		controller: controller,
	}
}
