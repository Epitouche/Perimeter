package service

import (
	"area/repository"
	"area/schemas"
)

type ReactionService interface {
	FindAll() []schemas.Reaction
	GetAllServicesByServiceId(serviceId uint64) (reactionJson []schemas.ReactionJson)
}

type reactionService struct {
	repository     repository.ReactionRepository
	serviceService ServiceService
	allAction      map[schemas.ServiceName][]schemas.Reaction
}

func NewReactionService(
	repository repository.ReactionRepository,
	serviceService ServiceService,
) ReactionService {
	newService := reactionService{
		repository:     repository,
		serviceService: serviceService,
		allAction: map[schemas.ServiceName][]schemas.Reaction{
			schemas.Spotify: {
				{
					Name:        "action1",
					Description: "do something",
				},
				{
					Name:        "action2",
					Description: "do something",
				},
				{
					Name:        "action3",
					Description: "do something",
				},
			},
			schemas.OpenWeatherMap: {
				{
					Name:        "action1",
					Description: "do something",
				},
				{
					Name:        "action2",
					Description: "do something",
				},
				{
					Name:        "action3",
					Description: "do something",
				},
			},
			schemas.Timer: {
				{
					Name:        "action1",
					Description: "do something",
				},
				{
					Name:        "action2",
					Description: "do something",
				},
				{
					Name:        "action3",
					Description: "do something",
				},
			},
		},
	}
	newService.InitialSaveAction()
	return &newService
}

func (service *reactionService) InitialSaveAction() {
	allService := service.serviceService.FindAll()
	// Find all service and save action
	for _, oneService := range allService {
		// Find all action by service name
		for _, oneAction := range service.allAction[schemas.ServiceName(oneService.Name)] {
			existingActions := service.repository.FindByServiceByName(oneService.Id, oneAction.Name)

			if len(existingActions) == 0 {
				oneAction.Service = oneService
				service.repository.Save(oneAction)
			}
		}
	}
}

func (service *reactionService) FindAll() []schemas.Reaction {
	return service.repository.FindAll()
}

func (service *reactionService) GetAllServicesByServiceId(
	serviceId uint64,
) (reactionJson []schemas.ReactionJson) {
	allRectionForService := service.repository.FindByServiceId(serviceId)
	for _, oneReaction := range allRectionForService {
		reactionJson = append(reactionJson, schemas.ReactionJson{
			Name:        oneReaction.Name,
			Description: oneReaction.Description,
		})
	}
	return reactionJson
}
