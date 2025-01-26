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

// openWeatherMapService provides methods to interact with the OpenWeatherMap API,
// manage service-related data, and handle area-specific information.
//
// Fields:
// - repository: Repository to access the OpenWeatherMap API.
// - serviceRepository: Repository to access the Service entity.
// - areaRepository: Repository to access the Area entity.
// - serviceInfo: Information about the OpenWeatherMap service.
type openWeatherMapService struct {
	repository        repository.OpenWeatherMapRepository // Repository to access the OpenWeatherMap API
	serviceRepository repository.ServiceRepository        // Repository to access the Service entity
	areaRepository    repository.AreaRepository           // Repository to access the Area entity
	serviceInfo       schemas.Service                     // Information about the OpenWeatherMap service
}

// NewOpenWeatherMapService creates a new instance of OpenWeatherMapService with the provided repositories.
// It initializes the service with predefined service information including name, description, OAuth support, color, and icon.
//
// Parameters:
//   - repository: an instance of OpenWeatherMapRepository for accessing weather data.
//   - serviceRepository: an instance of ServiceRepository for managing service-related data.
//   - areaRepository: an instance of AreaRepository for managing area-related data.
//
// Returns:
//   - OpenWeatherMapService: a new instance of OpenWeatherMapService.
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

// GetServiceInfo returns the service information for the OpenWeatherMap service.
// It retrieves the service details encapsulated in the schemas.Service struct.
func (service *openWeatherMapService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

// FindActionByName returns a function that performs a specific action based on the provided name.
// The returned function takes a channel, an option in the form of json.RawMessage, and an area of type schemas.Area.
//
// Parameters:
//   - name: A string representing the name of the action.
//
// Returns:
//   - A function that takes a channel, an option, and an area, and performs the corresponding action.
//   - If the name does not match any known actions, it returns nil.
func (service *openWeatherMapService) FindActionByName(
	name string,
) func(channel chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.SpecificWeather):
		return service.OpenWeatherMapActionSpecificWeather
	case string(schemas.SpecificTemperature):
		return service.OpenWeatherMapActionSpecificTemperature
	case string(schemas.AboveTemperature):
		return service.OpenWeatherMapActionAboveTemperature
	case string(schemas.BelowTemperature):
		return service.OpenWeatherMapActionBelowTemperature
	default:
		return nil
	}
}

// FindReactionByName returns a function that corresponds to the given reaction name.
// The returned function takes a JSON raw message and an area schema as parameters and returns a string.
// If the reaction name matches "CurrentWeather" or "CurrentTemperature", the corresponding function is returned.
// If the reaction name does not match any known reactions, nil is returned.
//
// Parameters:
//   - name: The name of the reaction to find.
//
// Returns:
//   - A function that takes a JSON raw message and an area schema, and returns a string.
//   - Nil if the reaction name does not match any known reactions.
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

// GetServiceActionInfo retrieves a list of actions supported by the OpenWeatherMap service.
// It initializes default values for specific weather and temperature actions, marshals them into JSON,
// and fetches the service information from the repository. The function returns a slice of Action
// structs, each representing a different type of weather-related action.
//
// Returns:
//
//	[]schemas.Action: A slice of Action structs containing the details of each supported action.
//
// Possible errors:
//   - If there is an error during JSON marshaling of the default values, an error message is printed.
//   - If there is an error finding the service information by name, an error message is printed.
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
			Description:        "This action triggers when the temperature is a specific value",
			Service:            service.serviceInfo,
			Option:             optionSpecificTemperature,
			MinimumRefreshRate: 10,
		},
		{
			Name:               string(schemas.AboveTemperature),
			Description:        "This action triggers when the temperature is above a specific value",
			Service:            service.serviceInfo,
			Option:             optionSpecificTemperature,
			MinimumRefreshRate: 10,
		},
		{
			Name:               string(schemas.BelowTemperature),
			Description:        "This action triggers when the temperature is below a specific value",
			Service:            service.serviceInfo,
			Option:             optionSpecificTemperature,
			MinimumRefreshRate: 10,
		},
	}
}

