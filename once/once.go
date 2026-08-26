package main

import (
	"fmt"
	"sync"
)

var num int
var once sync.Once
var increments sync.WaitGroup

func increment() {
	num++
}

func decrement() {
	num--
}

func main() {
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer increments.Done()
			if i%2 == 0 {
				once.Do(decrement)
			} else {
				once.Do(increment)
			}
		}(i)
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", num)
}
