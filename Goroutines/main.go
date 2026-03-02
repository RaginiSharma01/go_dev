package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("say hello func ended")
}

func sayHi() {
	fmt.Println("hi")
}

func main() {
	fmt.Println("learning go")
	go sayHello()
	go sayHi()

	time.Sleep(2000 * time.Millisecond)
	// sync wait group
}
