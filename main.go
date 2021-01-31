package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/pin/{pin:[0-9]+}/toggle", togglePinHandler)

	log.Fatal(http.ListenAndServe(":8081", cors.Handler(router)))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, raspberry!")
}

func togglePinHandler(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	pinNum, err := strconv.Atoi(requestVars["pin"])

	if err != nil {
		http.Error(w, "Could not parse pin number", http.StatusBadRequest)
		return
	}

	rpio.Open()
	log.Printf("Toggling pin %d", pinNum)
	pin := rpio.Pin(pinNum)
	pin.Output()
	pin.Toggle()
	rpio.Close()

	fmt.Fprintf(w, "Toggle pin %d", pinNum)
}
