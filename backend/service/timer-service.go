package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type TimerService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Service specific functions
	// Actions functions
	TimerActionSpecificHour(c chan string, option json.RawMessage, area schemas.Area)
	// Reactions functions
	TimerReactionGiveTime(option json.RawMessage, area schemas.Area) string
}

type timerService struct {
	repository        repository.TimerRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	serviceInfo       schemas.Service
}

func NewTimerService(
	repository repository.TimerRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
) TimerService {
	return &timerService{
		repository:        repository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
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

func (service *timerService) FindActionByName(
	name string,
) func(c chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.SpecificTime):
		return service.TimerActionSpecificHour
	default:
		return nil
	}
}

func (service *timerService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {
	case string(schemas.GiveTime):
		return service.TimerReactionGiveTime
	default:
		return nil
	}
}

func (service *timerService) GetServiceActionInfo() []schemas.Action {
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
			Name:               string(schemas.SpecificTime),
			Description:        "This action is a specific time action",
			Service:            service.serviceInfo,
			Option:             option,
			MinimumRefreshRate: 10,
		},
	}
}

func (service *timerService) GetServiceReactionInfo() []schemas.Reaction {
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
	area schemas.Area,
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
		time.Sleep(time.Second)
		return
	}

	databaseStored := schemas.TimerActionSpecificHourStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.TimerActionSpecificHourStorage{
				Time: time.Now(),
			}
			area.StorageVariable, err = json.Marshal(databaseStored)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return
			}
		}
	}

	if databaseStored.Time.IsZero() {
		println("initializing storage variable")
		databaseStored = schemas.TimerActionSpecificHourStorage{
			Time: time.Now(),
		}
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
	}

	// generate time.Time from actualTimeApi
	actualTime := time.Date(
		actualTimeApi.Year,
		time.Month(actualTimeApi.Month),
		actualTimeApi.Day,
		actualTimeApi.Hour,
		actualTimeApi.Minute,
		actualTimeApi.Seconds,
		actualTimeApi.MilliSeconds,
		time.Local,
	)

	if databaseStored.Time.Before(actualTime) {
		if actualTime.Hour() == optionJSON.Hour && actualTimeApi.Minute == optionJSON.Minute {
			response := "current time is " + actualTimeApi.Time
			databaseStored.Time = time.Now().Add(time.Minute)
			area.StorageVariable, err = json.Marshal(databaseStored)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return
			}
			println(response)
			c <- response
		}
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions

func (service *timerService) TimerReactionGiveTime(
	option json.RawMessage,
	area schemas.Area,
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
