package service

import (
	"area/repository"
	"area/schemas"
)

type ReactionService interface {
	Save(newService schemas.Reaction) error
	Update(newService schemas.Reaction) error
	Delete(newService schemas.Reaction) error
	FindAll() []schemas.Reaction
}

type reactionService struct {
	repository     repository.ReactionRepository
	serviceService ServiceService
	allAction      map[schemas.ServiceName][]schemas.Reaction
}

func NewReactionService(repository repository.ReactionRepository, serviceService ServiceService) ReactionService {
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

func (service *reactionService) Save(newService schemas.Reaction) error {
	service.repository.Save(newService)
	return nil
}

func (service *reactionService) Update(newService schemas.Reaction) error {
	service.repository.Update(newService)
	return nil
}

func (service *reactionService) Delete(newService schemas.Reaction) error {
	service.repository.Delete(newService)
	return nil
}

func (service *reactionService) FindAll() []schemas.Reaction {
	return service.repository.FindAll()
}
