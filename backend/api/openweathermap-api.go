package api

import (
	"area/controller"
)

type OpenWeatherMapAPI struct {
	controller controller.OpenWeatherMapController
}

func OewopenweathermapAPI(controller controller.OpenWeatherMapController) *OpenWeatherMapAPI {
	return &OpenWeatherMapAPI{
		controller: controller,
	}
}
