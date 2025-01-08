package service

import (
	"github.com/Epitouche/Perimeter/repository"
	"github.com/Epitouche/Perimeter/schemas"
)

type ReactionService interface {
	FindAll() (reactions []schemas.Reaction, err error)
	SaveAllReaction()
	FindById(reactionId uint64) (reaction schemas.Reaction, err error)
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

func (service *reactionService) FindAll() (reactions []schemas.Reaction, err error) {
	reactions, err = service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return reactions, nil
}

func (service *reactionService) GetAllServicesByServiceId(
	serviceId uint64,
) (reactionJSON []schemas.ReactionJSON) {
	allRectionForService, err := service.repository.FindByServiceId(serviceId)
	if err != nil {
		println("Error when get all reactions by service id")
	}
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
				reactionByName, err := service.repository.FindByName(reaction.Name)
				if err != nil {
					println("Error when find reaction by name")
				}
				if len(reactionByName) == 0 {
					service.repository.Save(reaction)
				}
			}
		} else {
			println("ServiceReaction interface not implemented")
		}
	}
}

func (service *reactionService) FindById(reactionId uint64) (reaction schemas.Reaction, err error) {
	reaction, err = service.repository.FindById(reactionId)
	if err != nil {
		return reaction, err
	}
	return reaction, nil
}

func (service *reactionService) GetReactionsInfo(
	id uint64,
) (response []schemas.Reaction, err error) {
	response, err = service.repository.FindByServiceId(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
