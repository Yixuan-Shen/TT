package main

import (
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	//Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Start scanning.
	println("Scanning the bluetooth...")
	access_error := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
	})
	must("Start scan", access_error)
}

func must(action string, access_error error) {
	if access_error != nil {
		panic("Insuccessfully to " + action + ": " + access_error.Error())
	}
}
