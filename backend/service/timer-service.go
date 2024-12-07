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

func getActualTime() (schemas.TimeAPISTRUCT, error) {
	apiURL := "https://www.timeapi.io/api/time/current/zone?timeZone=Europe/Paris"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return schemas.TimeAPISTRUCT{}, fmt.Errorf("error create request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.TimeAPISTRUCT{}, fmt.Errorf("error do request")
	}

	if resp.StatusCode != 200 {
		return schemas.TimeAPISTRUCT{}, fmt.Errorf("error status code %d", resp.StatusCode)
	}

	var result schemas.TimeAPISTRUCT
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.TimeAPISTRUCT{}, fmt.Errorf("error decode")
	}

	resp.Body.Close()
	return result, nil
}

func (service *timerService) TimerActionSpecificHour(c chan string, option string, idArea uint64) {
	println("option" + option)

	optionJson := schemas.TimerActionSpecificHour{}

	err := json.Unmarshal([]byte(option), &optionJson)
	if err != nil {
		println("error unmarshal option: " + err.Error())
		return
	}

	actualTimeApi, err := getActualTime()
	if err == nil {
		if actualTimeApi.Hour == optionJson.Hour && actualTimeApi.Minute == optionJson.Minute {
			println("current time is ", actualTimeApi.Time)
			c <- "response" // send sum to c
		}
	} else {
		println("error get actual time" + err.Error())
	}
	time.Sleep(15 * time.Second)
}

func (service *timerService) TimerReactionGiveTime(option string, idArea uint64) {
	println("give time")
}

func (service *timerService) GetServiceActionInfo() []schemas.Action {
	service.actionsName = append(service.actionsName, string(schemas.SpecificTime))
	return []schemas.Action{
		{
			Name:        string(schemas.SpecificTime),
			Description: "This action is a specific time action",
			Service:     service.serviceRepository.FindByName(schemas.Timer),
			Option:      "{\"hour\": 0, \"minute\": 0}",
		},
	}
}

func (service *timerService) GetServiceReactionInfo() []schemas.Reaction {
	service.reactionsName = append(service.reactionsName, string(schemas.GiveTime))
	return []schemas.Reaction{
		{
			Name:        string(schemas.GiveTime),
			Description: "This reaction is a give time reaction",
			Service:     service.serviceRepository.FindByName(schemas.Timer),
			Option:      "{}",
		},
	}
}

func (service *timerService) GetActionsName() []string {
	return service.actionsName
}

func (service *timerService) GetReactionsName() []string {
	return service.reactionsName
}
