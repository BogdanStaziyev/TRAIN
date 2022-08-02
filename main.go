package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"time"
)

type Trains []Train

type Train struct {
	TrainID            int       `json:"trainId"`
	DepartureStationID int       `json:"departureStationId"`
	ArrivalStationID   int       `json:"arrivalStationId"`
	Price              float32   `json:"price"`
	ArrivalTime        time.Time `json:"arrivalTime"`
	DepartureTime      time.Time `json:"departureTime"`
}

func main() {
	var departureStation, arrivalStation, criteria string
	fmt.Println("Введіть станцію відправлення: ")
	fmt.Scanln(&departureStation)

	fmt.Println("Введіть станцію прибуття: ")
	fmt.Scanln(&arrivalStation)

	fmt.Println("Критерій для сорутування: ")
	fmt.Scanln(&criteria)

	result, err := findTrains(departureStation, arrivalStation, criteria)
	if err != nil {
		log.Println(err)
	}
	for _, train := range result {
		fmt.Printf("{TrainID: %d, DepartureStationID: %d, ArrivalStationID: %d, Price: %.2f, ArrivalTime: time.Date(%d, time.%s, %d, %d, %d, %d, %d, time.%s), DepartureTime: time.Date(%d, time.%s, %d, %d, %d, %d, %d, time.%s)}\n", train.TrainID, train.DepartureStationID, train.ArrivalStationID, train.Price, train.ArrivalTime.Year(), train.ArrivalTime.Month().String(), train.ArrivalTime.Day(), train.ArrivalTime.Hour(), train.ArrivalTime.Minute(), train.ArrivalTime.Second(), train.ArrivalTime.Nanosecond(), train.ArrivalTime.Location(), train.DepartureTime.Year(), train.DepartureTime.Month().String(), train.DepartureTime.Day(), train.DepartureTime.Hour(), train.DepartureTime.Minute(), train.DepartureTime.Second(), train.DepartureTime.Nanosecond(), train.DepartureTime.Location())
	}
}

func findTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	if departureStation == "" {
		return nil, errors.New("empty departure station")
	} else if arrivalStation == "" {
		return nil, errors.New("empty arrival station")
	}

	sliceByte, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	//select all trains on a given route
	currentTrainsByStation, err := unmarshalByte(sliceByte, departureStation, arrivalStation)
	if err != nil {
		return nil, err
	}
	sortTrainsByStation, err := sortTrains(currentTrainsByStation, criteria)
	if err != nil {
		return nil, err
	}
	return sortTrainsByStation[:3], nil
}

func sortTrains(currentTrainsByStation Trains, criteria string) (Trains, error) {
	switch criteria {
	case "price":
		sort.SliceStable(currentTrainsByStation, func(i, j int) bool {
			return currentTrainsByStation[i].Price < currentTrainsByStation[j].Price
		})
	case "arrival-time":
		sort.SliceStable(currentTrainsByStation, func(i, j int) bool {
			return currentTrainsByStation[j].ArrivalTime.After(currentTrainsByStation[i].ArrivalTime)
		})
	case "departure-time":
		sort.SliceStable(currentTrainsByStation, func(i, j int) bool {
			return currentTrainsByStation[j].DepartureTime.After(currentTrainsByStation[i].DepartureTime)
		})
	default:
		return nil, errors.New("unsupported criteria")
	}
	return currentTrainsByStation, nil
}

type TrainsMapper []TrainMapper

type TrainMapper struct {
	TrainID            int     `json:"trainId"`
	DepartureStationID int     `json:"departureStationId"`
	ArrivalStationID   int     `json:"arrivalStationId"`
	Price              float32 `json:"price"`
	ArrivalTime        string  `json:"arrivalTime"`
	DepartureTime      string  `json:"departureTime"`
}

func unmarshalByte(sliceByte []byte, departureStation, arrivalStation string) (Trains, error) {
	var DepartureTime, ArrivalTime time.Time

	departureStationId, err := strconv.Atoi(departureStation)
	if err != nil || departureStationId <= 0 || len(arrivalStation) != 4 {
		return nil, errors.New("bad departure station input")
	}

	arrivalStationId, err := strconv.Atoi(arrivalStation)
	if err != nil || arrivalStationId <= 0 || len(arrivalStation) != 4 {
		return nil, errors.New("bad arrival station input")
	}

	var currentTrainsByStation Trains
	var convertTimeTrains TrainsMapper
	err = json.Unmarshal(sliceByte, &convertTimeTrains)
	if err != nil {
		log.Println("error Unmarshal")
	}
	for _, val := range convertTimeTrains {
		if departureStationId == val.DepartureStationID && arrivalStationId == val.ArrivalStationID {
			DepartureTime = parseTime(val.DepartureTime)
			ArrivalTime = parseTime(val.ArrivalTime)
			num := Train{val.TrainID, val.DepartureStationID, val.ArrivalStationID, val.Price, ArrivalTime, DepartureTime}
			currentTrainsByStation = append(currentTrainsByStation, num)
		}
	}
	return currentTrainsByStation, nil
}

func parseTime(date string) time.Time {
	returnDate, err := time.Parse("15:04:05", date)
	if err != nil {
		fmt.Println("error parseTime")
	}
	return returnDate
}
