package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Controller struct {
	ReqQueue  []TravelRequest
	TaxiQueue []Taxi
}

func (conn *Controller) init() {

	jsonFile, err := os.Open("data/taxis.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//data := TaxisNodes{}
	data := []Taxi{}
	_ = json.Unmarshal([]byte(byteValue), &data)

	//conn.TaxiQueue = append(conn.TaxiQueue, data.Taxis...)
	conn.TaxiQueue = append(conn.TaxiQueue, data...)
}

func (conn *Controller) addTravelRequest(req TravelRequest) bool {
	//func (conn *Controller) addTravelRequest(req *TravelRequest) bool {

	for _, r := range conn.ReqQueue {
		if r.TravelRequestID == req.TravelRequestID {
			return false
		}
	}

	conn.ReqQueue = append(conn.ReqQueue, req)
	//conn.ReqQueue = append(conn.ReqQueue, *req)

	// @todo: save in the json file.
	// @todo: implement DB.

	return true
}

func (conn *Controller) addTaxi(taxi Taxi) bool {

	for _, t := range conn.TaxiQueue {
		if t.TaxiID == taxi.TaxiID {
			return false
		}
	}

	filename := "data/taxis.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	data := []Taxi{}
	_ = json.Unmarshal([]byte(byteValue), &data)
	data = append(data, taxi)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}

	conn.TaxiQueue = append(conn.TaxiQueue, taxi)
	return true
}

func (conn *Controller) handleTravelRequest() (*Taxi, *TravelRequest) {

	if len(conn.TaxiQueue) == 0 || len(conn.ReqQueue) == 0 {
		return nil, nil
	}

	freeTaxi := conn.TaxiQueue[0]
	conn.TaxiQueue = conn.TaxiQueue[1:]

	conn.TaxiQueue = append(conn.TaxiQueue, freeTaxi) // add the taxi back to queue.

	req := conn.ReqQueue[0]
	conn.ReqQueue = conn.ReqQueue[1:]

	return &freeTaxi, &req
}
