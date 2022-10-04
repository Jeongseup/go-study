package main

import (
	"fmt"
	"mylog"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	mylog.Log()
	fmt.Println("logging 'handled' errors: ")
	mylog.FinalDestination()
}
