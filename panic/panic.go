package main

import "fmt"

func divide(num1, num2 int) int {
	if num2 == 0 {
		panic("Can't divide by 0")
	}

	return num1 / num2

}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(6, 0))
	fmt.Println(divide(2, 8))
}
