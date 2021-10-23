package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// How to return a ststus code from endpoint:
// w.WriteHeader(http.StatusNotFound)

var con Controller

func main() {
	// Init router
	r := mux.NewRouter()

	//con := initController()
	con.init()

	// Route handles & endpoints
	r.HandleFunc("/api/taxis", getTaxis).Methods("GET")
	r.HandleFunc("/api/taxi", createTaxi).Methods("POST")
	r.HandleFunc("/api/taxi/{id}", getTaxi).Methods("GET")
	r.HandleFunc("/api/travel_requests", getTravelRequests).Methods("GET")
	r.HandleFunc("/api/travel_request/{id}", getTravelRequest).Methods("GET")
	r.HandleFunc("/api/travel_request", createTravelRequest).Methods("POST")
	r.HandleFunc("/api/make_travel_request", makeTravelRequest).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getTaxis(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(con.TaxiQueue)
}

func getTaxi(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, taxi := range con.TaxiQueue {
		if taxi.TaxiID == params["id"] {
			json.NewEncoder(w).Encode(taxi)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("Taxi with the id: %s, is not fuond", params["id"]))
}

func createTaxi(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var taxi Taxi
	_ = json.NewDecoder(r.Body).Decode(&taxi)

	if !con.addTaxi(taxi) {
		json.NewEncoder(w).Encode(fmt.Sprintf("This taxi (id = %s) is already exists.\nThe taxi was not added", taxi.TaxiID))
		return
	}

	json.NewEncoder(w).Encode(taxi)
}

func getTravelRequests(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(con.ReqQueue)
}

func getTravelRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, req := range con.ReqQueue {
		if req.TravelRequestID == params["id"] {
			json.NewEncoder(w).Encode(req)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("Travel Request with the id: %s, is not fuond", params["id"]))
}

func createTravelRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req TravelRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	if !con.addTravelRequest(req) {
		json.NewEncoder(w).Encode(fmt.Sprintf("This travel request id (id = %s) is already exists.\nThe travel request was not added", req.TravelRequestID))
		return
	}

	json.NewEncoder(w).Encode(req)
}

func makeTravelRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taxi, req := con.handleTravelRequest()
	if taxi == nil || req == nil {
		// failure
		json.NewEncoder(w).Encode("An error occurred and the request was not processed..\nPlease try again.")
		return
	}

	// success
	re := MadeTravelRequest{*req, *taxi, con.TaxiQueue, con.ReqQueue}
	json.NewEncoder(w).Encode(re)
}
