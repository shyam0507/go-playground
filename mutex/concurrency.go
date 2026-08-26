package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 3)
var money = 100.0
var mutex sync.Mutex

func main() {
	start := time.Now()
	fmt.Println("Hello World")

	go doingLongWork("add", 100, 1)
	fmt.Println("Work Done ", <-ch)
	go doingLongWork("sub", 100, 2)
	fmt.Println("Work Done ", <-ch)
	go doingLongWork("sub", 100, 3)

	fmt.Println("Work Done ", <-ch)

	fmt.Println("Available money", money)

	duration := time.Since(start)
	fmt.Println("Finished all work", duration)
	close(ch)
}

func doingLongWork(op string, amount float64, counter int) {
	if op == "add" {
		mutex.Lock()
		money = money + amount + (amount * 0.1)
		mutex.Unlock()
	} else {
		mutex.Lock()
		money = money - amount - (amount * 0.05)
		mutex.Unlock()
	}
	time.Sleep(time.Second * 2)

	ch <- counter
}
