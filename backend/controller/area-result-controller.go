package controller

import (
	"area/service"
)

type AreaResultController interface{}

type areaResultController struct {
	service service.AreaResultService
}

func NewAreaResultController(service service.AreaResultService) AreaResultController {
	return &areaResultController{
		service: service,
	}
}
