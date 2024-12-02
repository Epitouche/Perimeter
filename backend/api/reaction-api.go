package api

import (
	"area/controller"
)

type ReactionApi struct {
	reactionController controller.ReactionController
}

func NewReactionApi(reactionController controller.ReactionController) *ReactionApi {
	return &ReactionApi{
		reactionController: reactionController,
	}
}
