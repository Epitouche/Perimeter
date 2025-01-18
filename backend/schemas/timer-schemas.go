package schemas

import "time"

type TimerAction string

const (
	SpecificTime TimerAction = "SpecificTime"
)

type TimerReaction string

const (
	GiveTime TimerReaction = "GiveTime" // GiveTime reaction
)

// TimerActionSpecificHour represents a specific time action with an hour and minute.
// Hour specifies the hour of the action (0-23).
// Minute specifies the minute of the action (0-59).
type TimerActionSpecificHour struct {
	Hour   int `json:"hour"`   // The hour
	Minute int `json:"minute"` // The minute
}

type TimerActionSpecificHourStorage struct {
	Time time.Time `json:"time"` // The time
}

type TimerReactionGiveTime struct{}

type TimeApiResponse struct {
	Year         int    `json:"year"`         // The year
	Month        int    `json:"month"`        // The month
	Day          int    `json:"day"`          // The day
	Hour         int    `json:"hour"`         // The hour
	Minute       int    `json:"minute"`       // The minute
	Seconds      int    `json:"seconds"`      // The seconds
	MilliSeconds int    `json:"milliSeconds"` // The milliseconds
	DateTime     string `json:"dateTime"`     // The date and time
	Date         string `json:"date"`         // The date
	Time         string `json:"time"`         // The time
	TimeZone     string `json:"timeZone"`     // The time zone
	DayOfWeek    string `json:"dayOfWeek"`    // The day of the week
	DstActive    bool   `json:"dstActive"`    // The daylight saving time status
}
