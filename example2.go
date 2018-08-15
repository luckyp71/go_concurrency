package main

import (
	"fmt"
)

/*
	Example concurrency in go using channel
*/

func main() {
	data := []string{"Name 1", "Name 2", "Name 3"}
	nameList := make(chan []string)
	go getNames(data, nameList)
	// Channel received value, then assign it to the output variable
	output := <-nameList

	// Display the output
	fmt.Println(output)

	// Iterate and display the output item
	for _, o := range output {
		fmt.Println(o)
	}

}

func getNames(names []string, nameList chan []string) {
	data := []string{}
	for _, datum := range names {
		data = append(data, datum)
	}
	// Send value to channel
	nameList <- data
}
