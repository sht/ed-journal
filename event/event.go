package event

import (
	"time"
)

// Event represents an incoming event from the journal
type Event struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
