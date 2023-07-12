package main

import (
	"fmt"
	"github.com/tinygo-org/bluetooth"
)

func main() {
	//create a new bluetooth device adapter
	adapter := bluetooth.DefaultAdapter

	//Open the bluetooth device adapter
	access := adapter.Enable()
	if access != nil {
		fmt.Printf("Cannot open the bluetooth device adapter sucessfully: %s\n", access)
		return
	}

	//scan the nearby bluetooth device
	access = adapter.Scan(func(device bluetooth.ScanResult) {
		fmt.Printf("Detect the nearby bluetooth device: %s\n", device.Address)
	})

	if access != nil {
		fmt.Printf("Cannot scan the nearby bluetooth device successfully: %s\n", access)
		return
	}

	//Connecting the the target bluetooth device
	deviceConnecting_address := ""
	device, err := adapter.Connect(deviceConnecting_address)
	if access != nil {
		fmt.Printf("Cannot connect with the corresponding bluetooth device: %s\n", access)
		return
	}

	//Send and receive the message from both bluetooth device adapter, if connect successfully
	go func() {
		for {
			//receive the data or message
			data, access := device.Receive()
			if access != nil {
				fmt.Printf("Receive data is denied: %s\n", access)
				break
			}
			fmt.Printf("Receiving the data: %s\n", string(data))
		}
	}()

	for {
		//Sending the data
		err = device.Send([]byte("Hello, Bluetooth!"))
		if access != nil {
			fmt.Printf("Error to send the data: %s\n", access)
			break
		}
	}

	//disconnect
	device.Disconnect()
}
