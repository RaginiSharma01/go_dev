package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) { // we r passing reference so we need to pass the address below

	fmt.Printf("worker %d started\n", i)
	// some task!
	fmt.Printf("worker %d end\n", i)
	//wg.Done() // to identify the tasks are executrd
	defer wg.Done() // just before the execution it will run the the wg.done

}

func main() {
	//fmt.Println("groutine started !")
	var wg sync.WaitGroup
	// 3 goroutine

	for i := 1; i <= 3; i++ {

		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Print("worker task executed!")
}
