package controller

import (
	"area/service"
)

type ReactionController interface {
}

type reactionController struct {
	service service.ReactionService
}

func NewReactionController(service service.ReactionService) ReactionController {
	return &reactionController{
		service: service,
	}
}
