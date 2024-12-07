package schemas

type TimerAction string

const (
	SpecificTime TimerAction = "SpecificTime"
)

type TimerReaction string

const (
	GiveTime TimerReaction = "GiveTime"
)

type TimerActionSpecificHour struct {
	Hour   uint8 `json:"hour"`
	Minute uint8 `json:"minute"`
}

type TimerReactionGiveTime struct{}

type TimeAPISTRUCT struct {
	Year         int    `json:"year"`
	Month        int    `json:"month"`
	Day          int    `json:"day"`
	Hour         uint8  `json:"hour"`
	Minute       uint8  `json:"minute"`
	Seconds      uint8  `json:"seconds"`
	MilliSeconds int    `json:"milliSeconds"`
	DateTime     string `json:"dateTime"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	TimeZone     string `json:"timeZone"`
	DayOfWeek    string `json:"dayOfWeek"`
	DstActive    bool   `json:"dstActive"`
}
