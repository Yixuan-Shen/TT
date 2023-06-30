package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	Distance_m float64
	//phone number
}

const BluetoothDisConst = 0.00000029981 / 2

// const BluetoothDisConst = 0.00029981 / 2

var initTime int64

// Calculate for the oneway distance between two devices
func distanceCal_oneway(sendT, receiveT int64) (distance float64) {
	electronmaticWave := 0.00029981
	receiveTime := receiveT - sendT
	distance = float64(receiveTime) * electronmaticWave
	fmt.Printf("The distance from the device1 is %.2f M.\n", distance)
	return
}

// One way distacne runner function
func onewayDistancePage(w http.ResponseWriter, r *http.Request) {
	sendTime := time.Now().UnixNano()
	fmt.Println("The time sent by another device is ->", sendTime)
	receiveTime := time.Now().UnixNano()
	fmt.Println("The time received by our device is ->", receiveTime)
	distance := distanceCal_oneway(sendTime, receiveTime)
	newDuration := Duration{sendTime, receiveTime, distance}
	json.NewEncoder(w).Encode(newDuration)
	fmt.Println("The distance between two devices is ->", distance)
	fmt.Println("Endpoint Hit: current_onewayDistancePage")
}

// Calculate for the round distance between two devices
func DistanceCal(time1, time2 int64) float64 {
	Duration := time2 - time1
	Distance := BluetoothDisConst * float64(Duration) // pseudo number for algo
	return Distance
}

// The round distacne runner function
func currentDistancePage(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UnixNano()
	// fmt.Println("Current Timestamp: ", currentTime)
	dis := DistanceCal(initTime, currentTime)
	newDuration := Duration{initTime, currentTime, dis}
	json.NewEncoder(w).Encode(newDuration)
	// fmt.Fprintf(w, "Current Distance: %f\n", dis)
	fmt.Println("Endpoint Hit: currentDistancePage")
}

// Show all the devices on the API
func AllDevices(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(devices)
	fmt.Println("Endpoint Hit: AllDevices")
}

// INitialize the deviceList
func initializeDL() (deviceList *list.List) {
	deviceList = list.New()
	return deviceList
}

func updateDeviceList(nearbyDeviceList *list.List, name string, distance float64) *list.List {
	nearbyDeviceList.PushBack(Device{uuid.New(), name, distance})
	return nearbyDeviceList
}

func clearDeviceList(deviceList *list.List) {
	deviceList = list.New()
}

// Search and connect the nearby device
func connection(w http.ResponseWriter, r *http.Request) {
	nearbyDeviceList := initializeDL() //create a new list
	var numberOfDevice int64
	fmt.Printf("Type the number of device you want to search:\n")
	fmt.Scanln(&numberOfDevice)
	var index int64
	//append new element in list for numberOfDevice times
	for index = 0; index < numberOfDevice; index++ {
		var searchD string
		fmt.Printf("Type the name of the device you want to search:\n")
		fmt.Scanln(&searchD)
		currentTime := time.Now().UnixNano()
		updateDeviceList(nearbyDeviceList, searchD, DistanceCal(initTime, currentTime))
	}
	fmt.Fprintf(w, "Scanning the nearby Devices...\n...\n")
	fmt.Fprintf(w, "Device been find as following:\n")
	//json.NewEncoder(w).Encode(nearbyDeviceList)
	//print the element in the list to the API website
	for element := nearbyDeviceList.Front(); element != nil; element = element.Next() {
		value := element.Value.(Device)
		fmt.Fprintf(w, "ID: %s, Name: %s, Distance: %fm\n", value.ID, value.Name, value.Distance_m)
	}
	fmt.Println("Endpoint Hit: connection")
}

// access the homepage
func homepage(w http.ResponseWriter, r *http.Request) {
	initTime = time.Now().UnixNano()
	// fmt.Fprintf(w, "Welcome to the homepage of TT. X_x\n")
	// fmt.Fprintf(w, "(Init Time reset!)\n")
	// fmt.Fprintf(w, "There are several features that we can use on this API:\n")
	// fmt.Fprintf(w, "+---------------------------------------+\n")
	// fmt.Fprintf(w, "/currentDistance -> round distance\n/devices -> devices info\n/currentOnewayDistance -> one way distance\n")
	// fmt.Fprintf(w, "+---------------------------------------+\n")
	// fmt.Fprintf(w, "\n\n\n\n\n\n\n\n\n\n\nFounder: Yixuan Shen, HanZhen Qin, Kaiyang Chang\n")
	fmt.Println("Endpoint Hit: homepage")
}

func getDeviceWithID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuidStr := vars["uuid"]
	UUID, _ := uuid.Parse(uuidStr)

	fmt.Println(`UUID: ` + uuidStr)
	fmt.Println("Endpoint Hit: getDeviceWithID")

	for _, device := range devices {
		if device.ID == UUID {
			json.NewEncoder(w).Encode(device)
			fmt.Println("device found with UUID: " + uuidStr)
			return
		}
	}

	fmt.Println("No device found")
	fmt.Fprintf(w, "No device found")
}

func deleteDeviceWithID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuidStr := vars["uuid"]
	UUID, _ := uuid.Parse(uuidStr)

	fmt.Fprintf(w, `UUID: `+uuidStr)
	fmt.Println("Endpoint Hit: deleteDeviceWithID")

	for index, device := range devices {
		if device.ID == UUID {
			devices = append(devices[:index], devices[index+1:]...)
			fmt.Fprintf(w, "Device deleted")
			return
		}
	}

	fmt.Fprintf(w, "No device found")
}

func addDeviceWithID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var device Device
	json.Unmarshal(requestBody, &device)

	for _, d := range devices {
		if d.ID == device.ID {
			fmt.Fprintf(w, "Device already exists")
			return
		}
	}

	devices = append(devices, device)
	fmt.Fprintf(w, "Device added")
}

// Handle the function link to the website
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/currentDistance", currentDistancePage)
	myRouter.HandleFunc("/devices", AllDevices).Methods("GET")

	myRouter.HandleFunc("/device", addDeviceWithID).Methods("POST")
	myRouter.HandleFunc("/device/{uuid}", getDeviceWithID).Methods("GET")
	myRouter.HandleFunc("/device/{uuid}", deleteDeviceWithID).Methods("DELETE")

	myRouter.HandleFunc("/currentOnewayDistance", onewayDistancePage)
	myRouter.HandleFunc("/connection", connection)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

var devices = []Device{
	Device{uuid.New(), "Device1", 50},
	Device{uuid.New(), "Device2", 30},
	Device{uuid.New(), "self", 0},
}

func main() {
	initTime = time.Now().UnixNano()

	fmt.Println("REST API with Mux Routers")
	fmt.Println("Server Started at: http://localhost:10000/")
	handleRequests()
}
