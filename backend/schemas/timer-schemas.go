package schemas

type TimerAction string

const (
	SpecificTime        TimerAction = "SpecificTime"
)

type TimerActionSpecificHour struct {
	Hour   uint8 `json:"hour"`
	Minute uint8 `json:"minute"`
}