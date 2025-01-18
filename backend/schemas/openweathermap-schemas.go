package schemas

import "errors"

type OpenWeatherMapAction string // The action type for OpenWeatherMap.

// SpecificWeather is the action to check the weather of a specific city.
const (
	SpecificWeather     OpenWeatherMapAction = "SpecificWeather"     // SpecificWeather is the action to check the weather of a specific city.
	SpecificTemperature OpenWeatherMapAction = "SpecificTemperature" // SpecificTemperature is the action to check the temperature of a specific city.
	AboveTemperature    OpenWeatherMapAction = "AboveTemperature"    // AboveTemperature is the action to check if the temperature is above a certain value.
	BelowTemperature    OpenWeatherMapAction = "BelowTemperature"    // BelowTemperature is the action to check if the temperature is below a certain value.
)

type OpenWeatherMapReaction string // The reaction type for OpenWeatherMap.

const (
	CurrentWeather     OpenWeatherMapReaction = "CurrentWeather"     // CurrentWeather is the reaction to get the current weather.
	CurrentTemperature OpenWeatherMapReaction = "CurrentTemperature" // CurrentTemperature is the reaction to get the current temperature.
)

// https://openweathermap.org/weather-conditions

type WeatherCondition string // WeatherCondition is a string type to represent the weather condition.

const (
	// Thunderstorm.
	Thunderstorm WeatherCondition = "Thunderstorm" // Thunderstorm is the weather condition for thunderstorms.
	// Drizzle.
	Drizzle WeatherCondition = "Drizzle" // Drizzle is the weather condition for drizzles.
	// Rain.
	Rain WeatherCondition = "Rain" // Rain is the weather condition for rain.
	// Snow.
	Snow WeatherCondition = "Snow" // Snow is the weather condition for snow.
	// Atmosphere.
	Mist    WeatherCondition = "Mist"    // Mist is the weather condition for mist.
	Smoke   WeatherCondition = "Smoke"   // Smoke is the weather condition for smoke.
	Haze    WeatherCondition = "Haze"    // Haze is the weather condition for haze.
	Dust    WeatherCondition = "Dust"    // Dust is the weather condition for dust.
	Fog     WeatherCondition = "Fog"     // Fog is the weather condition for fog.
	Sand    WeatherCondition = "Sand"    // Sand is the weather condition for sand.
	Ash     WeatherCondition = "Ash"     // Ash is the weather condition for ash.
	Squall  WeatherCondition = "Squall"  // Squall is the weather condition for squall.
	Tornado WeatherCondition = "Tornado" // Tornado is the weather condition for tornado.
	// Clear.
	Clear WeatherCondition = "Clear" // Clear is the weather condition for clear skies.
	// Clouds.
	Clouds WeatherCondition = "Clouds" // Clouds is the weather condition for cloudy skies.
)

type OpenWeatherMapStorageVariable int // OpenWeatherMapStorageVariable is an integer type to represent the storage variable.

const (
	OpenWeatherMapStorageVariableInit  OpenWeatherMapStorageVariable = 0 // OpenWeatherMapStorageVariableInit is the initial storage variable.
	OpenWeatherMapStorageVariableTrue  OpenWeatherMapStorageVariable = 1 // OpenWeatherMapStorageVariableTrue is the storage variable for true.
	OpenWeatherMapStorageVariableFalse OpenWeatherMapStorageVariable = 2 // OpenWeatherMapStorageVariableFalse is the storage variable for false.
)

type OpenWeatherMapActionSpecificWeather struct {
	City    string           `json:"city"`    // The city to check the weather for.
	Weather WeatherCondition `json:"weather"` // The weather condition to check.
}

type OpenWeatherMapActionSpecificTemperature struct {
	City        string `json:"city"`        // The city to check the temperature for.
	Temperature int64  `json:"temperature"` // The temperature to check.
}

// all reaction options schema.
type OpenWeatherMapReactionOption struct {
	City string `json:"city"` // The city to get the weather for.
}

type OpenWeatherMapCityCoordinatesResponse struct {
	Name    string  `json:"name"`    // The name of the city.
	Lat     float64 `json:"lat"`     // The latitude of the city.
	Lon     float64 `json:"lon"`     // The longitude of the city.
	Country string  `json:"country"` // The country of the city.
	State   string  `json:"state"`   // The state of the city.
}

type OpenWeatherMapCoordinatesWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int              `json:"id"`
		Main        WeatherCondition `json:"main"`
		Description string           `json:"description"`
		Icon        string           `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type OpenWeatherMapReactionGiveTime struct{}

type OpenWeatherMapReactionApiResponse struct{}

var ErrOpenWeatherMapApiKeyNotSet = errors.New("OPENWEATHERMAP_API_KEY is not set")
