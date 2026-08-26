package main

import "fmt"

func main() {
	defer fmt.Println("Three")
	defer fmt.Println("Four")

	fmt.Println("One")
	fmt.Println("Two")
}
