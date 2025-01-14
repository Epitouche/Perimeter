package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type OpenweathermapService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(
		name string,
	) func(channel chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Service specific functions
	// Actions functions
	OpenweathermapActionSpecificWeather(
		channel chan string,
		option json.RawMessage,
		area schemas.Area,
	)
	OpenweathermapActionSpecificTemperature(
		channel chan string,
		option json.RawMessage,
		area schemas.Area,
	)
	// Reactions functions
	OpenweathermapReactionCurrentWeather(
		option json.RawMessage,
		area schemas.Area,
	) string
	OpenweathermapReactionCurrentTemperature(
		option json.RawMessage,
		area schemas.Area,
	) string
}

type openweathermapService struct {
	repository        repository.OpenweathermapRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	serviceInfo       schemas.Service
}

func NewOpenweathermapService(
	repository repository.OpenweathermapRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
) OpenweathermapService {
	return &openweathermapService{
		repository:        repository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Openweathermap,
			Description: "This service is a weather service",
			Oauth:       false,
			Color:       "#946500",
			Icon:        "https://api.iconify.design/mdi:weather-cloudy.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *openweathermapService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *openweathermapService) FindActionByName(
	name string,
) func(channel chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.SpecificWeather):
		return service.OpenweathermapActionSpecificWeather
	case string(schemas.SpecificTemperature):
		return service.OpenweathermapActionSpecificTemperature
	default:
		return nil
	}
}

func (service *openweathermapService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {

	case string(schemas.CurrentWeather):
		return service.OpenweathermapReactionCurrentWeather
	case string(schemas.CurrentTemperature):
		return service.OpenweathermapReactionCurrentTemperature
	default:
		return nil
	}
}

func (service *openweathermapService) GetServiceActionInfo() []schemas.Action {
	// SpecificWeather
	defaultValueSpecificWeather := schemas.OpenweathermapActionSpecificWeather{
		City:    "Bordeaux",
		Weather: "Rain",
	}
	optionSpecificWeather, err := json.Marshal(defaultValueSpecificWeather)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	// SpecificTemperature
	defaultValueSpecificTemperature := schemas.OpenweathermapActionSpecificTemperature{
		City:        "Bordeaux",
		Temperature: 12,
	}
	optionSpecificTemperature, err := json.Marshal(defaultValueSpecificTemperature)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}

	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Openweathermap,
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

func (service *openweathermapService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.OpenweathermapReactionOption{
		City: "Bordeaux",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Openweathermap,
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

	var result []schemas.OpenweathermapCityCoordinatesResponse
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

func getWeatherOfCoodinate(coordinates struct {
	Lat float64
	Lon float64
},
) (weather schemas.OpenweathermapCoordinatesWeatherResponse, err error) {
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

	var result schemas.OpenweathermapCoordinatesWeatherResponse
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

// Actions functions

func (service *openweathermapService) OpenweathermapActionSpecificWeather(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	// Find the area
	optionJSON := schemas.OpenweathermapActionSpecificWeather{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	coordinates, err := getCoordinatesOfCity(optionJSON.City)
	if err != nil {
		fmt.Println(err)
	}

	weatherOfSpecifiedCity, err := getWeatherOfCoodinate(coordinates)
	if err != nil {
		println("error get actual weather info" + err.Error())
	} else {
		if weatherOfSpecifiedCity.Weather[0].Main == optionJSON.Weather {
			response := "current weather in " + optionJSON.City + " is " + string(weatherOfSpecifiedCity.Weather[0].Main)
			println(response)
			channel <- response
		}
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

func (service *openweathermapService) OpenweathermapActionSpecificTemperature(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
	optionJSON := schemas.OpenweathermapActionSpecificTemperature{}

	err := json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal temperature option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	coordinates, err := getCoordinatesOfCity(optionJSON.City)
	if err != nil {
		fmt.Println(err)
	}
	weatherOfSpecifiedCity, err := getWeatherOfCoodinate(coordinates)

	if err != nil {
		println("error get actual temperature info" + err.Error())
	} else {
		if weatherOfSpecifiedCity.Main.Temp == optionJSON.Temperature {
			response := "current temperature in " + optionJSON.City + " is " + fmt.Sprintf("%f", weatherOfSpecifiedCity.Main.Temp) + "°C"
			println(response)
			channel <- response
		}
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions

func (service *openweathermapService) OpenweathermapReactionCurrentWeather(
	option json.RawMessage,
	area schemas.Area,
) string {
	optionJSON := schemas.OpenweathermapReactionOption{}

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

	weatherOfSpecifiedCity, err := getWeatherOfCoodinate(coordinates)
	if err != nil {
		println("error get actual weather info" + err.Error())
		return "error get actual weather info" + err.Error()
	} else {
		response := "current weather in " + optionJSON.City + " is " + string(weatherOfSpecifiedCity.Weather[0].Main)
		println(response)
		return response
	}
}

func (service *openweathermapService) OpenweathermapReactionCurrentTemperature(
	option json.RawMessage,
	area schemas.Area,
) string {
	optionJSON := schemas.OpenweathermapReactionOption{}

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
	weatherOfSpecifiedCity, err := getWeatherOfCoodinate(coordinates)
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
