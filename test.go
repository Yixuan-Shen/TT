package main

import (
	"fmt"
	"time"
)

const BluetoothDisConst = 0.00029981 / 2

func DistanceCal(time1, time2 int64) float64 {
	Duration := time2 - time1
	Distance := BluetoothDisConst * float64(Duration) // pseudo number for algo
	return Distance
}

func main() {
	currentTime1 := time.Now().UnixNano()
	fmt.Println("Current Timestamp: ", currentTime1)
	time.Sleep(30 * time.Microsecond) // Simulate bluetooth connection and passing time
	currentTime2 := time.Now().UnixNano()
	fmt.Println("Current Timestamp: ", currentTime2)
	dis := DistanceCal(currentTime1, currentTime2)
	fmt.Println(dis)
}