// GetServiceReactionInfo retrieves the reaction information for the OpenWeatherMap service.
// It sets a default city to "Bordeaux" and marshals it into JSON format for the reaction options.
// The function updates the service information by finding the service by its name.
// It returns a slice of Reaction objects, each containing the name, description, service information, and options.
//
// Returns:
//
//	[]schemas.Reaction: A slice of Reaction objects with the current weather and temperature reactions.
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

// getCoordinatesOfCity retrieves the geographical coordinates (latitude and longitude) of a given city
// using the OpenWeatherMap API.
//
// Parameters:
//   - city: The name of the city for which to retrieve coordinates.
//
// Returns:
//   - coordinates: A struct containing the latitude and longitude of the city.
//   - err: An error if the coordinates could not be retrieved.
//
// The function requires an environment variable "OPENWEATHERMAP_API_KEY" to be set with a valid API key.
// If the API key is not set, it returns an error schemas.ErrOpenWeatherMapApiKeyNotSet.
//
// Example usage:
//
//	coordinates, err := getCoordinatesOfCity("London")
//	if err != nil {
//	    log.Fatalf("Error retrieving coordinates: %v", err)
//	}
//	fmt.Printf("Coordinates of London: Lat=%f, Lon=%f\n", coordinates.Lat, coordinates.Lon)
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

	defer resp.Body.Close()

	var result []schemas.OpenWeatherMapCityCoordinatesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return coordinates, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	coordinates.Lat = result[0].Lat
	coordinates.Lon = result[0].Lon
	return coordinates, nil
}

// getWeatherOfCoordinate fetches the weather information for the given coordinates
// from the OpenWeatherMap API.
//
// Parameters:
//   - coordinates: A struct containing the latitude (Lat) and longitude (Lon) of the location.
//
// Returns:
//   - weather: A schemas.OpenWeatherMapCoordinatesWeatherResponse struct containing the weather information.
//   - err: An error if the request fails or the API key is not set.
//
// The function retrieves the OpenWeatherMap API key from the environment variable "OPENWEATHERMAP_API_KEY".
// If the API key is not set, it returns an error. It constructs the API request URL with the provided coordinates
// and sends a GET request to the OpenWeatherMap API. The response is decoded into the weather struct and returned.
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

	defer resp.Body.Close()

	var result schemas.OpenWeatherMapCoordinatesWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return weather, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	weather = result

	return weather, nil
}

// initializedOpenWeatherMapStorageVariable initializes the OpenWeatherMap storage variable for a given area.
// It attempts to unmarshal the storage variable from the area's StorageVariable field. If unmarshaling fails,
// it initializes the storage variable to a default false value and updates the area's StorageVariable field.
// The function returns the initialized storage variable and any error encountered during the process.
//
// Parameters:
//   - area: The area for which the storage variable is being initialized.
//   - service: The OpenWeatherMap service used to update the area repository.
//
// Returns:
//   - variable: The initialized OpenWeatherMap storage variable.
//   - err: An error encountered during the initialization process, if any.
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

