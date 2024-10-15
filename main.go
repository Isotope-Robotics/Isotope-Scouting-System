package main

import (
	"github.com/Isotope-Robotics/Isotope-Scouting-System/web"
)

const httpPort = 8080

func main() {

	web := web.NewWeb()
	go web.ServeWebInterface(httpPort)
}
