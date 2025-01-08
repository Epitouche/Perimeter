package controller

import (
	"fmt"

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
	response, err = controller.service.GetReactionsInfo(id)
	if err != nil {
		return nil, fmt.Errorf("can't get reactions info: %w", err)
	}
	return response, nil
}
