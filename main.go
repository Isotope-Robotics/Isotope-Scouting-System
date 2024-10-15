// Author: pat@patfairbank.com (Patrick Fairbank)
// Modified for Isotope Robotics by: Ethen Brandenburg

package main

import (
	"github.com/Isotope-Robotics/Isotope-Scouting-System/web"
)

const httpPort = 8080

func main() {

	web := web.NewWeb()
	go web.ServeWebInterface(httpPort)
}
