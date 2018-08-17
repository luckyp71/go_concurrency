package main

import (
	"fmt"
)

/*
	Example concurrency in go using unbuffered channel
*/

func main() {

	// A buffered channel will block when the buffer is full
	ch := make(chan int, 2)
	// Send data to channel
	ch <- 90
	ch <- 50

	// Send data to the channel by using go routine to prevent deadlock when the buffered channel is full
	ch3 := func() {
		ch <- 40
	}

	go ch3()

	// Channel received value and assign it to v variable
	v1 := <-ch
	v2 := <-ch
	v3 := <-ch

	fmt.Println(v1, v2, v3)
}
