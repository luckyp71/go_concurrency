package main

import (
	"fmt"
	"sync"
)

/*
	Example concurrency in go using go routine and sync
*/

var wg sync.WaitGroup

func main() {
	fmt.Println("start")

	// Pass one argument indicates we are going to wait for one thing
	wg.Add(1)

	go sayName("name a")

	// Wait for all things to be done
	wg.Wait()

}

func sayName(text string) {
	fmt.Println(text)
	wg.Done() // this is done
}
