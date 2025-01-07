package controller

import (
	"github.com/Epitouche/Perimeter/service"
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
