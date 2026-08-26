package main

import "fmt"

func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Recover ", r)
	}
}
func divide(num1, num2 int) int {
	defer handlePanic()
	if num2 == 0 {
		panic("Can't divide by 0")

	} else {
		return num1 / num2
	}
}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(6, 0))
	fmt.Println(divide(2, 8))
}
