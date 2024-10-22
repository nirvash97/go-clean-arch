package schedule

import (
	"time"
)

type Schedule struct {
	Date    time.Time         `json:"date"`
	Display bool              `json:"display"`
	Meeting []ScheduleMeeting `json:"meeting"`
}

type ScheduleMeeting struct {
	MeetingTime string `json:"meeting_time"`
	Status      string `json:"status"`
}

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"` // double 32 bit
}
