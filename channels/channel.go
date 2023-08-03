package main

import (
	"fmt"
	"time"
)

// function
func printString(mychannel chan bool) {
	fmt.Println("THis is simple channel program where you can learn how we create a channel using userdefine function")
	// sending data to a channel 'mychannel'
	mychannel <- true
}

func main() {
	// create an channel
	str := make(chan bool)

	go printString(str)
	time.Sleep(1 * time.Second)
}
