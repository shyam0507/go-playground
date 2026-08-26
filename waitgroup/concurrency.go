package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println(os.Por)
	start := time.Now()
	fmt.Println("Hello World")
	wg.Add(3)
	go doingLongWork(1)
	go doingLongWork(1)
	go doingLongWork(1)

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Finished all work", duration)
}

func doingLongWork(counter int) {
	time.Sleep(time.Second * 2)
	fmt.Println("Work Done ", counter)
	wg.Done()
}
