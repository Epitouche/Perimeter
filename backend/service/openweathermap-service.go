package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type OpenWeatherMapService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(
		name string,
	) func(channel chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Service specific functions
	// Actions functions
	OpenWeatherMapActionSpecificWeather(
		channel chan string,
		option json.RawMessage,
		area schemas.Area,
	)
	OpenWeatherMapActionSpecificTemperature(
		channel chan string,
		option json.RawMessage,
		area schemas.Area,
	)
	// Reactions functions
	OpenWeatherMapReactionCurrentWeather(
		option json.RawMessage,
		area schemas.Area,
	) string
	OpenWeatherMapReactionCurrentTemperature(
		option json.RawMessage,
		area schemas.Area,
	) string
}

type openWeatherMapService struct {
	repository        repository.OpenWeatherMapRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	serviceInfo       schemas.Service
}

func NewOpenWeatherMapService(
	repository repository.OpenWeatherMapRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
) OpenWeatherMapService {
	return &openWeatherMapService{
		repository:        repository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.OpenWeatherMap,
			Description: "This service is a weather service",
			Oauth:       false,
			Color:       "#946500",
			Icon:        "https://api.iconify.design/mdi:weather-cloudy.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *openWeatherMapService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *openWeatherMapService) FindActionByName(
	name string,
) func(channel chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.SpecificWeather):
		return service.OpenWeatherMapActionSpecificWeather
	case string(schemas.SpecificTemperature):
		return service.OpenWeatherMapActionSpecificTemperature
	default:
		return nil
	}
}

func (service *openWeatherMapService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {

	case string(schemas.CurrentWeather):
		return service.OpenWeatherMapReactionCurrentWeather
	case string(schemas.CurrentTemperature):
		return service.OpenWeatherMapReactionCurrentTemperature
	default:
		return nil
	}
}

func (service *openWeatherMapService) GetServiceActionInfo() []schemas.Action {
	// SpecificWeather
	defaultValueSpecificWeather := schemas.OpenWeatherMapActionSpecificWeather{
		City:    "Bordeaux",
		Weather: "Rain",
	}
	optionSpecificWeather, err := json.Marshal(defaultValueSpecificWeather)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	// SpecificTemperature
	defaultValueSpecificTemperature := schemas.OpenWeatherMapActionSpecificTemperature{
		City:        "Bordeaux",
		Temperature: 12,
	}
	optionSpecificTemperature, err := json.Marshal(defaultValueSpecificTemperature)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}

	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.OpenWeatherMap,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:               string(schemas.SpecificWeather),
			Description:        "This action is a specific weather action",
			Service:            service.serviceInfo,
			Option:             optionSpecificWeather,
			MinimumRefreshRate: 10,
		},
		{
			Name:               string(schemas.SpecificTemperature),
			Description:        "This action is a specific temperature action",
			Service:            service.serviceInfo,
			Option:             optionSpecificTemperature,
			MinimumRefreshRate: 10,
		},
	}
}

func (service *openWeatherMapService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.OpenWeatherMapReactionOption{
		City: "Bordeaux",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.OpenWeatherMap,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.CurrentWeather),
			Description: "This reaction is a current weather reaction",
			Service:     service.serviceInfo,
			Option:      option,
		},
		{
			Name:        string(schemas.CurrentTemperature),
			Description: "This reaction is a current teamperature reaction",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

// Service specific functions

func getCoordinatesOfCity(city string) (coordinates struct {
	Lat float64
	Lon float64
}, err error,
) {
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if APIKey == "" {
		return coordinates, schemas.ErrOpenWeatherMapApiKeyNotSet
	}

	apiURL := "http://api.openweathermap.org/geo/1.0/direct"
	data := url.Values{}
	data.Set("q", city)
	data.Set("limit", "1")
	data.Set("appid", APIKey)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return coordinates, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return coordinates, fmt.Errorf("unable to make request because %w", err)
	}

	var result []schemas.OpenWeatherMapCityCoordinatesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return coordinates, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()

	coordinates.Lat = result[0].Lat
	coordinates.Lon = result[0].Lon
	return coordinates, nil
}

func getWeatherOfCoordinate(coordinates struct {
	Lat float64
	Lon float64
},
) (weather schemas.OpenWeatherMapCoordinatesWeatherResponse, err error) {
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if APIKey == "" {
		return weather, schemas.ErrOpenWeatherMapApiKeyNotSet
	}

	apiURL := "https://api.openweathermap.org/data/2.5/weather"
	data := url.Values{}
	data.Set("lat", fmt.Sprintf("%f", coordinates.Lat))
	data.Set("lon", fmt.Sprintf("%f", coordinates.Lon))
	data.Set("appid", APIKey)
	data.Set("units", "metric") // to get temperature in celsius

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return weather, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return weather, fmt.Errorf("unable to make request because %w", err)
	}

	var result schemas.OpenWeatherMapCoordinatesWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return weather, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()

	weather = result
	return weather, nil
}

