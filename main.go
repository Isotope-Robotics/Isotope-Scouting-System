// Author: pat@patfairbank.com (Patrick Fairbank)
// Modified for Isotope Robotics by: Ethen Brandenburg

package main

import (
	"log"

	"github.com/Isotope-Robotics/Isotope-Scouting-System/Event"
	"github.com/Isotope-Robotics/Isotope-Scouting-System/web"
)

const httpPort = 8081
const eventDbPath = "./event.db"

func main() {

	_event, err := Event.NewEvent(eventDbPath)
	if err != nil {
		log.Fatalln("Error during startup: ", err)
	}

	web := web.NewWeb()
	go web.ServeWebInterface(httpPort)

	_event.Run()
}
