package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Duration struct {
	// Duration in nanoseconds
	StartingTime int64
	EndingTime   int64
	Distance     float64
}

type Device struct {
	// Device ID
	ID uuid.UUID
	// Device Name
	Name string
	// Device Distance
	Distance float64
}

const BluetoothDisConst = 0.00000029981 / 2

// const BluetoothDisConst = 0.00029981 / 2

var initTime int64

func DistanceCal(time1, time2 int64) float64 {
	Duration := time2 - time1
	Distance := BluetoothDisConst * float64(Duration) // pseudo number for algo
	return Distance
}

func currentDistancePage(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UnixNano()
	// fmt.Println("Current Timestamp: ", currentTime)
	dis := DistanceCal(initTime, currentTime)
	newDuration := Duration{initTime, currentTime, dis}
	json.NewEncoder(w).Encode(newDuration)
	// fmt.Fprintf(w, "Current Distance: %f\n", dis)
	fmt.Println("Endpoint Hit: currentDistancePage")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Init Time reset!\n")
	initTime = time.Now().UnixNano()
	fmt.Println("Endpoint Hit: homePage")
}

func AllDevices(w http.ResponseWriter, r *http.Request) {
	pseudoDevices := []Device{
		Device{uuid.New(), "Device1", 0.0},
		Device{uuid.New(), "Device2", 30.0},
		Device{uuid.New(), "Device3", 27.0},
	}
	json.NewEncoder(w).Encode(pseudoDevices)
	fmt.Println("Endpoint Hit: AllDevices")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/currentDistance", currentDistancePage)
	myRouter.HandleFunc("/devices", AllDevices)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	// currentTime1 := time.Now().UnixNano()
	// fmt.Println("Current Timestamp: ", currentTime1)
	// time.Sleep(30 * time.Microsecond) // Simulate bluetooth connection and passing time
	// currentTime2 := time.Now().UnixNano()
	// fmt.Println("Current Timestamp: ", currentTime2)
	// dis := DistanceCal(currentTime1, currentTime2)
	// fmt.Println(dis)
	initTime = time.Now().UnixNano()
	fmt.Println("REST API with Mux Routers")
	handleRequests()
}
