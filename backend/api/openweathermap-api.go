package api

import (
	"github.com/Epitouche/Perimeter/controller"
)

type OpenweathermapAPI struct {
	controller controller.OpenweathermapController
}

func OewopenweathermapAPI(controller controller.OpenweathermapController) *OpenweathermapAPI {
	return &OpenweathermapAPI{
		controller: controller,
	}
}
