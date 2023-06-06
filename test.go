package main

import (
	"fmt"
	"time"
)

func main()  {
	currentTime := time.Now().UnixNano()
	fmt.Println("Current Timestamp: ", currentTime)
}