package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1
	go func() {
		ch <- 2
	}()

	time.Sleep(1 * time.Second)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// goroutine is parked in sendq (channel send queue) since buffer is full
//  sendq : trying to send 2
