package main

import (
	"fmt"
	"sync"
	"time"
)

var queue []int

var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go addToQueue(i + 1)

	}
	wg.Wait()

	fmt.Println(len(queue))
}

func addToQueue(n int) {
	time.Sleep(time.Second * 30)
	mutex.Lock()
	queue = append(queue, n)
	wg.Done()
	mutex.Unlock()
}
