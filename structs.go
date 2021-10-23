package main

type TaxisNodes struct {
	Taxis []Taxi `json:"taxis"`
}

type Taxi struct {
	TaxiID string `json:"taxiid"`
	IsFree bool   `json:"isfree"`
}

type TravelRequest struct {
	TravelRequestID string `json:"travelrequestid"`
	Source          string `json:"source"`
	Dest            string `json:"dest"`
}

type MadeTravelRequest struct {
	TravelReq          TravelRequest   `json:"travel_request"`
	Taxi               Taxi            `json:"taxi"`
	FreeTaxis          []Taxi          `json:"free_taxis"`
	TravelRequestQueue []TravelRequest `json:"travel_requests_queue"`
}

type ControllerInterface interface {
	init()
	addTravelRequest(req TravelRequest) bool
	handleTravelRequest() (*Taxi, *TravelRequest)
	addTaxi(taxi Taxi) bool
}

//func interfaceChecking(c ControllerInterface) {
//	fmt.Println(c)
//}
