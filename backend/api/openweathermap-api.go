package api

import (
	"area/controller"
)

type OpenweathermapAPI struct {
	controller controller.OpenweathermapController
}

func OewopenweathermapAPI(controller controller.OpenweathermapController) *OpenweathermapAPI {
	return &OpenweathermapAPI{
		controller: controller,
	}
}
