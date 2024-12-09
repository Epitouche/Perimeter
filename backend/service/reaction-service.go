package service

import (
	"area/repository"
	"area/schemas"
)

type ReactionService interface {
	FindAll() []schemas.Reaction
	SaveAllReaction()
	FindById(reactionId uint64) schemas.Reaction
	GetReactionsInfo(id uint64) (response []schemas.Reaction, err error)
	GetAllServicesByServiceId(serviceId uint64) (reactionJSON []schemas.ReactionJSON)
}

type ServiceReaction interface {
	GetServiceReactionInfo() []schemas.Reaction
}

type reactionService struct {
	repository     repository.ReactionRepository
	serviceService ServiceService
}

func NewReactionService(
	repository repository.ReactionRepository,
	serviceService ServiceService,
) ReactionService {
	newService := &reactionService{
		repository:     repository,
		serviceService: serviceService,
	}
	newService.SaveAllReaction()
	return newService
}

func (service *reactionService) FindAll() []schemas.Reaction {
	return service.repository.FindAll()
}

func (service *reactionService) GetAllServicesByServiceId(
	serviceId uint64,
) (reactionJSON []schemas.ReactionJSON) {
	allRectionForService := service.repository.FindByServiceId(serviceId)
	for _, oneReaction := range allRectionForService {
		reactionJSON = append(reactionJSON, schemas.ReactionJSON{
			Name:        oneReaction.Name,
			Description: oneReaction.Description,
		})
	}
	return reactionJSON
}

func (service *reactionService) SaveAllReaction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceReaction, ok := services.(ServiceReaction); ok {
			reactions := serviceReaction.GetServiceReactionInfo()
			for _, reaction := range reactions {
				reactionByName := service.repository.FindByName(reaction.Name)
				if len(reactionByName) == 0 {
					service.repository.Save(reaction)
				}
			}
		} else {
			println("ServiceReaction interface not implemented")
		}
	}
}

func (service *reactionService) FindById(reactionId uint64) schemas.Reaction {
	return service.repository.FindById(reactionId)
}

func (service *reactionService) GetReactionsInfo(
	id uint64,
) (response []schemas.Reaction, err error) {
	return service.repository.FindByServiceId(id), nil
}
