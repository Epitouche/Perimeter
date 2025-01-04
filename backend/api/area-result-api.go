package api

import (
	"area/controller"
)

type AreaResultApi struct {
	controller controller.AreaResultController
}

func NewAreaResultApi(controller controller.AreaResultController) *AreaResultApi {
	return &AreaResultApi{
		controller: controller,
	}
}
