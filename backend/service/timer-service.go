package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Epitouche/Perimeter/repository"
	"github.com/Epitouche/Perimeter/schemas"
)

// Constructor

type TimerService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	GetActionsName() []string
	GetReactionsName() []string
	// Service specific functions
	// Actions functions
	TimerActionSpecificHour(c chan string, option json.RawMessage, idArea uint64)
	// Reactions functions
	TimerReactionGiveTime(option json.RawMessage, idArea uint64) string
}

type timerService struct {
	repository        repository.TimerRepository
	serviceRepository repository.ServiceRepository
	actionsName       []string
	reactionsName     []string
	serviceInfo       schemas.Service
}

func NewTimerService(
	repository repository.TimerRepository,
	serviceRepository repository.ServiceRepository,
) TimerService {
	return &timerService{
		repository:        repository,
		serviceRepository: serviceRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Timer,
			Description: "This service is a time service",
			Oauth:       false,
			Color:       "#BB00FF",
			Icon:        "https://api.iconify.design/mdi:clock.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *timerService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *timerService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	case string(schemas.SpecificTime):
		return service.TimerActionSpecificHour
	default:
		return nil
	}
}

func (service *timerService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	case string(schemas.GiveTime):
		return service.TimerReactionGiveTime
	default:
		return nil
	}
}

func (service *timerService) GetServiceActionInfo() []schemas.Action {
	service.actionsName = append(service.actionsName, string(schemas.SpecificTime))
	defaultValue := schemas.TimerActionSpecificHour{
		Hour:   0,
		Minute: 0,
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Timer,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:        string(schemas.SpecificTime),
			Description: "This action is a specific time action",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

func (service *timerService) GetServiceReactionInfo() []schemas.Reaction {
	service.reactionsName = append(service.reactionsName, string(schemas.GiveTime))
	defaultValue := struct{}{}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Timer,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.GiveTime),
			Description: "This reaction is a give time reaction",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

func (service *timerService) GetActionsName() []string {
	return service.actionsName
}

func (service *timerService) GetReactionsName() []string {
	return service.reactionsName
}

// Service specific functions

func getActualTime() (schemas.TimeApiResponse, error) {
	apiURL := "https://www.timeapi.io/api/time/current/zone?timeZone=Europe/Paris"

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return schemas.TimeApiResponse{}, schemas.ErrCreateRequest
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.TimeApiResponse{}, schemas.ErrDoRequest
	}

	if resp.StatusCode != http.StatusOK {
		return schemas.TimeApiResponse{}, fmt.Errorf("error status code %d", resp.StatusCode)
	}

	var result schemas.TimeApiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.TimeApiResponse{}, schemas.ErrDecode
	}

	resp.Body.Close()
	return result, nil
}

// Actions functions

func (service *timerService) TimerActionSpecificHour(
	c chan string,
	option json.RawMessage,
	idArea uint64,
) {
	optionJSON := schemas.TimerActionSpecificHour{}

	err := json.Unmarshal(option, &optionJSON)
	if err != nil {
		println("error unmarshal timer option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	actualTimeApi, err := getActualTime()
	if err != nil {
		println("error get actual time" + err.Error())
	} else {
		if actualTimeApi.Hour == optionJSON.Hour && actualTimeApi.Minute == optionJSON.Minute {
			response := "current time is " + actualTimeApi.Time
			println(response)
			c <- response
		}
	}
	time.Sleep(time.Minute)
}

// Reactions functions

func (service *timerService) TimerReactionGiveTime(
	option json.RawMessage,
	idArea uint64,
) string {
	actualTimeApi, err := getActualTime()
	if err != nil {
		println("error get actual time" + err.Error())
		return "error get actual time"
	} else {
		response := "current time is " + actualTimeApi.Time
		println(response)
		return response
	}
}