func initializedOpenWeatherMapStorageVariable(
	area schemas.Area,
	service openWeatherMapService,
) (variable schemas.OpenWeatherMapStorageVariable, err error) {
	variable = schemas.OpenWeatherMapStorageVariableInit
	err = json.Unmarshal(area.StorageVariable, &variable)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshaling storage variable: " + err.Error())
			return variable, err
		} else {
			println("initializing storage variable")
			variable = schemas.OpenWeatherMapStorageVariableFalse
			area.StorageVariable, err = json.Marshal(variable)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return variable, err
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return variable, err
			}
		}
	}

	if variable == schemas.OpenWeatherMapStorageVariableInit {
		variable = schemas.OpenWeatherMapStorageVariableFalse
		area.StorageVariable, err = json.Marshal(variable)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return variable, err
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return variable, err
		}
	}
	return variable, nil
}

// Actions functions

func (service *openWeatherMapService) OpenWeatherMapActionSpecificWeather(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	// Find the area
	optionJSON := schemas.OpenWeatherMapActionSpecificWeather{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	variableDatabaseStorage, err := initializedOpenWeatherMapStorageVariable(area, *service)
	if err != nil {
		println("error initializing storage variable: " + err.Error())
	}

	coordinates, err := getCoordinatesOfCity(optionJSON.City)
	if err != nil {
		fmt.Println(err)
	}

	weatherOfSpecifiedCity, err := getWeatherOfCoordinate(coordinates)
	if err != nil {
		println("error get actual weather info" + err.Error())
	} else {
		if weatherOfSpecifiedCity.Weather[0].Main == optionJSON.Weather && variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableFalse {
			response := "current weather in " + optionJSON.City + " is " + string(weatherOfSpecifiedCity.Weather[0].Main)
			variableDatabaseStorage = schemas.OpenWeatherMapStorageVariableTrue
			area.StorageVariable, err = json.Marshal(variableDatabaseStorage)
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
			channel <- response
		} else {
			if variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableTrue {
				variableDatabaseStorage = schemas.OpenWeatherMapStorageVariableFalse
				area.StorageVariable, err = json.Marshal(variableDatabaseStorage)
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
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

func (service *openWeatherMapService) OpenWeatherMapActionSpecificTemperature(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	optionJSON := schemas.OpenWeatherMapActionSpecificTemperature{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal temperature option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	variableDatabaseStorage, err := initializedOpenWeatherMapStorageVariable(area, *service)
	if err != nil {
		println("error initializing storage variable: " + err.Error())
	}

	coordinates, err := getCoordinatesOfCity(optionJSON.City)
	if err != nil {
		fmt.Println(err)
	}
	weatherOfSpecifiedCity, err := getWeatherOfCoordinate(coordinates)
	if err != nil {
		println("error get actual temperature info" + err.Error())
		if int64(math.Round(weatherOfSpecifiedCity.Main.Temp)) == optionJSON.Temperature &&
			variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableFalse {
			response := "current temperature in " + optionJSON.City + " is " + fmt.Sprintf(
				"%f",
				weatherOfSpecifiedCity.Main.Temp,
			) + "°C"
			println(response)
			channel <- response
			variableDatabaseStorage = schemas.OpenWeatherMapStorageVariableTrue
			area.StorageVariable, err = json.Marshal(variableDatabaseStorage)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return
			}
		} else {
			if variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableTrue {
				variableDatabaseStorage = schemas.OpenWeatherMapStorageVariableFalse
				area.StorageVariable, err = json.Marshal(variableDatabaseStorage)
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
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions

func (service *openWeatherMapService) OpenWeatherMapReactionCurrentWeather(
	option json.RawMessage,
	area schemas.Area,
) string {
	optionJSON := schemas.OpenWeatherMapReactionOption{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return "error unmarshal weather option: " + err.Error()
	}

	coordinates, err := getCoordinatesOfCity(optionJSON.City)
	if err != nil {
		fmt.Println(err)
	}

	weatherOfSpecifiedCity, err := getWeatherOfCoordinate(coordinates)
	if err != nil {
		println("error get actual weather info" + err.Error())
		return "error get actual weather info" + err.Error()
	} else {
		response := "current weather in " + optionJSON.City + " is " + string(weatherOfSpecifiedCity.Weather[0].Main)
		println(response)
		return response
	}
}

func (service *openWeatherMapService) OpenWeatherMapReactionCurrentTemperature(
	option json.RawMessage,
	area schemas.Area,
) string {
	optionJSON := schemas.OpenWeatherMapReactionOption{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal temperature option: " + err.Error())
		time.Sleep(time.Second)
		return "error unmarshal temperature option: " + err.Error()
	}
	coordinates, err := getCoordinatesOfCity(optionJSON.City)
	if err != nil {
		fmt.Println(err)
	}
	weatherOfSpecifiedCity, err := getWeatherOfCoordinate(coordinates)
	if err != nil {
		println("error get actual temperature info" + err.Error())
		return "error get actual temperature info" + err.Error()
	} else {
		response := "current temperature in " + optionJSON.City + " is " + fmt.Sprintf("%f", weatherOfSpecifiedCity.Main.Temp) + "°C"
		println(response)
		return response
		// TODO: save to database
	}
}
