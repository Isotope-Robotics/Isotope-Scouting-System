package Event

import (
	"log"
	"time"

	"github.com/Isotope-Robotics/Isotope-Scouting-System/model"
)

const (
	eventLoopWarningMS = 3000
	eventLoopPeriodMs  = 10
)

type Event struct {
	EventSettings *model.EventSettings
}

func NewEvent(dbPath string) (*Event, error) {
	event := new(Event)

	return event, nil
}

// Runs the Event Loops for the Event
func (event *Event) Run() {
	for {
		loopStartTime := time.Now()

		if time.Since(loopStartTime).Microseconds() > eventLoopWarningMS {
			log.Printf("Warning: Event loop iteration took a long time: %dus", time.Since(loopStartTime).Microseconds())
		}

		time.Sleep(time.Millisecond * eventLoopPeriodMs)
	}
}

func Update() {

}
