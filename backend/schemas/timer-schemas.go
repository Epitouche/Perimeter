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
