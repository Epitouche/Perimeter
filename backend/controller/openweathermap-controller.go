package controller

import (
	"area/service"
)

type OpenweathermapController interface{}

type openweathermapController struct {
	service service.OpenweathermapService
}

func NewOpenweathermapController(service service.OpenweathermapService) OpenweathermapController {
	return &openweathermapController{
		service: service,
	}
}
