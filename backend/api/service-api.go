package api

import (
	"area/controller"
)

type ServiceApi struct {
	serviceController controller.ServiceController
}

func NewServiceAPI(serviceController controller.ServiceController) *ServiceApi {
	return &ServiceApi{
		serviceController: serviceController,
	}
}
