package schemas

type OpenweathermapAction string

const (
	SpecificWeather     OpenweathermapAction = "SpecificWeather"
	SpecificTemperature OpenweathermapAction = "SpecificTemperature"
)

type OpenweathermapReaction string

const (
	CurrentWeather     OpenweathermapReaction = "CurrentWeather"
	CurrentTemperature OpenweathermapReaction = "CurrentTemperature"
)

type OpenweathermapActionSpecificWeather struct {
	City    string `json:"city"`
	Weather string `json:"weather"`
}

type OpenweathermapActionSpecificTemperature struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
}

type OpenweathermapReactionCurrentWeather struct {
	City string `json:"city"`
}

type OpenweathermapCityCoordinatesResponse struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

type OpenweathermapCoordinatesWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
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

type OpenweathermapReactionGiveTime struct{}

type OpenweathermapReactionApiResponse struct{}
