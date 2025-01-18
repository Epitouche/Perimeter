package controller

import (
	"area/service"
)

type OpenWeatherMapController interface{}

// openweathermapController is a controller that handles requests related to OpenWeatherMap data.
// It uses the OpenWeatherMapService to interact with the OpenWeatherMap API and retrieve weather information.
type openweathermapController struct {
	service service.OpenWeatherMapService
}

// NewOpenWeatherMapController creates a new instance of OpenWeatherMapController
// with the provided OpenWeatherMapService.
//
// Parameters:
//   - service: an instance of OpenWeatherMapService that provides weather data.
//
// Returns:
//   - An instance of OpenWeatherMapController.
func NewOpenWeatherMapController(service service.OpenWeatherMapService) OpenWeatherMapController {
	return &openweathermapController{
		service: service,
	}
}
