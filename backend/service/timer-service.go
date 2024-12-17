package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"area/repository"
	"area/schemas"
)

type TimerService interface {
	TimerActionSpecificHour(c chan string, option string, idArea uint64)
	TimerReactionGiveTime(option string, idArea uint64)
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option string, idArea uint64)
	FindReactionbyName(name string) func(option string, idArea uint64)
	GetActionsName() []string
	GetReactionsName() []string
}

type timerService struct {
	repository        repository.TimerRepository
	serviceRepository repository.ServiceRepository
	actionsName       []string
	reactionsName     []string
}

func NewTimerService(
	repository repository.TimerRepository,
	serviceRepository repository.ServiceRepository,
) TimerService {
	return &timerService{
		repository:        repository,
		serviceRepository: serviceRepository,
	}
}

func (service *timerService) FindActionbyName(
	name string,
) func(c chan string, option string, idArea uint64) {
	switch name {
	case string(schemas.SpecificTime):
		return service.TimerActionSpecificHour
	default:
		return nil
	}
}

func (service *timerService) FindReactionbyName(name string) func(option string, idArea uint64) {
	switch name {
	case string(schemas.GiveTime):
		return service.TimerReactionGiveTime
	default:
		return nil
	}
}

func getActualTime() (schemas.TimeApiResponse, error) {
	apiURL := "https://www.timeapi.io/api/time/current/zone?timeZone=Europe/Paris"

	req, err := http.NewRequest("GET", apiURL, nil)
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

func (service *timerService) TimerActionSpecificHour(c chan string, option string, idArea uint64) {
	optionJSON := schemas.TimerActionSpecificHour{}

	err := json.Unmarshal([]byte(option), &optionJSON)
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

func (service *timerService) TimerReactionGiveTime(option string, idArea uint64) {
	actualTimeApi, err := getActualTime()
	if err != nil {
		println("error get actual time" + err.Error())
	} else {
		response := "current time is " + actualTimeApi.Time
		println(response)
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
	return []schemas.Action{
		{
			Name:        string(schemas.SpecificTime),
			Description: "This action is a specific time action",
			Service:     service.serviceRepository.FindByName(schemas.Timer),
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
	return []schemas.Reaction{
		{
			Name:        string(schemas.GiveTime),
			Description: "This reaction is a give time reaction",
			Service:     service.serviceRepository.FindByName(schemas.Timer),
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
