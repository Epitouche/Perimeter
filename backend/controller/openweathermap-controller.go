package controller

import (
	"area/service"
)

type OpenWeatherMapController interface{}

type openweathermapController struct {
	service service.OpenWeatherMapService
}

func NewOpenWeatherMapController(service service.OpenWeatherMapService) OpenWeatherMapController {
	return &openweathermapController{
		service: service,
	}
}
