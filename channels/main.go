package main

import (
	"fmt"
	"sync"
)

func main() {

	myChannel := make(chan int, 3)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		//reciving value
		val, isChannelOpen := <-myChannel
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		close(myChannel)

		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
