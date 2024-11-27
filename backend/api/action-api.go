package api

import (
	"area/controller"
)

type ActionApi struct {
	actionController controller.ActionController
}

func NewActionApi(actionController controller.ActionController) *ActionApi {
	return &ActionApi{
		actionController: actionController,
	}
}
