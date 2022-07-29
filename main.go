package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type TrainsDTO []TrainDTO

type TrainDTO struct {
	TrainID            int     `json:"trainId"`
	DepartureStationID int     `json:"departureStationId"`
	ArrivalStationID   int     `json:"arrivalStationId"`
	Price              float32 `json:"price"`
	ArrivalTime        string  `json:"arrivalTime"`
	DepartureTime      string  `json:"departureTime"`
}

type Trains []TrainDTO

type Train struct {
	TrainID            int       `json:"trainId"`
	DepartureStationID int       `json:"departureStationId"`
	ArrivalStationID   int       `json:"arrivalStationId"`
	Price              float32   `json:"price"`
	ArrivalTime        time.Time `json:"arrivalTime"`
	DepartureTime      time.Time `json:"departureTime"`
}

func main() {
	sliseBite, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println("error read all")
	}
	b := &TrainsDTO{}
	err = json.Unmarshal(sliseBite, &b)
	if err != nil {
		log.Println("error Unmarshal")
	}
	for _, val := range *b {
		if val.TrainID == 8801 {
			fmt.Println(val)
		}
	}

	//	... запит даних від користувача
	//result, err := FindTrains(departureStation, arrivalStation, criteria))
	//	... обробка помилки
	//	... друк result
}

func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	// ... код
	return nil, nil // маєте повернути правильні значення
}
