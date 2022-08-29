package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("start")
			fmt.Printf("Worker %d received %d\n", id, <-c)
			fmt.Println("end")
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(0)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func main() {
	//channelClose()
	chanDemo()
	time.Sleep(5 * time.Minute)
}
