// Author: pat@patfairbank.com (Patrick Fairbank)
// Modified for Isotope Robotics by: Ethen Brandenburg

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHelloFunc)
	http.ListenAndServe(":8081", nil)
}

func handleHelloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")

}
