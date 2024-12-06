package controller

import (
	"area/schemas"
	"area/service"
)

type ReactionController interface {
	GetReactionsInfo(id uint64) (response []schemas.Reaction, err error)
}

type reactionController struct {
	service service.ReactionService
}

func NewReactionController(service service.ReactionService) ReactionController {
	return &reactionController{
		service: service,
	}
}

func (controller *reactionController) GetReactionsInfo(
	id uint64,
) (response []schemas.Reaction, err error) {
	return controller.service.GetReactionsInfo(id)
}
