package schemas

type OpenweathermapAction string

const (
	CurrentWeatherAction OpenweathermapAction = "CurrentWeather"
)

type OpenweathermapReaction string

const (
	CurrentWeatherReaction OpenweathermapReaction = "CurrentWeather"
)

type Openweather struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type OpenweathermapReactionGiveTime struct{}

type OpenweathermapReactionApiResponse struct{}