// OpenWeatherMapActionSpecificWeather retrieves the weather information for a specified city
// and updates the area storage variable based on the weather condition.
//
// Parameters:
// - channel: A channel to send the response string.
// - option: A JSON raw message containing the weather options.
// - area: The area schema containing the action and storage variable.
//
// The function performs the following steps:
// 1. Unmarshals the JSON option into OpenWeatherMapActionSpecificWeather struct.
// 2. Initializes the storage variable for the specified area.
// 3. Retrieves the coordinates of the specified city.
// 4. Fetches the weather information for the retrieved coordinates.
// 5. Compares the current weather with the specified weather condition.
// 6. Updates the storage variable and area repository based on the weather condition.
// 7. Sends the response string to the provided channel if the weather condition matches.
// 8. Sleeps for a duration based on the area's action refresh rate or minimum refresh rate.
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
		if weatherOfSpecifiedCity.Weather[0].Main == optionJSON.Weather {
			if variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableFalse {
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

	WaitAction(area)
}

// OpenWeatherMapActionSpecificTemperature checks the current temperature of a specified city
// and updates the storage variable accordingly. If the current temperature matches the specified
// temperature, it sends a response message to the provided channel. The function also handles
// updating the area repository with the new storage variable state.
//
// Parameters:
// - channel: A channel to send the response message if the temperature matches.
// - option: A JSON raw message containing the city and temperature to check against.
// - area: The area schema containing the action and storage variable information.
//
// The function performs the following steps:
//  1. Unmarshals the option JSON into an OpenWeatherMapActionSpecificTemperature struct.
//  2. Initializes the storage variable for the specified area.
//  3. Retrieves the coordinates of the specified city.
//  4. Gets the current weather information for the city's coordinates.
//  5. Compares the current temperature with the specified temperature and updates the storage
//     variable and area repository accordingly.
//  6. Sleeps for a duration based on the area's action refresh rate or minimum refresh rate.
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
	} else {
		if int64(math.Round(weatherOfSpecifiedCity.Main.Temp)) == optionJSON.Temperature {
			if variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableFalse {
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

	WaitAction(area)
}

// OpenWeatherMapActionAboveTemperature checks if the temperature in a specified city is above a given threshold.
// If the temperature is above the threshold and the storage variable indicates it was previously below, it sends a message to the provided channel and updates the storage variable.
// If the temperature is below the threshold and the storage variable indicates it was previously above, it updates the storage variable accordingly.
// The function also handles errors related to JSON unmarshalling, storage variable initialization, coordinate retrieval, and weather data retrieval.
// Parameters:
// - channel: A channel to send messages when the temperature condition is met.
// - option: A JSON raw message containing the city and temperature threshold.
// - area: The area schema containing action and storage variable information.
func (service *openWeatherMapService) OpenWeatherMapActionAboveTemperature(
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
	} else {
		if int64(math.Round(weatherOfSpecifiedCity.Main.Temp)) > optionJSON.Temperature {
			if variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableFalse {
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

	WaitAction(area)
}

// OpenWeatherMapActionBelowTemperature checks if the temperature of a specified city
// is below a given threshold and updates the area storage variable accordingly.
// If the temperature is below the threshold and the storage variable is false, it sends
// a message to the provided channel and updates the storage variable to true.
// If the temperature is above the threshold and the storage variable is true, it updates
// the storage variable to false.
//
// Parameters:
//   - channel: A channel to send the response message if the temperature is below the threshold.
//   - option: A JSON raw message containing the city and temperature threshold.
//   - area: The area schema containing the storage variable and action refresh rates.
//
// The function performs the following steps:
//  1. Unmarshals the option JSON into an OpenWeatherMapActionSpecificTemperature struct.
//  2. Initializes the storage variable for the specified area.
//  3. Retrieves the coordinates of the specified city.
//  4. Gets the weather information for the city's coordinates.
//  5. Checks if the temperature is below the threshold and updates the storage variable and area accordingly.
//  6. Sleeps for the minimum refresh rate or action refresh rate of the area.
func (service *openWeatherMapService) OpenWeatherMapActionBelowTemperature(
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
	} else {
		if int64(math.Round(weatherOfSpecifiedCity.Main.Temp)) < optionJSON.Temperature {
			if variableDatabaseStorage == schemas.OpenWeatherMapStorageVariableFalse {
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

	WaitAction(area)
}

// Reactions functions

// OpenWeatherMapReactionCurrentWeather retrieves the current weather for a specified city
// using the OpenWeatherMap API and returns a string describing the weather.
//
// Parameters:
//   - option: a JSON-encoded raw message containing the city name.
//   - area: a schemas.Area object (not used in the function).
//
// Returns:
//
//	A string describing the current weather in the specified city, or an error message
//	if there was an issue with unmarshalling the option or retrieving the weather information.
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

// OpenWeatherMapReactionCurrentTemperature retrieves the current temperature for a specified city
// using the OpenWeatherMap API and returns a formatted string with the temperature information.
//
// Parameters:
//   - option: A JSON raw message containing the city information.
//   - area: An area schema (not used in the current implementation).
//
// Returns:
//
//	A string containing the current temperature in the specified city or an error message if any error occurs during the process.
//
// The function performs the following steps:
//  1. Unmarshals the JSON option to extract the city information.
//  2. Retrieves the coordinates of the specified city.
//  3. Fetches the weather information for the retrieved coordinates.
//  4. Formats and returns the current temperature information or an error message if any step fails.
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
