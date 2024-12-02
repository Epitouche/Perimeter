package api

import (
	"area/controller"
)

type ReactionApi struct {
	controller controller.ReactionController
}

func NewReactionApi(controller controller.ReactionController) *ReactionApi {
	return &ReactionApi{
		controller: controller,
	}
}
