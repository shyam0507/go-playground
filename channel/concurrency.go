package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 3)

func main() {
	start := time.Now()
	fmt.Println("Hello World")

	go doingLongWork(1)
	go doingLongWork(2)
	go doingLongWork(3)

	fmt.Println("Work Done ", <-ch)
	fmt.Println("Work Done ", <-ch)
	fmt.Println("Work Done ", <-ch)

	duration := time.Since(start)
	fmt.Println("Finished all work", duration)
	close(ch)
}

func doingLongWork(counter int) {
	time.Sleep(time.Second * 2)

	ch <- counter
}
