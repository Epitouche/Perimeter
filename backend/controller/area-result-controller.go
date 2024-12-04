package controller

import (
	"area/service"
)

type AreaResultController interface{}

type areaResultController struct {
	service service.AreaResultService
}

func NewAreaResultController(service service.AreaResultService) AreaController {
	return &areaResultController{
		service: service,
	}
}
