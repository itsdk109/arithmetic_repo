package main

import (
	"fmt"
	"time"
)

func PrintMsg(msg string) {
	fmt.Println(msg)
}

func Printmsg(){
 fmt.Println("Without Calling printing message")
}

func main() {
	go PrintMsg("Hello World!")
	time.Sleep(1 * time.Second)

	go Printmsg()
	time.Sleep(1* time.Second)

	fmt.Println("Main function Ended")
}
