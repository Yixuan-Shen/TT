package main

import (
	"fmt"
	"time"
)

/*
* Overview
* The distanceCalculation_oneway function: use to calculate the time from another
* device to self device, and then get the distance between the two devices.
* receiveT - sendT
* The distanceCalculation_oneway function only take care the incoming signal time,
* as well as we receive it, we take the calculation.
 */
func distanceCalculation_oneway(sendT, receiveT int64) (distance float64) {
	electronmaticWave := 0.00029981
	receiveTime := receiveT - sendT
	distance = float64(receiveTime) * electronmaticWave
	fmt.Printf("The distance from the device1 is %.2f M.\n", distance)
	return
}

/*
* Overview
* Initialize the time and pass time to the distanceCalculation_oneway.
 */
func timeRecord() {
	sendTime := time.Now().UnixNano()
	fmt.Println("The time sent by another device is ->", sendTime)
	receiveTime := time.Now().UnixNano()
	fmt.Println("The time received by our device is ->", receiveTime)
	distanceCalculation_oneway(sendTime, receiveTime)
}

func main() {
	timeRecord() //test for timeRecord
}
